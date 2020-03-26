package service

import (
	"errors"

	"jcfw.com/legal-api/model"
	"jcfw.com/legal-api/util"
)

//通过手机号检查账户是否存在
func (b *BaseService) ExistAccountByPhone(phone string) (account model.LegalCompanyAccount, err error) {
	err = b.DB.Select("id,saas_id,company_id,name,uuid,phone").Where("phone = ?", phone).Find(&account).Error
	if err != nil {
		return account, errors.New("账号不存在")
	}
	return account, nil
}

//通过手机和密码登录
func (b *BaseService) LoginByPhoneAndPassword(phone, password string) (accountInfo model.LegalCompanyAccount, e error) {
	accountInfo, err := b.ExistAccountByPhone(phone)
	if err != nil {
		return accountInfo, err
	}
	//判断密码
	accountAuth := model.LegalCompanyAccountAuth{}
	err = b.DB.Select("id,hash").Where("auth_name = ? and account_id = ?", phone, accountInfo.ID).Find(&accountAuth).Error
	if err != nil || !util.PasswordVerify(password, accountAuth.Hash) {
		return accountInfo, errors.New("密码错误")
	}
	return accountInfo, nil
}

//增加一个账号
func (b *BaseService) AddAcount(data *model.LegalCompanyAccount, password string) error {
	//检查是否存在同一个账号
	_, err := b.ExistAccountByPhone(data.Phone)
	if err == nil {
		return errors.New("已存在同手机号账号")
	}
	tx := b.DB.Begin()
	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return errors.New("账号添加失败")
	}
	//生成密码hash
	hash, err := util.PasswordHash(password)
	if err != nil {
		return err
	}
	err = tx.Create(&model.LegalCompanyAccountAuth{
		AccountID: int(data.ID),
		AuthName:  data.Phone,
		SaasID:    data.SaasID,
		Hash:      hash,
	}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("账号添加失败")
	}
	tx.Commit()
	return nil
}

// 更新信息
func (b *BaseService) ModifyAccount(data *model.LegalCompanyAccount) error {
	return b.DB.Updates(data).Error
}

//删除账号
func (b *BaseService) DeleteAccount(ids []int) error {
	account := model.LegalCompanyAccount{}
	accountAuth := model.LegalCompanyAccountAuth{}
	tx := b.DB.Begin()
	err := tx.Where(ids).Delete(&account).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除失败")
	}
	err = tx.Where("account_id IN (?)", ids).Delete(&accountAuth).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除失败")
	}
	tx.Commit()
	return nil
}

package model

import (
	"github.com/jinzhu/gorm"
)

// LegalCompanyAccountAuth 账户认证
type LegalCompanyAccountAuth struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID    int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	AccountID int    `gorm:"column:account_id;type:int(11);not null" json:"account_id"` // account ID
	AuthName  string `gorm:"column:auth_name;type:varchar(32)" json:"auth_name"`        // 登录名
	Hash      string `gorm:"column:hash;type:varchar(255)" json:"hash"`                 // 盐值
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"`    // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyAccountAuth) TableName() string {
	return "legal_company_account_auth"
}

// LegalCompanyAccountRole 员工角色
type LegalCompanyAccountRole struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID    int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	AccountID int    `gorm:"column:account_id;type:int(11);not null" json:"account_id"` // 账号id
	RoleID    int    `gorm:"column:role_id;type:int(11);not null" json:"role_id"`       // 角色id
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"`    // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyAccountRole) TableName() string {
	return "legal_company_account_role"
}

// LegalCompanyAnnouncement 公司公告
type LegalCompanyAnnouncement struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID    int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	Title     string `gorm:"column:title;type:varchar(128);not null" json:"title"`      // 标题
	Content   string `gorm:"column:content;type:varchar(512);not null" json:"content"`  // 内容
	PublishAt int    `gorm:"column:publish_at;type:int(11);not null" json:"publish_at"` // 发布时间
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"`    // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyAnnouncement) TableName() string {
	return "legal_company_announcement"
}

// LegalCompanyDepartment 公司部门
type LegalCompanyDepartment struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID    int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	Name      string `gorm:"column:name;type:varchar(32);not null" json:"name"`      // 部门名称
	Code      string `gorm:"column:code;type:varchar(32);not null" json:"code"`      // 部门代码
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"` // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyDepartment) TableName() string {
	return "legal_company_department"
}

// LegalCompanyPost 公司岗位
type LegalCompanyPost struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID    int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	Name      string `gorm:"column:name;type:varchar(32);not null" json:"name"`      // 岗位名称
	Code      string `gorm:"column:code;type:varchar(32);not null" json:"code"`      // 岗位代码
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"` // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyPost) TableName() string {
	return "legal_company_post"
}

// LegalSystemRole 系统角色
type LegalSystemRole struct {
	gorm.Model
	UUID      string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"`      // uuid
	Name      string `gorm:"column:name;type:varchar(32);not null" json:"name"`      // 角色名称
	Code      string `gorm:"column:code;type:varchar(32);not null" json:"code"`      // 角色代码
	Remark    string `gorm:"column:remark;type:varchar(512);not null" json:"remark"` // 备注
	DeletedBy int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalSystemRole) TableName() string {
	return "legal_system_role"
}

// LegalCompanyAccount 公司账号
type LegalCompanyAccount struct {
	gorm.Model
	UUID         string `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"` // uuid
	SaasID       int    `gorm:"column:saas_id;type:int(11);not null" json:"saas_id"`
	CompanyID    int    `gorm:"column:company_id;type:int(11);not null" json:"company_id"`       // 公司id
	PostID       int    `gorm:"column:post_id;type:int(11);not null" json:"post_id"`             // 岗位id
	DepartmentID int    `gorm:"column:department_id;type:int(11);not null" json:"department_id"` // 部门id
	JobNumber    string `gorm:"column:job_number;type:varchar(6);not null" json:"job_number"`    // 工号
	Name         string `gorm:"column:name;type:varchar(32);not null" json:"name"`               // 姓名
	Phone        string `gorm:"column:phone;type:varchar(11);not null" json:"phone"`             // 手机号
	Email        string `gorm:"column:email;type:varchar(64);not null" json:"email"`             // 邮箱
	EntryDate    int    `gorm:"column:entry_date;type:int(11);not null" json:"entry_date"`       // 入职时间
	Remark       string `gorm:"column:remark;type:varchar(512);not null" json:"remark"`          // 备注
	DeletedBy    int8   `gorm:"column:deleted_by;type:tinyint(1);not null" json:"deleted_by"`
	UpdatedBy    int    `gorm:"column:updated_by;type:int(11);not null" json:"updated_by"`
	CreatedBy    int    `gorm:"column:created_by;type:int(11);not null" json:"created_by"`
}

// TableName get sql table name.获取数据库表名
func (m *LegalCompanyAccount) TableName() string {
	return "legal_company_account"
}

package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

//密码加密
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//验证加密字串
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

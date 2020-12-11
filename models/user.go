package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 用户

var AnonymousUser = User{
	ID:       0,
	Username: "Anonymous User",
	Email:    "Email@not.email",
}

type User struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Openid       string    `json:"openid"`    // 微信登录openid
	LastTime     time.Time `json:"last_time"` // 微信登录openid
	Library      []Library `json:"library"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) HashPassword(pwd string) {

	u.PasswordHash = hashAndSalt([]byte(pwd))
}

// 密码加密
// pwd []byte 密码
func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	return string(hash)
}

// 验证密码
// plainPwd string 计划密码
func (u *User) ComparePasswords(plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(plainPwd)); err == nil {
		return true
	}
	return false
}

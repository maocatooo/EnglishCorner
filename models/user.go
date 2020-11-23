package models

// 用户
type User struct {
	ID uint `json:"id"`

	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Openid       string `json:"openid"` // 微信登录openid

}

func (User) TableName() string {
	return "users"
}

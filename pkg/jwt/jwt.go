package jwt

import (
	"EnglishCorner/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type EcClaims struct {
	jwt.StandardClaims
	User models.User `json:"user"`
}

const TokenExpireDuration = time.Hour * 24 * 3

var Secret = []byte("ec")

func GetExTime() int64 {
	return time.Now().Add(TokenExpireDuration).Unix()
}

func GetClaims(user models.User) EcClaims {
	c := EcClaims{
		User: user,
	}
	c.ExpiresAt = GetExTime()
	c.Issuer = string(Secret)
	return c
}

func GetToken(user models.User) (string, error) {
	c := GetClaims(user)
	return GenToken(c)
}

// GenToken 生成JWT
func GenToken(c EcClaims) (string, error) {

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (*EcClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &EcClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*EcClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func RefreshToken(token string) (string, error) {
	var (
		newToken string
		err      error
	)
	c, err := ParseToken(token)
	c.ExpiresAt = GetExTime()
	newToken, err = GenToken(*c)
	return newToken, err

}

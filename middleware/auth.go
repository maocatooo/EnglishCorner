package middleware

import (
	"EnglishCorner/models"
	"EnglishCorner/pkg/jwt"
	"EnglishCorner/utils/resful/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware(anonymous ...bool) func(c *gin.Context) {
	a := false
	if len(anonymous) != 0 && anonymous[0] {
		a = true
	}
	return loginRequiredUser(a)
}

func loginRequiredUser(anonymous bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			ok   = true
			emsg = "请登录后访问"
		)
		{
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" {
				ok = false
				goto GOTOTHIS
			}
			// 按空格分割
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == string(jwt.Secret)) {
				ok = false
				goto GOTOTHIS
			}
			//
			mc, err := jwt.ParseToken(parts[1])
			if err != nil {
				ok = false
				goto GOTOTHIS
			}

			c.Set("user", mc.User)
		}
	GOTOTHIS:
		if !ok {
			if anonymous {
				c.Set("user", models.AnonymousUser)
			} else {
				response.UnAuthError(c, emsg)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

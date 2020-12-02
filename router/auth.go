package router

import (
	"EnglishCorner/apis/auth"
	"EnglishCorner/middleware"
	"github.com/gin-gonic/gin"
)

func authRouter(rg *gin.RouterGroup) {

	g := rg.Group("/auth")
	ga := g.Use(middleware.JWTAuthMiddleware(true))
	ga.POST("/register", auth.RegisterAuth)
	ga.GET("/code", auth.GetCode)
	ga.POST("/login", auth.LoginAuth)
	ga.POST("/wechat/login/:code", auth.WeChatLogin)
}

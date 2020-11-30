package router

import (
	"EnglishCorner/apis/auth"
	"github.com/gin-gonic/gin"
)

func authRouter(rg *gin.RouterGroup) {

	g := rg.Group("/auth")
	ga := g.Use()
	ga.POST("/register", auth.RegisterAuth)
	ga.GET("/code", auth.GetCode)
	ga.POST("/login", auth.LoginAuth)
}

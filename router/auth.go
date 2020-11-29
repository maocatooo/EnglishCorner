package router

import (
	"EnglishCorner/apis/auth"
	"github.com/gin-gonic/gin"
)

func authRouter(rg *gin.RouterGroup) {

	g := rg.Group("/auth")
	ga := g.Use()
	ga.GET("/register", auth.RegisterAuth)
	ga.GET("/code", auth.GetCode)
	ga.GET("/login", auth.LoginAuth)
}

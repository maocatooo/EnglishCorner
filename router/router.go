package router

import (
	"EnglishCorner/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Cors())
	v1 := r.Group("/api/v1")
	wordsRouter(v1)
	libraryRouter(v1)
	authRouter(v1)
	return r
}

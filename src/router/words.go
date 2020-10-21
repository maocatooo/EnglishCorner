package router

import (
	"EnglishCorner/src/apis"
	"github.com/gin-gonic/gin"
)

func wordsRouter(rg *gin.RouterGroup) {

	{
		wrg := rg.Group("/words")
		rga := wrg.Use()
		rga.GET("/", apis.WordsList)
		rga.POST("/", apis.WordsPost)
		rga.GET("/:id", apis.Words)
		rga.POST("/upload", apis.Upload)
	}
}

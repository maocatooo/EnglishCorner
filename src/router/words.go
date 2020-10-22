package router

import (
	"EnglishCorner/src/apis"
	"github.com/gin-gonic/gin"
)

func wordsRouter(rg *gin.RouterGroup) {

	{
		rg := rg.Group("/words")
		ga := rg.Use()
		ga.GET("", apis.WordsList)
		ga.POST("", apis.WordsPost)
		ga.GET("/:id", apis.Words)
		ga.POST("/upload", apis.Upload)
	}
}

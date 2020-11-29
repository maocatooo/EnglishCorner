package router

import (
	"EnglishCorner/apis/words"
	"github.com/gin-gonic/gin"
)

func wordsRouter(rg *gin.RouterGroup) {

	{
		rg := rg.Group("/words")
		ga := rg.Use()
		ga.GET("", words.WordsList)
		ga.POST("", words.WordsPost)
		ga.GET("/:id", words.Words)
		ga.POST("/upload", words.Upload)
	}
}

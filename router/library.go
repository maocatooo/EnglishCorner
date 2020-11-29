package router

import (
	"EnglishCorner/apis/library"
	"github.com/gin-gonic/gin"
)

func libraryRouter(rg *gin.RouterGroup) {

	{
		g := rg.Group("/library")
		ga := g.Use()
		ga.GET("", library.LibraryList)
		ga.GET("/:id/words", library.LibraryWords)
	}
}

package router

import (
	"EnglishCorner/src/apis"
	"github.com/gin-gonic/gin"
)

func libraryRouter(rg *gin.RouterGroup) {

	{
		g := rg.Group("/library")
		ga := g.Use()
		ga.GET("", apis.LibraryList)
	}
}

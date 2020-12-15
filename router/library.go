package router

import (
	"EnglishCorner/apis/library"
	"EnglishCorner/middleware"
	"github.com/gin-gonic/gin"
)

func libraryRouter(rg *gin.RouterGroup) {

	{
		g := rg.Group("/library")
		ga := g.Use(middleware.JWTAuthMiddleware(true))
		ga.GET("", library.LibraryList)
		ga.GET("/:id/words", library.LibraryWords)
	}
}

package apis

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	"EnglishCorner/utils/resful/request"
	"EnglishCorner/utils/resful/response"
	"github.com/gin-gonic/gin"
)

func LibraryList(c *gin.Context) {
	var (
		ls []models.Library
		DB = db.GetDB()
	)
	o, l := request.GetOffsetLimit(c)
	DB.Limit(l).Offset(o).Find(&ls)
	response.OK(c, ls)
	//c.JSON(200, ls)
}

func LibraryWords(c *gin.Context) {
	var (
		lb models.Library
		ws []models.Word
		DB = db.GetDB()
	)
	id, ok := c.Params.Get("id")
	if !ok {
		response.ParamsError(c, "not Library")
		return
	}
	res := DB.First(&lb, id)
	if res.Error != nil {
		response.ParamsError(c, "not Library")
		return
	}
	o, l := request.GetOffsetLimit(c)
	DB.Model(&lb).Limit(l).Offset(o).Preload(
		"Sentences").Preload(
		"Translations").Association("Words").Find(&ws)
	//DB.Limit(l).Offset(o).Find(&ls)
	response.OK(c, ws)
}

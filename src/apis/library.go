package apis

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/models"
	"EnglishCorner/src/utils"
	"github.com/gin-gonic/gin"
)

func LibraryList(c *gin.Context) {
	var (
		ls []models.Library
		DB = db.GetDB()
		l  = 10
		o  = 0
	)

	offset, ok := c.GetQuery("offset")
	if ok {
		if i, e := utils.StrToInt(offset); e == nil {
			o = i
		}
	}
	limit, ok := c.GetQuery("limit")
	if ok {
		if i, e := utils.StrToInt(limit); e == nil {
			l = i
		}
	}
	DB.Limit(l).Offset(o).Find(&ls)
	c.JSON(200, ls)
}

func LibraryWords(c *gin.Context) {
	var (
		lb models.Library
		ws []models.Word
		DB = db.GetDB()
		l  = 10
		o  = 0
	)
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(422, "not Library")
	}
	res := DB.First(&lb, id)
	if res.Error != nil {
		c.JSON(422, "not Library")
	}
	offset, ok := c.GetQuery("offset")
	if ok {
		if i, e := utils.StrToInt(offset); e == nil {
			o = i
		}
	}
	limit, ok := c.GetQuery("limit")
	if ok {
		if i, e := utils.StrToInt(limit); e == nil {
			l = i
		}
	}
	DB.Model(&lb).Limit(l).Offset(o).Preload(
		"Sentences").Preload(
		"Translations").Association("Words").Find(&ws)
	//DB.Limit(l).Offset(o).Find(&ls)
	c.JSON(200, ws)
}

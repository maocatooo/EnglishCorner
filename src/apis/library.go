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

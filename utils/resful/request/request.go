package request

import (
	"EnglishCorner/utils"
	"github.com/gin-gonic/gin"
)

func GetOffsetLimit(c *gin.Context) (o, l int) {
	l = 10
	o = 0
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
	return
}

package apis

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	rep "EnglishCorner/utils/resful/response"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(c *gin.Context) {
	var (
		user  models.User
		DB    = db.GetDB()
		count int64
	)
	_ = c.BindJSON(&user)

	DB.Model(models.User{}).Where("email=?", user.Email).Count(&count)

	rep.Tell(c, "注册成功")
}

func LoginAuth(c *gin.Context) {
	var (
		words []models.Word
		DB    = db.GetDB()
	)
	id, ok := c.GetQuery("id")
	if !ok {
		id = "1"
	}
	DB.Preload(
		"Sentences").Preload(
		"Translations").Where(
		"id >= ?", id).Limit(
		10).Find(&words, id)
	c.JSON(200, words)
}

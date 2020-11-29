package auth

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	"EnglishCorner/pkg"
	rep "EnglishCorner/utils/resful/response"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type rUser struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"password_again"`
	Code          string `json:"code"`
}

func GetCode(c *gin.Context) {
	captchaId, image := pkg.Captcha()
	key := pkg.GetKey()
	r := db.GetRedisConn()
	r.Do("Set", key, captchaId)
	res, error := redis.String(r.Do("get", key))

	fmt.Println(error)
	fmt.Println(res)
	defer r.Close()
	rep.OK(c, gin.H{
		"image":     image,
		"captchaId": captchaId,
	})

}

func RegisterAuth(c *gin.Context) {
	var (
		ru    rUser
		DB    = db.GetDB()
		count int64
	)

	_ = c.BindJSON(&ru)

	DB.Model(models.User{}).Where("email=?", ru.Email).Count(&count)
	if count > 1 {
		rep.ParamsError(c, "email 已经被注册")
	}
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
	rep.OK(c, gin.H{})
}

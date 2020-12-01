package auth

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	"EnglishCorner/pkg/captcha"
	"EnglishCorner/pkg/jwt"
	rep "EnglishCorner/utils/resful/response"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type rUser struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"password_again"`
	Code          string `json:"code"`
	Key           string `json:"key"`
}

func GetCode(c *gin.Context) {
	captchaId, image := captcha.Captcha()
	key := captcha.GetKey()
	r := db.GetRedisConn()
	_, _ = r.Do("Set", key, captchaId)
	defer r.Close()
	rep.OK(c, gin.H{
		"image": image,
		"key":   key,
	})

}

func RegisterAuth(c *gin.Context) {
	var (
		ru    rUser
		DB    = db.GetDB()
		count int64
		err   error
		r     = db.GetRedisConn()
		user  models.User
	)
	defer r.Close()
	err = c.ShouldBindJSON(&ru)
	if err != nil {
		rep.ParamsError(c, "参数错误")
		return
	}
	if len([]rune(ru.Password)) < 6 {
		rep.ParamsError(c, "密码长度不能小于6位")
		return
	} else if ru.Password != ru.PasswordAgain {
		rep.ParamsError(c, "两次密码不一致")
		return
	}

	captchaId, _ := redis.String(r.Do("get", ru.Key))
	fmt.Println(captchaId)
	if ok := captcha.CaptchaVerify(captchaId, ru.Code); !ok {
		rep.ParamsError(c, "验证码错误或已经过期")
		return
	} else {
		r.Do("delete", ru.Key)
	}

	DB.Model(models.User{}).Where("email=?", ru.Email).Count(&count)
	if count > 1 {
		rep.ParamsError(c, "email 已经被注册")
		return
	}
	user.HashPassword(ru.Password)
	user.Email = ru.Email
	user.LastTime = time.Now()
	_ = DB.Model(models.User{}).Create(&user)

	rep.Tell(c, "注册成功")
}

func LoginAuth(c *gin.Context) {
	var (
		ru   rUser
		DB   = db.GetDB()
		user models.User
		err  error
		eMsg = "邮箱密码错误"
	)
	err = c.ShouldBindJSON(&ru)
	if err != nil {
		rep.ParamsError(c, "参数错误")
		return
	}
	res := DB.Model(models.User{}).Where("email = ?", ru.Email).First(&user)
	if res.RowsAffected < 1 || !user.ComparePasswords(ru.Password) {
		rep.ParamsError(c, eMsg)
		return
	}
	user.PasswordHash = ""
	token, _ := jwt.GetToken(user)

	rep.OK(c, gin.H{
		"token": token,
		"user":  user,
	})
}

func WeChatLogin(c *gin.Context) {
	res, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code")
	if err != nil {

	}
	body, _ := ioutil.ReadAll(res.Body)
	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	if err != nil {

	}
	//res.Body.Read()

}

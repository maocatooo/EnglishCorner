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
		r.Do("delete", ru.Key)
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

	rep.OKAndTell(c, gin.H{
		"token": token,
		"user":  user,
	}, "login success")
}

func GetOpenid(code string) (string, bool) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		"wxa891caa4f40cf683", "684a3375b12d426f65039851b87f550c", code)
	res, err := http.Get(url)
	if err != nil {
		return "", false
	}
	body, _ := ioutil.ReadAll(res.Body)
	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	fmt.Println(r)
	if err != nil {
		return "", false
	}
	return r["openid"].(string), true
}

func WeChatLogin(c *gin.Context) {
	var (
		DB      = db.GetDB()
		user    models.User
		isLogin = false
	)
	u, ok := c.Get("user")
	if ok && u.(models.User).ID != 0 {
		isLogin = true
		user = u.(models.User)
	}
	code, ok := c.Params.Get("code")
	if !ok {
		rep.ParamsError(c, code)
		c.Abort()
		return
	}
	oid, ok := GetOpenid(code)
	if !ok {
		rep.ParamsError(c, code)
		c.Abort()
		return
	}
	if isLogin {
		res := DB.Model(models.User{}).Where("id = ?", user.ID).Update("openid", oid)
		//res := DB.Model(models.User{}).Where("openid = ?", oid).First(&user)
		if res.Error != nil {
			rep.ParamsError(c, code)
			c.Abort()
			return
		}
	} else {
		res := DB.Model(models.User{}).Where("openid = ?", oid).First(&user)
		if res.Error != nil {
			rep.ParamsError(c, code)
			c.Abort()
			return
		}
	}

	user.PasswordHash = ""
	token, _ := jwt.GetToken(user)
	rep.OKAndTell(c, gin.H{
		"token": token,
		"user":  user,
	}, "login success")

}

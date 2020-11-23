package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success       int = 2000 + iota // 正常
	un_auth_error                   // 登录失效
	params_error                    // 参数错误
	server_error                    // 服务器错误
)

type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func resultResponse(c *gin.Context, code int, data interface{}, msg string) {
	var res Response
	res.Code = code
	res.Data = data
	res.Msg = msg
	c.JSON(http.StatusOK, res)
}

func Tell(c *gin.Context, msg string) {
	resultResponse(c, success, nil, msg)
}

func OK(c *gin.Context, data interface{}) {
	resultResponse(c, success, data, "")
}

func ServerError(c *gin.Context, msg string) {
	resultResponse(c, server_error, nil, msg)
}

func ParamsError(c *gin.Context, msg string) {
	resultResponse(c, params_error, nil, msg)
}

func UnAuthError(c *gin.Context, msg string) {
	resultResponse(c, un_auth_error, nil, msg)
}

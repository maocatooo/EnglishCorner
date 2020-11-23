package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(Logger())
	// 自定义错误处理

}

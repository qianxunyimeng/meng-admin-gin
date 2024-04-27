// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/28 00:31:00
// @Desc

package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	// 演示模式 修改数据的操作拦截
	r.Use(DemoEvn())

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// 自定义错误处理
	r.Use(CustomError)

	// 跨域处理
	r.Use(Options)

	r.Use(Secure)
}

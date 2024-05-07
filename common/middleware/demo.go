// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/28 00:34:00
// @Desc
package middleware

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/global"
	"net/http"
)

func DemoEvn() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if global.MA_CONFIG.System.Env == "demo" {
			if method == "GET" ||
				method == "OPTIONS" ||
				c.Request.RequestURI == "/api/v1/login" ||
				c.Request.RequestURI == "/api/v1/logout" {
				c.Next()
			} else {
				// 演示模式下 禁止用户修改数据
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "谢谢您的参与，演示环境切勿修改数据！\U0001F600\U0001F600\U0001F600",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

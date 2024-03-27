// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 21:03:00
// @Desc
package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TrafficKey = "MA_Request_Id"
)

// 获取请求id，如果不存在则生成一个
func GetTranceId(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

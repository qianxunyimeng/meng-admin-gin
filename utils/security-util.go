// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/7 20:21:00
// @Desc
package utils

import "github.com/gin-gonic/gin"

func ExtractClaims(c *gin.Context) *CustomClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return &CustomClaims{}
	}
	return claims.(*CustomClaims)
}

func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	return data.UserId
}

// IsAdminOfUserId 判断用户是否是超级管理员（超级管理员只有一位，系统默认用户admin）
func IsAdminOfUserId(userId int) bool {
	return userId == 1
}

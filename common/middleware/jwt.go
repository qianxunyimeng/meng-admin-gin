// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/31 21:02:00
// @Desc
package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/code"
	"meng-admin-gin/global"
	"meng-admin-gin/utils"
	"net/http"
	"strconv"
	"time"
)

var jwtService = service.JwtService{}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": code.FORBIDDEN,
				"data": nil,
				"msg":  "未登录或非法访问",
			})
			c.Abort()
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenInvalid) {
				c.JSON(http.StatusOK, gin.H{
					"code": code.FORBIDDEN,
					"data": nil,
					"msg":  "token is invalid",
				})
				utils.ClearCookie(c)
				c.Abort()
				return
			} else if errors.Is(err, utils.TokenExpired) {
				c.JSON(http.StatusOK, gin.H{
					"code": code.EXPIRED,
					"data": nil,
					"msg":  "token is expired",
				})
				utils.ClearCookie(c)
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": code.FORBIDDEN,
				"data": nil,
				"msg":  err.Error(),
			})
			utils.ClearCookie(c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			fmt.Println("token is expired,token续期：")
			dr, _ := utils.ParseDuration(global.MA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetCookie(c, newToken, int(dr.Seconds()))
			fmt.Println("续期后的newToken:", newToken)
			if global.MA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.MA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(models.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}

		c.Set("claims", claims)
		c.Set("JWT_PAYLOAD", claims)
		c.Next()
	}
}

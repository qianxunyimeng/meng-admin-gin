// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/31 21:02:00
// @Desc
package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"meng-admin-gin/app/admin/model"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/code"
	"meng-admin-gin/common/response"
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
				"code": code.UNAUTHORIZED,
				"data": nil,
				"msg":  "未登录或非法访问",
			})
			c.Abort()
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				c.JSON(http.StatusOK, gin.H{
					"code": response.Error,
					"data": nil,
					"msg":  "授权已过期",
				})
				utils.ClearCookie(c)
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": response.Error,
				"data": nil,
				"msg":  err.Error(),
			})
			utils.ClearCookie(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.MA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetCookie(c, newToken, int(dr.Seconds()))
			if global.MA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.MA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
	}
}

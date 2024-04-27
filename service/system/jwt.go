// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/31 22:53:00
// @Desc
package system

import (
	"context"
	"meng-admin-gin/common/models"
	"meng-admin-gin/global"
)

type JwtService struct{}

// 检查当前用户是否已经登录
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.MA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// 设置当前用户登录状态
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	err = global.MA_REDIS.Set(context.Background(), userName, jwt, global.MA_JWT_EXP).Err()
	return err
}

// 把旧的jwt 存入redis，防止下次用旧的token请求接口
func (jwtService *JwtService) JsonInBlacklist(jwtList models.JwtBlacklist) (err error) {
	global.MA_CACHE.Set(jwtList.Jwt, struct{}{}, int(global.MA_JWT_EXP.Seconds()))
	return
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, err := global.MA_CACHE.Get(jwt)
	if err != nil {
		return false
	}
	return true
}

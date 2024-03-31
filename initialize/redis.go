// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/29 22:30:00
// @Desc
package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"meng-admin-gin/global"
)

func InitRedis() {
	redisCfg := global.MA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.MA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.MA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.MA_REDIS = client
	}
}

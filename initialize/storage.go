// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:41:00
// @Desc
package initialize

import (
	"go.uber.org/zap"
	"meng-admin-gin/core/captcha"
	"meng-admin-gin/global"
)

func InitStorage() {
	// 初始化redis
	InitRedis()
	cache, err := InitCache()
	if err != nil {
		global.MA_LOG.Fatal("初始化Cache异常：", zap.String("error", err.Error()))
	}
	global.MA_CACHE = cache

	//5. 设置验证码store
	captcha.SetStore(captcha.NewCacheStore(cache, 600))
}

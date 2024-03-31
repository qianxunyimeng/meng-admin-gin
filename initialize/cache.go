// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/29 22:47:00
// @Desc
package initialize

import (
	"meng-admin-gin/core/storage"
	"meng-admin-gin/core/storage/cache"
	"meng-admin-gin/global"
)

func InitCache() (storage.AdapterCache, error) {
	if global.MA_CONFIG.System.CacheType == "redis" {
		options, err := global.MA_CONFIG.Redis.GetRedisOptions()
		r, err := cache.NewRedis(global.MA_REDIS, options)
		if err != nil {
			return nil, err
		}
		if global.MA_REDIS == nil {
			global.MA_REDIS = r.GetClient()
		}
		return r, nil
	}
	return cache.NewMemory(), nil
}

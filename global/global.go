package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"meng-admin-gin/config"
	"meng-admin-gin/core/storage"
	"time"
)

var (
	MA_DB                  *gorm.DB
	MA_CONFIG              config.Config
	MA_VP                  *viper.Viper
	MA_LOG                 *zap.Logger
	MA_TRANS               ut.Translator
	MA_REDIS               *redis.Client
	MA_CACHE               storage.AdapterCache
	MA_Concurrency_Control = &singleflight.Group{}
	MA_JWT_EXP             time.Duration
)

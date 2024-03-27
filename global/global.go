package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"meng-admin-gin/config"
)

var (
	MA_DB     *gorm.DB
	MA_CONFIG config.Config
	MA_VP     *viper.Viper
	MA_LOG    *zap.Logger
	MA_TRANS  ut.Translator
)

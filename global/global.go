package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"meng-admin-gin/config"
)

var (
	MA_CONFIG config.Config
	MA_VP     *viper.Viper
	MA_LOG    *zap.Logger
	MA_TRANS  ut.Translator
)

package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"meng-admin-gin/core/inner"
	"meng-admin-gin/global"
	"meng-admin-gin/utils"
	"os"
)

func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.MA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.MA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.MA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := inner.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.MA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

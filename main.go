package main

import (
	"go.uber.org/zap"
	"meng-admin-gin/global"
	"meng-admin-gin/initialize"
)

func main() {

	// 初始化 viper
	global.MA_VP = initialize.InitViper()

	// 初始化 zap
	global.MA_LOG = initialize.InitZap()
	zap.ReplaceGlobals(global.MA_LOG)

	trans, _ := initialize.InitTranslate("zh")
	global.MA_TRANS = trans
	//initialize.InitRoute()

	initialize.RunServer()
}

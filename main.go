package main

import (
	"fmt"
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
	// 初始化存储系统
	initialize.InitStorage()

	// gin表单校验翻译器
	trans, _ := initialize.InitTranslate("zh")
	global.MA_TRANS = trans

	// 初始化数据库链接
	global.MA_DB = initialize.InitGorm()
	fmt.Println(global.MA_DB)
	if global.MA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.MA_DB.DB()
		defer db.Close()
	}

	initialize.RunServer()
}

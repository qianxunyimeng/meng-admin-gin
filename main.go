package main

import (
	"meng-admin-gin/cmd"
	"meng-admin-gin/global"
)

func main() {

	//// 初始化 viper
	//global.MA_VP = initialize.InitViper()
	//
	//// 初始化 zap
	//global.MA_LOG = initialize.InitZap()
	//zap.ReplaceGlobals(global.MA_LOG)
	//// 初始化存储系统
	//initialize.InitStorage()
	//
	//// gin表单校验翻译器
	//trans, _ := initialize.InitTranslate("zh")
	//global.MA_TRANS = trans
	//
	//// 初始化数据库链接
	//global.MA_DB = initialize.InitGorm()
	//if global.MA_DB == nil {
	//	global.MA_LOG.Panic("初始化数据库链接失败")
	//}
	//initialize.RegisterTables() // 初始化表
	//// 程序结束前关闭数据库链接
	//db, _ := global.MA_DB.DB()
	//defer db.Close()
	//
	//initialize.RunServer()

	cmd.Execute()
	//// 程序结束前关闭数据库链接
	db, _ := global.MA_DB.DB()
	defer db.Close()
}

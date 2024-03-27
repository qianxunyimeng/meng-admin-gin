// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 22:07:00
// @Desc 初始化 gorm
package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"meng-admin-gin/common/models/system"
	"meng-admin-gin/global"
	"os"
	"time"
)

type DBBASE interface {
	GetLogMode() string
}

type _gorm struct{}

var Gorm = new(_gorm)

func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 4 * time.Second, // 慢查询时间级别
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE = &global.MA_CONFIG.Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

func InitGorm() *gorm.DB {
	// TODO 这里可以做成多种数据库

	return InitMysql()
}

// 初始化数据库表
func RegisterTables() {
	db := global.MA_DB
	err := db.AutoMigrate(
		system.SysApi{},  // api
		system.SysUser{}, // 用户
		system.SysMenu{}, // 菜单
		system.SysRole{}, // 角色
		system.SysDept{}, // 部门
		system.SysPost{}, // 岗位
	)
	if err != nil {
		global.MA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.MA_LOG.Info("register table success")
}

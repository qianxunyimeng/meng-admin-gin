package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"meng-admin-gin/core/inner"
	"meng-admin-gin/global"
	"meng-admin-gin/utils"
	"os"
	"path/filepath"
)

func InitViper() *viper.Viper {
	var configPath string
	// 使用flag包来接收命令行参数 -c 指定配置文件
	// go run main.go -c ./config/config.yaml
	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()
	if configPath == "" { // 命令行参数 没有指定配置文件路径
		if configPathEnv := os.Getenv(inner.ConfigPathEnv); configPathEnv == "" {
			fmt.Println("gin mode:", gin.Mode())
			switch gin.Mode() {
			case gin.DebugMode: // 开发模式
				configPath = inner.ConfigDevFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, inner.ConfigDevFile)
			case gin.ReleaseMode: // 生产模式
				configPath = inner.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, inner.ConfigReleaseFile)
			case gin.TestMode: // 测试模式
				configPath = inner.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, inner.ConfigTestFile)
			}
		} else {
			configPath = configPathEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", inner.ConfigPathEnv, configPath)
		}
	}

	v := viper.New()
	// 配置文件路径
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.MA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.MA_CONFIG); err != nil {
		panic(err)
	}

	dr, err := utils.ParseDuration(global.MA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	global.MA_JWT_EXP = dr

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.MA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

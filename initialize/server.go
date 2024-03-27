package initialize

import (
	"fmt"
	"github.com/fvbock/endless"
	"meng-admin-gin/core"
	"meng-admin-gin/global"
	"meng-admin-gin/router"
	"net/http"
	"time"
)

var AppRouters = make([]func(), 0)

type server interface {
	ListenAndServe() error
}

// 启动http 服务
func RunServer() {

	// 初始化总路由
	Routes()

	for _, f := range AppRouters {
		f()
	}

	address := fmt.Sprintf(":%d", global.MA_CONFIG.System.Port)
	s := initServer(address, core.Runtime.GetEngine())

	fmt.Printf(`
	欢迎使用 meng admin
	当前版本：v1.0.0
`)
	s.ListenAndServe().Error()
}

func initServer(address string, router http.Handler) server {
	fmt.Println("server other initServer")
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}

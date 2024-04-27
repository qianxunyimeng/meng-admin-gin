// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 21:50:00
// @Desc

package api

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	common "meng-admin-gin/common/middleware"
	"meng-admin-gin/core"
	"meng-admin-gin/global"
	"meng-admin-gin/initialize"
	"meng-admin-gin/router"
	"meng-admin-gin/utils"
	"net/http"
	"time"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "启动api服务",
		Example:      "meng-admin server -c config.dev.yaml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

type server interface {
	ListenAndServe() error
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config.dev.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

// 执行 server 子命令前的初始化操作
func setup() {
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
	if global.MA_DB == nil {
		global.MA_LOG.Panic("初始化数据库链接失败")
	}
	//initialize.RegisterTables() // 初始化表
	// 程序结束前关闭数据库链接
	//db, _ := global.MA_DB.DB()
	//defer db.Close()

	usageStr := `starting api server...`
	log.Println(usageStr)
}

// 初始化gin
func run() error {
	// 生产模式
	if global.MA_CONFIG.System.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	address := fmt.Sprintf(":%d", global.MA_CONFIG.System.Port)
	s := initServer(address, core.Runtime.GetEngine())
	tip()
	fmt.Println("Server run at:")
	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", global.MA_CONFIG.System.Port)
	fmt.Printf("-  Network: %s://%s:%d/ \r\n", "http", utils.GetLocaHonst(), global.MA_CONFIG.System.Port)
	fmt.Println("Swagger run at:")
	fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", global.MA_CONFIG.System.Port)
	fmt.Printf("-  Network: %s://%s:%d/swagger/admin/index.html \r\n", "http", utils.GetLocaHonst(), global.MA_CONFIG.System.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", utils.GetCurrentTimeStr())
	s.ListenAndServe().Error()

	return nil
}

func initRouter() {

	var r *gin.Engine
	h := core.Runtime.GetEngine()

	if h == nil {
		h = gin.New()
		core.Runtime.SetEngine(h)
	}

	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		//os.Exit(-1)
	}

	common.InitMiddleware(r)

}

// endless 平滑重启服务
// kill -1 pid  如果知道pid使用该方法重启服务
// ps aux | grep "meng-admin" | grep -v grep | awk '{print $2}' | xargs -i kill -1 {}  使用进程名称重启服务，test_endless是进程名称，需要替换自己的
func initServer(address string, router http.Handler) server {
	fmt.Println("server other initServer")
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func tip() {
	usageStr := `欢迎使用 meng-admin 0.0.1 可以使用 -h 查看命令`
	fmt.Printf("%s\n", usageStr)
}

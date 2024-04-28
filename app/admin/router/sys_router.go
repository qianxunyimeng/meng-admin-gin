package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meng-admin-gin/app/admin/api"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
	"meng-admin-gin/docs"
	"meng-admin-gin/global"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	if global.MA_CONFIG.System.Env != "prod" {
		sysSwaggerRouter(g)
	}

	systemInternalRouteInit(r, authMiddleware)

	return g
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	//r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("admin")))
	// docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	docs.SwaggerInfo.BasePath = global.MA_CONFIG.System.RouterPrefix
	r.GET(global.MA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// 初始化系统内置路由
func systemInternalRouteInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	fmt.Println("系统内置路由")
	api := api.SysBaseApi{}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", api.Login)
		v1.GET("/captcha", api.GenerateCaptcha)
	}

	registerBaseRouter(v1)
}

func registerBaseRouter(v1 *gin.RouterGroup) {
	v1auth := v1.Group("").Use(middleware.JWTAuth())
	api := api.SysBaseApi{}
	{
		v1auth.POST("/logout", api.Logout)
	}
}

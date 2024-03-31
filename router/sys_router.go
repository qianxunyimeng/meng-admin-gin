package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meng-admin-gin/api/system"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	systemInternalRouteInit(r, authMiddleware)
	return g
}

// 初始化系统内置路由
func systemInternalRouteInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	fmt.Println("系统内置路由")
	api := system.SysApiGroup{}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", api.Login)
		v1.GET("/captcha", api.GenerateCaptcha)
	}

	registerBaseRouter(v1)
}

func registerBaseRouter(v1 *gin.RouterGroup) {
	v1auth := v1.Group("").Use(middleware.JWTAuth())
	api := system.SysApiGroup{}
	{
		v1auth.POST("/logout", api.Logout)
	}
}

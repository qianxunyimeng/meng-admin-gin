package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meng-admin-gin/api/system"
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
		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}

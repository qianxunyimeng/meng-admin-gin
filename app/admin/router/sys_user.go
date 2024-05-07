package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/app/admin/api"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := api.SysUserApi{}
	v1 := r.Group("sys-user")
	{
		v1.POST("", api.Register)
		// Refresh time can be longer than token timeout
	}

	v1auth := r.Group("").Use(middleware.JWTAuth())
	{
		v1auth.GET("/getInfo", api.GetInfo)
	}
}

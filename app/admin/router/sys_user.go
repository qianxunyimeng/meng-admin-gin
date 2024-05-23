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
	r1 := r.Group("sys-user")
	{
		r1.GET("", api.GetPage)
		r1.GET("/:id", api.Get)
		r1.POST("", api.Register)
		r1.PUT("", api.Update)
		r1.DELETE("", api.Delete)
		r1.PUT("/status/change", api.UpdateStatus)
	}
	user := r.Group("user")
	{
		user.PUT("/pwd/reset", api.ResetPwd)
	}

	r1auth := r.Group("").Use(middleware.JWTAuth())
	{
		r1auth.GET("/getInfo", api.GetInfo)
	}
}

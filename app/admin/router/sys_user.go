package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/api/system"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := system.SysApiGroup{}
	v1 := r.Group("sys-user")
	{
		v1.POST("", api.Register)
		// Refresh time can be longer than token timeout

	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/api/system"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

// 需认证的路由代码
func registerSysMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := system.MenuApi{}

	r := v1.Group("/menu").Use(middleware.JWTAuth())
	{
		r.GET("", api.GetMenu)
		r.POST("", api.Insert)
	}

}

package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/app/admin/api"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

// 需认证的路由代码
func registerSysMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := api.MenuApi{}

	r := v1.Group("/menu").Use(middleware.JWTAuth())
	{
		r.GET("", api.GetMenuList)
		r.POST("", api.Insert)
		r.GET("/:menuId", api.GetMenuById)
		r.PUT("/:menuId", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("/menu").Use(middleware.JWTAuth())
	{
		r1.GET("/getRouters", api.GetRouters)
	}

}

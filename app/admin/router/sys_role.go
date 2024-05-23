// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/15 21:41:00
// @Desc
package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/app/admin/api"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

func registerSysRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := api.SysRoleApi{}
	r := v1.Group("/role").Use(middleware.JWTAuth())
	{
		r.GET("", api.GetPageList)
		r.POST("", api.Insert)
		r.GET("/:id", api.GetDetail)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/status/change", api.Update2Status)
		r.PUT("/auth/menu", api.UpdateMenuAuthority)
	}
}

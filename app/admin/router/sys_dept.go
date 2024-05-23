// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/12 16:52:00
// @Desc 部门管理 路由

package router

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/app/admin/api"
	"meng-admin-gin/common/middleware"
	jwt "meng-admin-gin/core/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

func registerSysDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := api.SysDeptApi{}
	r := v1.Group("/dept").Use(middleware.JWTAuth())
	{
		r.GET("", api.GetPageList)
		r.GET("/tree", api.GetDeptTree)
		r.POST("", api.Insert)
		r.GET("/:id", api.GetDetailById)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

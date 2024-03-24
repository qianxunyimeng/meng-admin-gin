package router

import (
	"github.com/gin-gonic/gin"
	jwt "meng-admin-gin/core/jwtauth"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	return g
}

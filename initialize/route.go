package initialize

import (
	"github.com/gin-gonic/gin"
	"log"
	"meng-admin-gin/core"
	"meng-admin-gin/global"
)

func Routes() {

	var r *gin.Engine
	h := core.Runtime.GetEngine()

	if h == nil {
		h = gin.New()
		core.Runtime.SetEngine(h)
	}

	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		//os.Exit(-1)
	}

	r.Use(gin.Recovery())
	if global.MA_CONFIG.System.Env != "public" {
		r.Use(gin.Logger())
	}

}

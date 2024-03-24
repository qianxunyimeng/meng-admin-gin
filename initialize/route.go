package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"meng-admin-gin/core"
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

	fmt.Println(r)

}

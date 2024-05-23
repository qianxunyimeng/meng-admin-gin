// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/18 22:40:00
// @Desc
package dto

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"meng-admin-gin/global"
	"net/http"
)

type ObjectById struct {
	Id  int   `uri:"id"`
	Ids []int `json:"ids"`
}

func (s *ObjectById) Bind(ctx *gin.Context) error {
	var err error
	log := global.MA_LOG
	err = ctx.ShouldBindUri(s)
	if err != nil {
		log.Warn("ShouldBindUri error: " + err.Error())
		return err
	}
	if ctx.Request.Method == http.MethodDelete {
		err = ctx.ShouldBind(&s)
		if err != nil {
			log.Warn("ShouldBind error: " + err.Error())
			return err
		}
		if len(s.Ids) > 0 {
			return nil
		}
		if s.Ids == nil {
			s.Ids = make([]int, 0)
		}
		if s.Id != 0 {
			s.Ids = append(s.Ids, s.Id)
		}
	}
	if err = vd.Validate(s); err != nil {
		log.Error("Validate error: " + err.Error())
		return err
	}
	return err
}

func (s *ObjectById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

type ObjectGetReq struct {
	Id int `uri:"id"`
}

func (s *ObjectGetReq) Bind(ctx *gin.Context) error {
	var err error
	log := global.MA_LOG
	err = ctx.ShouldBindUri(s)
	if err != nil {
		log.Warn("ShouldBindUri error: " + err.Error())
		return err
	}
	if err = vd.Validate(s); err != nil {
		log.Error("Validate error: " + err.Error())
		return err
	}
	return err
}

func (s *ObjectGetReq) GetId() interface{} {
	return s.Id
}

type ObjectDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ObjectDeleteReq) Bind(ctx *gin.Context) error {
	var err error
	log := global.MA_LOG
	err = ctx.ShouldBind(&s)
	if err != nil {
		log.Warn("ShouldBind error: " + err.Error())
		return err
	}
	if len(s.Ids) > 0 {
		return nil
	}
	if s.Ids == nil {
		s.Ids = make([]int, 0)
	}

	if err = vd.Validate(s); err != nil {
		log.Error("Validate error: " + err.Error())
		return err
	}
	return err
}

func (s *ObjectDeleteReq) GetId() interface{} {
	return s.Ids
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/api"
	"meng-admin-gin/common/code"
	"meng-admin-gin/core/jwtauth/user"
)

type MenuApi struct {
	api.Api
}

func (e MenuApi) GetMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "menu",
	})
}

func (e MenuApi) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := service.SysMenuService{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.INVALID_PARAM, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

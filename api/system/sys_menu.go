package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/common/dto"
	"meng-admin-gin/common/respcode"
	"meng-admin-gin/core/api"
	"meng-admin-gin/core/jwtauth/user"
	sysService "meng-admin-gin/service/system"
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
	s := sysService.SysMenuService{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(respcode.ErrorParam, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.Error(respcode.Error, err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

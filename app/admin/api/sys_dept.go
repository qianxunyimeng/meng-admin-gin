// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/12 16:55:00
// @Desc 部门管理 api

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/api"
	"meng-admin-gin/common/code"
	"meng-admin-gin/utils"
)

type SysDeptApi struct {
	api.Api
}

func (e SysDeptApi) GetPageList(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	list := make([]models.SysDept, 0)
	list, err = s.GetDeptPage(&req)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

func (e SysDeptApi) GetDeptTree(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	list := make([]models.SysDept, 0)
	list, err = s.GetDeptPage(&req)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

func (e SysDeptApi) GetDetailById(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	var object models.SysDept

	err = s.Get(&req, &object)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

func (e SysDeptApi) Insert(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}

	// 设置创建人
	req.SetCreateBy(utils.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

func (e SysDeptApi) Update(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	req.SetUpdateBy(utils.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e SysDeptApi) Delete(c *gin.Context) {
	s := service.SysDeptService{}
	req := dto.SysDeptDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

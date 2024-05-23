// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/15 21:42:00
// @Desc
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/api"
	"meng-admin-gin/common/code"
	"meng-admin-gin/utils"
)

type SysRoleApi struct {
	api.Api
}

func (e SysRoleApi) GetPageList(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.SysRoleGetPageReq{}
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
	list := make([]models.SysRole, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	//e.OK(list, "查询成功")
	e.PageOK(list, int(count), req.GetPageNum(), req.GetPageSize(), "查询成功")
}

// GetDetail
// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param roleId path string false "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [get]
// @Security Bearer
func (e SysRoleApi) GetDetail(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}

	var object models.SysRole

	err = s.Get(&req, &object)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [post]
// @Security Bearer
func (e SysRoleApi) Insert(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.SysRoleInsertReq{}
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
	req.CreateBy = utils.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}
	//cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, "创建失败,"+err.Error())
		return
	}
	//_, err = global.LoadPolicy(c)
	//if err != nil {
	//	e.Logger.Error(err)
	//	e.Error(500, err, "创建失败,"+err.Error())
	//	return
	//}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户角色
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [put]
// @Security Bearer
func (e SysRoleApi) Update(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	//cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	req.SetUpdateBy(utils.GetUserId(c))

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}

	//_, err = global.LoadPolicy(c)
	//if err != nil {
	//	e.Logger.Error(err)
	//	e.Error(500, err, "更新失败,"+err.Error())
	//	return
	//}

	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色/Role
// @Param data body dto.SysRoleDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [delete]
// @Security Bearer
func (e SysRoleApi) Delete(c *gin.Context) {
	s := new(service.SysRoleService)
	req := dto.SysRoleDeleteReq{}
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

	//cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(500, err)
		return
	}

	e.OK(req.GetId(), fmt.Sprintf("删除角色角色 %v 状态成功！", req.GetId()))
}

// Update2Status 修改用户角色状态
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-status/{id} [put]
// @Security Bearer
func (e SysRoleApi) Update2Status(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.UpdateStatusReq{}
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
	req.SetUpdateBy(utils.GetUserId(c))
	err = s.UpdateStatus(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(500, err)
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

func (e SysRoleApi) UpdateMenuAuthority(c *gin.Context) {
	s := service.SysRoleService{}
	req := dto.UpdateMenuAuthReq{}
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
	err = s.UpdateRoleMenu(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(500, err)
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

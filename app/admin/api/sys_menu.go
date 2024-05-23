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

type MenuApi struct {
	api.Api
}

// GetMenuList Menu列表数据
// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /api/v1/menu [get]
func (e MenuApi) GetMenuList(c *gin.Context) {
	s := service.SysMenuService{}
	req := dto.SysMenuGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.INVALID_PARAM, err)
		return
	}
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list).Error
	if err != nil {
		e.Error(code.ERROR, err)
		return
	}
	e.OK(list, "获取成功")

}

// GetMenuById 根据菜单id获取菜单详细
func (e MenuApi) GetMenuById(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenuService)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.INVALID_PARAM, err)
		return
	}

	var object = models.SysMenu{}
	err = s.GetMenuById(&req, &object).Error
	if err != nil {
		e.Error(code.ERROR, err)
		return
	}
	e.OK(object, "查询成功")
}

// Insert 创建菜单
// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMenuInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /api/v1/menu [post]
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
	//fmt.Println("userId: ", utils.GetUserId(c))
	req.SetCreateBy(utils.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改菜单
// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysMenuInsertReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [put]
func (e MenuApi) Update(c *gin.Context) {
	req := dto.SysMenuUpdatetReq{}
	s := new(service.SysMenuService)
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
	err = s.Update(&req).Error
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param data body dto.SysMenuDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [delete]
func (e MenuApi) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	s := new(service.SysMenuService)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(control, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}
	err = s.Remove(control).Error
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, "删除失败")
		return
	}
	e.OK(control.GetId(), "删除成功")
}

// GetRouters 获取登录用户的菜单列表
// @Summary 获取登录用户的菜单列表
// @Description 获取JSON
// @Tags 菜单
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /api/v1/menu/getRouters [get]
func (e MenuApi) GetRouters(c *gin.Context) {
	s := new(service.SysMenuService)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.Error(code.ERROR, err)
		return
	}

	menus, err := s.GetRouterByUserId(utils.GetRoleCode(c))
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	result, err := s.BuildMenus(menus)
	if err != nil {
		e.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	e.OK(result, "")
}

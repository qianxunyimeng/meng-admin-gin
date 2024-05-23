// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 17:37:00
// @Desc

package dto

import (
	"github.com/jinzhu/copier"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/common/dto"
	common "meng-admin-gin/common/model"
)

type SysMenuGetPageReq struct {
	dto.Pagination `search:"-"`
	MenuName       string `form:"menuName" search:"type:contains;column:menu_name;table:sys_menu" comment:"菜单名称"` // 菜单名称
	Visible        string `form:"visible" search:"type:exact;column:visible;table:sys_menu" comment:"显示状态"`       // 显示状态
	Status         string `form:"status" search:"type:exact;column:status;table:sys_menu" comment:"数据状态"`         // 显示状态
}

func (m *SysMenuGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysMenuInsertReq struct {
	MenuId       int    `uri:"menuId" comment:"编码"`                                        // 编码
	MenuName     string `form:"menuName" comment:"菜单name" vd:"@:len($)>0; msg:'菜单名称不能为空'"` //菜单name
	Title        string `form:"title" comment:"显示名称"`                                      //显示名称
	Icon         string `form:"icon" comment:"图标"`                                         //图标
	Path         string `form:"path" comment:"路径" validate:"required"`                     //路径
	Paths        string `form:"paths" comment:"id路径"`                                      //id路径
	MenuType     string `form:"menuType" comment:"菜单类型"`                                   //菜单类型
	ViewType     string `form:"viewType" comment:"试图类型，1:普通页面 2:外链页面 3:内嵌页面"`
	Action       string `form:"action" comment:"请求方式"`      //请求方式
	ParentId     int    `form:"parentId" comment:"上级菜单"`    //上级菜单
	NoCache      bool   `form:"noCache" comment:"是否缓存"`     //是否缓存
	Breadcrumb   string `form:"breadcrumb" comment:"是否面包屑"` //是否面包屑
	Component    string `form:"component" comment:"组件"`     //组件
	Permission   string `form:"permission" comment:"权限字符"`
	Sort         int    `form:"sort" comment:"排序"`      //排序
	Visible      string `form:"visible" comment:"是否显示"` //是否显示
	Status       string `form:"status" comment:"菜单状态，1:正常 0:停用"`
	IsFrame      string `form:"isFrame" comment:"是否frame"` //是否frame
	IsInternally string `form:"isInternally" comment:"是否是系统内置数据 1 是 0 否"`
	common.ControlBy
}

func (s *SysMenuInsertReq) Generate(model *models.SysMenu) {
	//if s.MenuId != 0 {
	//	model.MenuId = s.MenuId
	//}
	//model.MenuName = s.MenuName
	//model.Title = s.Title
	//model.Icon = s.Icon
	//model.Path = s.Path
	//model.Paths = s.Paths
	//model.MenuType = s.MenuType
	//model.Action = s.Action
	////model.SysApi = s.SysApi
	////model.Permission = s.Permission
	//model.ParentId = s.ParentId
	//model.NoCache = s.NoCache
	//model.Breadcrumb = s.Breadcrumb
	//model.Component = s.Component
	//model.Sort = s.Sort
	//model.Visible = s.Visible
	//model.IsFrame = s.IsFrame
	//if s.CreateBy != 0 {
	//	model.CreateBy = s.CreateBy
	//}
	//if s.UpdateBy != 0 {
	//	model.UpdateBy = s.UpdateBy
	//}

	copier.Copy(model, s)
}

func (s *SysMenuInsertReq) GetId() interface{} {
	return s.MenuId
}

type SysMenuGetReq struct {
	MenuId int `uri:"menuId"`
}

func (s *SysMenuGetReq) GetId() interface{} {
	return s.MenuId
}

type SysMenuDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysMenuDeleteReq) GetId() interface{} {
	return s.Ids
}

type SysMenuUpdatetReq struct {
	MenuId       int    `uri:"menuId" comment:"编码"`                                        // 编码
	MenuName     string `form:"menuName" comment:"菜单name" vd:"@:len($)>0; msg:'菜单名称不能为空'"` //菜单name
	Title        string `form:"title" comment:"显示名称"`                                      //显示名称
	Icon         string `form:"icon" comment:"图标"`                                         //图标
	Path         string `form:"path" comment:"路径"`                                         //路径
	Paths        string `form:"paths" comment:"id路径"`                                      //id路径
	MenuType     string `form:"menuType" comment:"菜单类型"`                                   //菜单类型
	ViewType     string `form:"viewType" comment:"试图类型，1:普通页面 2:外链页面 3:内嵌页面"`
	Action       string `form:"action" comment:"请求方式"`      //请求方式
	ParentId     int    `form:"parentId" comment:"上级菜单"`    //上级菜单
	NoCache      bool   `form:"noCache" comment:"是否缓存"`     //是否缓存
	Breadcrumb   string `form:"breadcrumb" comment:"是否面包屑"` //是否面包屑
	Component    string `form:"component" comment:"组件"`     //组件
	Permission   string `form:"permission" comment:"权限字符"`
	Sort         int    `form:"sort" comment:"排序"`      //排序
	Visible      string `form:"visible" comment:"是否显示"` //是否显示
	Status       string `form:"status" comment:"菜单状态，1:正常 0:停用"`
	IsFrame      string `form:"isFrame" comment:"是否frame"` //是否frame
	IsInternally string `form:"isInternally" comment:"是否是系统内置数据 1 是 0 否"`
	common.ControlBy
}

func (s *SysMenuUpdatetReq) Generate(model *models.SysMenu) {

	copier.CopyWithOption(model, s, copier.Option{
		IgnoreEmpty: true,
	})
}

func (s *SysMenuUpdatetReq) GetId() interface{} {
	return s.MenuId
}

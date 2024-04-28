// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 17:37:00
// @Desc

package dto

import (
	"meng-admin-gin/app/admin/model"
	common "meng-admin-gin/common/model"
)

type SysMenuInsertReq struct {
	MenuId       int    `uri:"id" comment:"编码"`                                            // 编码
	MenuName     string `form:"menuName" comment:"菜单name" vd:"@:len($)>0; msg:'菜单名称不能为空'"` //菜单name
	Title        string `form:"title" comment:"显示名称"`                                      //显示名称
	Icon         string `form:"icon" comment:"图标"`                                         //图标
	Path         string `form:"path" comment:"路径" vd:"@:len($)>0; msg:'路由地址不能为空'"`         //路径
	Paths        string `form:"paths" comment:"id路径"`                                      //id路径
	MenuType     string `form:"menuType" comment:"菜单类型"`                                   //菜单类型
	Action       string `form:"action" comment:"请求方式"`                                     //请求方式
	ParentId     int    `form:"parentId" comment:"上级菜单"`                                   //上级菜单
	NoCache      bool   `form:"noCache" comment:"是否缓存"`                                    //是否缓存
	Breadcrumb   string `form:"breadcrumb" comment:"是否面包屑"`                                //是否面包屑
	Component    string `form:"component" comment:"组件"`                                    //组件
	Sort         int    `form:"sort" comment:"排序"`                                         //排序
	Visible      string `form:"visible" comment:"是否显示"`                                    //是否显示
	Status       string `form:"status" comment:"菜单状态，1:正常 0:停用"`
	IsFrame      string `form:"isFrame" comment:"是否frame"` //是否frame
	IsInternally string `form:"isInternally" comment:"是否是系统内置数据 1 是 0 否"`
	common.ControlBy
}

func (s *SysMenuInsertReq) Generate(model *model.SysMenu) {
	if s.MenuId != 0 {
		model.MenuId = s.MenuId
	}
	model.MenuName = s.MenuName
	model.Title = s.Title
	model.Icon = s.Icon
	model.Path = s.Path
	model.Paths = s.Paths
	model.MenuType = s.MenuType
	model.Action = s.Action
	//model.SysApi = s.SysApi
	//model.Permission = s.Permission
	model.ParentId = s.ParentId
	model.NoCache = s.NoCache
	model.Breadcrumb = s.Breadcrumb
	model.Component = s.Component
	model.Sort = s.Sort
	model.Visible = s.Visible
	model.IsFrame = s.IsFrame
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysMenuInsertReq) GetId() interface{} {
	return s.MenuId
}

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 23:39:00
// @Desc
package model

import (
	common "meng-admin-gin/common/model"
)

type SysMenu struct {
	MenuId       int       `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName     string    `json:"menuName" gorm:"size:128;comment:菜单名称"`
	Title        string    `json:"title" gorm:"size:128;"`
	Icon         string    `json:"icon" gorm:"size:128;"`
	Path         string    `json:"path" gorm:"size:128;"`
	Paths        string    `json:"paths" gorm:"size:128;"`
	MenuType     string    `json:"menuType" gorm:"size:1;"`
	ViewType     string    `json:"viewType" gorm:"size:1;comment:试图类型，1:普通页面 2:外链页面 3:内嵌页面"`
	Action       string    `json:"action" gorm:"size:16;"`
	Permission   string    `json:"permission" gorm:"size:255;"`
	ParentId     int       `json:"parentId" gorm:"size:11;"`
	NoCache      bool      `json:"noCache" gorm:"size:8;"`
	Breadcrumb   string    `json:"breadcrumb" gorm:"size:255;"`
	Component    string    `json:"component" gorm:"size:255;"`
	Sort         int       `json:"sort" gorm:"size:4;comment:显示排序"`
	Visible      string    `json:"visible" gorm:"size:1;comment:菜单状态，1:显示 0:隐藏"`
	Status       string    `json:"status" gorm:"size:1;comment:菜单状态，1:正常 0:停用"`
	IsFrame      string    `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	SysApi       []SysApi  `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`
	Apis         []int     `json:"apis" gorm:"-"`
	DataScope    string    `json:"dataScope" gorm:"-"`
	Params       string    `json:"params" gorm:"-"`
	RoleId       int       `gorm:"-"`
	Children     []SysMenu `json:"children,omitempty" gorm:"-"`
	IsSelect     bool      `json:"is_select" gorm:"-"`
	IsInternally string    `json:"is_internally" gorm:"size:1;DEFAULT:0;comment:是否是系统内置数据，内置数据不可删除"`
	common.ControlBy
	common.ModelTime
}

type SysMenuSlice []SysMenu

func (x SysMenuSlice) Len() int           { return len(x) }
func (x SysMenuSlice) Less(i, j int) bool { return x[i].Sort < x[j].Sort }
func (x SysMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (*SysMenu) TableName() string {
	return "sys_menu"
}

func (e *SysMenu) Generate() common.ActiveRecord {
	o := *e
	return &o
}

func (e *SysMenu) GetId() interface{} {
	return e.MenuId
}

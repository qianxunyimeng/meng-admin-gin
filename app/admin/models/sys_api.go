// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 23:39:00
// @Desc
package models

import (
	common "meng-admin-gin/common/model"
)

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:标题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Action string `json:"action" gorm:"size:16;comment:请求类型"`
	Type   string `json:"type" gorm:"size:16;comment:接口类型"`
	common.ModelTime
	common.ControlBy
}

func (*SysApi) TableName() string {
	return "sys_api"
}

func (e *SysApi) Generate() common.ActiveRecord {
	o := *e
	return &o
}

func (e *SysApi) GetId() interface{} {
	return e.Id
}

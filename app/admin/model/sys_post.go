// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 23:37:00
// @Desc
package model

import (
	common "meng-admin-gin/common/model"
)

type SysPost struct {
	PostId   int    `gorm:"primaryKey;autoIncrement" json:"postId"` //岗位编号
	PostName string `gorm:"size:128;" json:"postName"`              //岗位名称
	PostCode string `gorm:"size:128;" json:"postCode"`              //岗位代码
	Sort     int    `gorm:"size:4;" json:"sort"`                    //岗位排序
	Status   int    `gorm:"size:4;" json:"status"`                  //状态
	Remark   string `gorm:"size:255;" json:"remark"`                //描述
	common.ControlBy
	common.ModelTime

	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
}

func (*SysPost) TableName() string {
	return "sys_post"
}

func (e *SysPost) Generate() common.ActiveRecord {
	o := *e
	return &o
}

func (e *SysPost) GetId() interface{} {
	return e.PostId
}

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 22:57:00
// @Desc
package model

import (
	common "meng-admin-gin/common/model"
)

type SysRole struct {
	RoleId    uint       `json:"roleId" gorm:"primaryKey;comment:角色ID;size:90"` // 角色ID
	RoleName  string     `json:"roleName" gorm:"comment:角色名"`                   // 角色名
	Status    string     `json:"status" gorm:"size:4;"`                         // 状态 1禁用 2正常
	RoleKey   string     `json:"roleKey" gorm:"size:128;"`                      //角色代码
	RoleSort  int        `json:"roleSort" gorm:""`                              //角色排序
	Remark    string     `json:"remark" gorm:"size:255;"`                       //备注
	Admin     bool       `json:"admin" gorm:"size:4;"`                          // 是否是admin
	DataScope string     `json:"dataScope" gorm:"size:128;"`                    // 数据权限
	Params    string     `json:"params" gorm:"-"`
	MenuIds   []int      `json:"menuIds" gorm:"-"`
	DeptIds   []int      `json:"deptIds" gorm:"-"`
	SysDept   []SysDept  `json:"sysDept" gorm:"many2many:sys_role_dept;foreignKey:RoleId;joinForeignKey:role_id;references:DeptId;joinReferences:dept_id;"`
	SysMenu   *[]SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
	common.ControlBy
	common.ModelTime
}

func (*SysRole) TableName() string {
	return "sys_role"
}

func (e *SysRole) Generate() common.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRole) GetId() interface{} {
	return e.RoleId
}

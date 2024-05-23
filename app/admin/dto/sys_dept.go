// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/12 17:22:00
// @Desc
package dto

import common "meng-admin-gin/common/model"

type SysDeptGetPageReq struct {
	DeptId   int    `form:"deptId" search:"type:exact;column:dept_id;table:sys_dept" comment:"id"`          //id
	ParentId int    `form:"parentId" search:"type:exact;column:parent_id;table:sys_dept" comment:"上级部门"`    //上级部门
	DeptPath string `form:"deptPath" search:"type:exact;column:dept_path;table:sys_dept" comment:""`        //路径
	DeptName string `form:"deptName" search:"type:contains;column:dept_name;table:sys_dept" comment:"部门名称"` //部门名称
	Sort     int    `form:"sort" search:"type:exact;column:sort;table:sys_dept" comment:"排序"`               //排序
	Leader   string `form:"leader" search:"type:contains;column:leader;table:sys_dept" comment:"负责人"`       //负责人
	Phone    string `form:"phone" search:"type:exact;column:phone;table:sys_dept" comment:"手机"`             //手机
	Email    string `form:"email" search:"type:exact;column:email;table:sys_dept" comment:"邮箱"`             //邮箱
	Status   string `form:"status" search:"type:exact;column:status;table:sys_dept" comment:"状态"`           //状态
}

func (m *SysDeptGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysDeptGetReq struct {
	Id int `uri:"id"`
}

func (s *SysDeptGetReq) GetId() interface{} {
	return s.Id
}

type SysDeptDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysDeptDeleteReq) GetId() interface{} {
	return s.Ids
}

type SysDeptInsertReq struct {
	DeptId   int    `uri:"id" comment:"编码"`                        // 编码
	ParentId int    `json:"parentId" comment:"上级部门" vd:"?"`        //上级部门
	DeptPath string `json:"deptPath" comment:""`                   //路径
	DeptName string `json:"deptName" comment:"部门名称" vd:"len($)>0"` //部门名称
	Sort     int    `json:"sort" comment:"排序" vd:"?"`              //排序
	Leader   string `json:"leader" comment:"负责人" vd:"?"`           //负责人
	Phone    string `json:"phone" comment:"手机" vd:"?"`             //手机
	Email    string `json:"email" comment:"邮箱" vd:"?"`             //邮箱
	Status   int    `json:"status" comment:"状态" vd:"$>0"`          //状态
	common.ControlBy
}

// GetId 获取数据对应的ID
func (s *SysDeptInsertReq) GetId() interface{} {
	return s.DeptId
}

type SysDeptUpdateReq struct {
	DeptId   int    `uri:"id" comment:"编码"`                        // 编码
	ParentId int    `json:"parentId" comment:"上级部门" vd:"?"`        //上级部门
	DeptPath string `json:"deptPath" comment:""`                   //路径
	DeptName string `json:"deptName" comment:"部门名称" vd:"len($)>0"` //部门名称
	Sort     int    `json:"sort" comment:"排序" vd:"?"`              //排序
	Leader   string `json:"leader" comment:"负责人"`                  //负责人
	Phone    string `json:"phone" comment:"手机" vd:"?"`             //手机
	Email    string `json:"email" comment:"邮箱" vd:"?"`             //邮箱
	Status   int    `json:"status" comment:"状态" vd:"$>0"`          //状态
	common.ControlBy
}

// GetId 获取数据对应的ID
func (s *SysDeptUpdateReq) GetId() interface{} {
	return s.DeptId
}

type DeptLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLabel `gorm:"-" json:"children"`
}

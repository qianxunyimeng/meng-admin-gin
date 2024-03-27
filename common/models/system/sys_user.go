// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 22:39:00
// @Desc
package system

import (
	"meng-admin-gin/common/models"
)

type SysUser struct {
	//gorm.Model
	UserId   int      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`                                   // 用户id
	Username string   `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password string   `json:"-" gorm:"comment:用户登录密码"`                                                              // 登录密码
	Salt     string   `json:"-" gorm:"size:255;comment:加盐"`                                                         // 加盐
	NickName string   `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Avatar   string   `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone    string   `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email    string   `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable   int      `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	DeptId   int      `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int      `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string   `json:"remark" gorm:"size:255;comment:备注"` //用户是否被冻结 1正常 2冻结
	DeptIds  []int    `json:"deptIds" gorm:"-"`
	PostIds  []int    `json:"postIds" gorm:"-"`
	RoleIds  []int    `json:"roleIds" gorm:"-"`
	Dept     *SysDept `json:"dept"`
	models.ControlBy
	models.ModelTime
}

func (SysUser) TableName() string {
	return "sys_users"
}

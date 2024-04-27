// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 22:39:00
// @Desc
package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	//gorm.Model
	UserId   int      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`                                          // 用户id
	Username string   `json:"userName" gorm:"index;comment:用户登录名"`                                                         // 用户登录名
	Password string   `json:"-" gorm:"comment:用户登录密码"`                                                                     // 登录密码
	Salt     string   `json:"-" gorm:"size:255;comment:加盐"`                                                                // 加盐
	NickName string   `json:"nickName" gorm:"default:系统用户;size:128;comment:用户昵称"`                                          // 用户昵称
	Sex      string   `json:"sex" gorm:"size:10;comment:性别"`                                                               // 性别
	Avatar   string   `json:"avatar" gorm:"default:https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp;comment:用户头像"` // 用户头像
	Phone    string   `json:"phone"  gorm:"size:11;comment:用户手机号"`                                                         // 用户手机号
	Email    string   `json:"email"  gorm:"size:128;comment:用户邮箱"`                                                         // 用户邮箱
	Status   int      `json:"status" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	RoleId   int      `json:"roleId" gorm:"size:20;comment:角色ID"`
	DeptId   int      `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int      `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string   `json:"remark" gorm:"size:255;comment:备注"` //用户是否被冻结 1正常 2冻结
	DeptIds  []int    `json:"deptIds" gorm:"-"`
	PostIds  []int    `json:"postIds" gorm:"-"`
	RoleIds  []int    `json:"roleIds" gorm:"-"`
	Dept     *SysDept `json:"dept" gorm:"-"`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_users"
}

// 在新增数据库之前对密码加密
func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

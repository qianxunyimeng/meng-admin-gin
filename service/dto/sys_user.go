// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/28 22:16:00
// @Desc
package dto

import (
	"meng-admin-gin/common/models"
	"meng-admin-gin/common/models/system"
)

type SysUserRegisterReq struct {
	UserId   int    `json:"userId" comment:"用户ID"` // 用户ID
	Username string `json:"userName" comment:"用户名" vd:"len($)>0"`
	Password string `json:"password" comment:"密码" vd:"len($)>0"`
	NickName string `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Phone    string `json:"phone" comment:"手机号" vd:"len($)>0"`
	RoleId   int    `json:"roleId" comment:"角色ID"`
	Avatar   string `json:"avatar" comment:"头像"`
	Sex      string `json:"sex" comment:"性别"`
	Email    string `json:"email" comment:"邮箱" vd:"len($)>0,email"`
	DeptId   int    `json:"deptId" comment:"部门" vd:"$>0"`
	PostId   int    `json:"postId" comment:"岗位"`
	Remark   string `json:"remark" comment:"备注"`
	Status   int    `json:"status" comment:"状态" default:1`
	models.ControlBy
}

type LoginReq struct {
	Username  string `form:"username" json:"username" binding:"required"`   // 用户名
	Password  string `form:"password" json:"password" binding:"required"`   // 密码
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`     // 验证码
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"` // 验证码ID
}

func (s *SysUserRegisterReq) Generate(model *system.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Username = s.Username
	model.Password = s.Password
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.RoleId = s.RoleId
	model.Avatar = s.Avatar
	model.Sex = s.Sex
	model.Email = s.Email
	model.DeptId = s.DeptId
	model.PostId = s.PostId
	model.Remark = s.Remark
	model.Status = s.Status
	model.CreateBy = s.CreateBy
}

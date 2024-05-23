// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:31:00
// @Desc
package dto

import "meng-admin-gin/app/admin/models"

type LoginResp struct {
	User      models.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type Login struct {
	Username  string `form:"username" json:"username" vd:"len($)>0; msg:'用户名不能为空'"`     // 用户名
	Password  string `form:"password" json:"password" vd:"len($)>0; msg:'密码不能为空'"`      // 密码
	Captcha   string `form:"captcha" json:"captcha" vd:"len($)>0; msg:'验证码不能为空'"`       // 验证码
	CaptchaId string `form:"captchaId" json:"captchaId" vd:"len($)>0; msg:'验证码ID不能为空'"` // 验证码ID
}

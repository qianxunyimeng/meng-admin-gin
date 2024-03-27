// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 19:48:00
// @Desc
package request

type Login struct {
	Username  string `form:"username" json:"username" binding:"required"`   // 用户名
	Password  string `form:"password" json:"password" binding:"required"`   // 密码
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`     // 验证码
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"` // 验证码ID
}

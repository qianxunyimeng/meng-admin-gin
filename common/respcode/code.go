// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 22:23:00
// @Desc
package respcode

const (
	Success = iota
	Error
	ErrorParam   // 参数错误
	UnAuthorized // 未授权
	FORBIDDEN    // 未登录 或 非法访问
)

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 20:44:00
// @Desc
package response

type Responses interface {
	SetCode(int32)
	SetTraceID(string) //用于链路追踪
	SetMsg(interface{})
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses // 初始化/重置
}

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 23:34:00
// @Desc
package model

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}

type Responses interface {
	SetCode(int32)
	SetTraceID(string) //用于链路追踪
	SetMsg(string)
	SetError(error)
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses // 初始化/重置
}

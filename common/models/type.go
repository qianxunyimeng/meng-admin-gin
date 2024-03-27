// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 23:34:00
// @Desc
package models

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}

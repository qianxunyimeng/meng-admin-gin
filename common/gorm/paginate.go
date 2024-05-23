// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/12 17:33:00
// @Desc
package gorm

import (
	"gorm.io/gorm"
)

func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}

		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

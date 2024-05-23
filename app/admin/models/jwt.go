// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/31 23:15:00
// @Desc
package models

import "gorm.io/gorm"

// jwt黑名单
type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

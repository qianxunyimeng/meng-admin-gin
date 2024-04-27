// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:31:00
// @Desc
package dto

import (
	"meng-admin-gin/common/models"
)

type LoginResp struct {
	User      models.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

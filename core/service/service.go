// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/28 23:34:00
// @Desc
package service

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"meng-admin-gin/core/storage"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *zap.Logger
	Error error
	Cache storage.AdapterCache
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/28 23:36:00
// @Desc
package system

import (
	"errors"
	"meng-admin-gin/common/dto"
	"meng-admin-gin/common/models"
	"meng-admin-gin/core/service"
)

type SysUserService struct {
	service.Service
}

func (s *SysUserService) Insert(d *dto.SysUserRegisterReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = s.Orm.Model(&data).Where("username = ?", d.Username).Count(&i).Error
	if err != nil {
		s.Log.Error(err.Error())
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		s.Log.Error(err.Error())
		return err
	}
	d.Generate(&data)
	err = s.Orm.Create(&data).Error
	if err != nil {
		s.Log.Error(err.Error())
		return err
	}
	return nil
}

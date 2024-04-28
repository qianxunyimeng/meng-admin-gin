// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:27:00
// @Desc
package service

import (
	"errors"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/model"
	"meng-admin-gin/common/service"
	"meng-admin-gin/utils"
)

type SysBaseService struct {
	service.Service
}

func (s *SysBaseService) Login(r *dto.Login) (userInter *model.SysUser, err error) {
	var user model.SysUser
	err = s.Orm.Where("username = ?", r.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(r.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

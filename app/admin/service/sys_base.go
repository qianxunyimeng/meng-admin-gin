// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:27:00
// @Desc
package service

import (
	"errors"
	"fmt"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/common/service"
	"meng-admin-gin/utils"
)

type SysBaseService struct {
	service.Service
}

func (s *SysBaseService) Login(r *dto.Login) (userInter *models.SysUser, err error) {
	var user models.SysUser
	err = s.Orm.Where("user_name = ?", r.Username).First(&user).Error

	if err == nil {
		if ok := utils.BcryptCheck(r.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	fmt.Println("roleId: ", user.RoleId)
	err = s.Orm.Table("sys_role").Where("role_id = ?", user.RoleId).Limit(1).Scan(&user.Role).Error
	if err == nil {
		if ok := utils.BcryptCheck(r.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:27:00
// @Desc
package system

import (
	"errors"
	"fmt"
	"meng-admin-gin/common/models/system"
	"meng-admin-gin/core/service"
	"meng-admin-gin/model/base/request"
	"meng-admin-gin/utils"
)

type SysBaseService struct {
	service.Service
}

func (s *SysBaseService) Login(r *request.Login) (userInter *system.SysUser, err error) {
	var user system.SysUser
	err = s.Orm.Where("username = ?", r.Username).First(&user).Error
	fmt.Println(err)
	if err == nil {
		if ok := utils.BcryptCheck(r.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

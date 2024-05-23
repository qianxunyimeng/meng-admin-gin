// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/28 23:36:00
// @Desc
package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	cDto "meng-admin-gin/common/dto"
	cGorm "meng-admin-gin/common/gorm"
	"meng-admin-gin/common/service"
	"meng-admin-gin/utils"
)

type SysUserService struct {
	service.Service
}

func (s *SysUserService) Insert(d *dto.SysUserRegisterReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = s.Orm.Model(&data).Where("user_name = ?", d.UserName).Count(&i).Error
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

func (s *SysUserService) GetById(id int, result *models.SysUser) error {
	var data models.SysUser
	fmt.Println(id, result)
	err := s.Orm.Model(&data).
		First(result, id).Error

	if err != nil {
		s.Log.Error(err.Error())
		return err
	}
	return nil
}

// GetPage 获取SysUser列表
func (e *SysUserService) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	var err error
	var data models.SysUser

	err = e.Orm.Debug().Model(&data).
		Preload("Dept").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cGorm.Paginate(c.GetPageNum(), c.GetPageSize()),
			//actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Error(err.Error())
		return err
	}
	return nil
}

// Get 获取SysUser对象
func (e *SysUserService) Get(d *dto.SysUserById, model *models.SysUser) error {
	var data models.SysUser

	err := e.Orm.Model(&data).Debug().
		//Scopes(
		//	actions.Permission(data.TableName(), p),
		//).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Error("db error: " + err.Error())
		return err
	}
	if err != nil {
		e.Log.Error("db error: " + err.Error())
		return err
	}
	return nil
}

// Update 修改SysUser对象
func (e *SysUserService) Update(c *dto.SysUserUpdateReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Error("Service UpdateSysUser error: " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	//c.Generate(&model)
	copier.Copy(&model, c)
	update := e.Orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Error("db error: " + err.Error())
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		e.Log.Warn("db update error")
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUserService) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.
		//Scopes(
		//actions.Permission(model.TableName(), p),
		//)
		First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Error("Service UpdateSysUser error: " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		e.Log.Error("Service UpdateSysUser error: " + err.Error())
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUserService) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.
		//	Scopes(
		//	actions.Permission(model.TableName(), p),
		//).
		First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Error("At Service ResetSysUserPwd error: " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	//c.Generate(&model)
	copier.Copy(&model, c)
	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		e.Log.Error("At Service ResetSysUserPwd error: " + err.Error())
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUserService) Remove(c *dto.SysUserById) error {
	var err error
	var data models.SysUser

	db := e.Orm.Model(&data).
		//Scopes(
		//	actions.Permission(data.TableName(), p),
		//).
		Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Error("Error found in  RemoveSysUser : " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUserService) UpdatePwd(id int, oldPassword, newPassword string) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.Orm.Model(c).
		//Scopes(
		//	actions.Permission(c.TableName(), p),
		//).
		Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		e.Log.Error("db error: " + err.Error())
		return err
	}
	//var ok bool
	//ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
	//ok = utils.BcryptCheck(oldPassword, c.Password)
	//if err != nil {
	//	e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
	//	return err
	//}
	if ok := utils.BcryptCheck(oldPassword, c.Password); !ok {
		e.Log.Error("旧密码错误")
		return errors.New("旧密码错误")
	}
	//if !ok {
	//	err = errors.New("incorrect Password")
	//	e.Log.Warnf("user[%d] %s", id, err.Error())
	//	return err
	//}
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		e.Log.Error("db error: " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		e.Log.Warn("db update error")
		return err
	}
	return nil
}

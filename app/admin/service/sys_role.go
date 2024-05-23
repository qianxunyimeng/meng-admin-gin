// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/15 21:49:00
// @Desc
package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	cDto "meng-admin-gin/common/dto"
	cGorm "meng-admin-gin/common/gorm"
	"meng-admin-gin/common/service"
)

type SysRoleService struct {
	service.Service
}

// GetPage 获取SysRole列表
func (e *SysRoleService) GetPage(c *dto.SysRoleGetPageReq, list *[]models.SysRole, count *int64) error {
	var err error
	var data models.SysRole

	err = e.Orm.Model(&data).
		//.Preload("SysMenu").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.OrderDest("sort", false),
			cGorm.Paginate(c.GetPageNum(), c.GetPageSize()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		return err
	}
	return nil
}

// Get 获取SysRole对象
func (e *SysRoleService) Get(d *dto.SysRoleGetReq, model *models.SysRole) error {
	var err error
	//var data models.SysRole
	db := e.Orm.First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Error(err.Error())
		return err
	}
	if err != nil {
		e.Log.Error(err.Error())
		return err
	}
	menuIds := make([]int, 0)
	query := e.Orm.Table("sys_menu").Select("sys_menu.parent_id").Joins("INNER JOIN sys_role_menu rm ON sys_menu.menu_id = rm.menu_id AND rm.role_id = ?", d.GetId())
	err = e.Orm.Table("sys_menu").Select("sys_menu.menu_id").Joins("LEFT JOIN sys_role_menu rm ON sys_menu.menu_id = rm.menu_id").Where("rm.role_id = ?", d.GetId()).Where("sys_menu.menu_id not in (?)", query).Order("sys_menu.parent_id,sys_menu.sort").Scan(&menuIds).Error
	//model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		e.Log.Error(err.Error())
		return err
	}

	model.MenuIds = menuIds

	return nil
}

// Insert 创建SysRole对象
func (e *SysRoleService) Insert(c *dto.SysRoleInsertReq) error {
	var err error
	var data models.SysRole

	copier.Copy(&data, c)

	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var count int64
	err = tx.Model(&data).Where("role_code = ?", c.RoleCode).Count(&count).Error
	if err != nil {
		e.Log.Error(err.Error())
		return err
	}

	if count > 0 {
		err = errors.New("roleCode已存在，需更换在提交！")
		e.Log.Error(err.Error())
		return err
	}

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Error(err.Error())
		return err
	}

	return nil
}

// Update 修改SysRole对象
func (e *SysRoleService) Update(c *dto.SysRoleUpdateReq) error {
	var err error
	var data models.SysRole
	if c.RoleCode == "admin" {
		return errors.New("admin角色不能修改")
	}
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	tx.First(&data, c.GetId())
	copier.Copy(&data, c)

	// 更新关联的数据，使用 FullSaveAssociations 模式
	//db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&data)
	db := tx.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Error(err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	return nil
}

// Remove 删除SysRole
func (e *SysRoleService) Remove(c *dto.SysRoleDeleteReq) error {
	var err error

	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.Preload("SysMenu").Preload("SysDept").First(&model, c.GetId())
	if model.RoleCode == "admin" {
		return errors.New("admin角色不能修改")
	}
	//删除 SysRole 时，同时删除角色所有 关联其它表 记录 (SysMenu 和 SysMenu)
	db := tx.Select(clause.Associations).Delete(&model)

	if err = db.Error; err != nil {
		e.Log.Error(err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	// 清除 sys_casbin_rule 权限表里 当前角色的所有记录
	//_, _ = cb.RemoveFilteredPolicy(0, model.RoleKey)

	return nil
}

// UpdateStatus 修改SysRole对象status
func (e *SysRoleService) UpdateStatus(c *dto.UpdateStatusReq) error {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.First(&model, c.GetId())
	if model.RoleCode == "admin" {
		return errors.New("admin角色不能修改")
	}
	copier.Copy(&model, c)
	// 更新关联的数据，使用 FullSaveAssociations 模式
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if err = db.Error; err != nil {
		e.Log.Error(err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

func (e *SysRoleService) UpdateRoleMenu(c *dto.UpdateMenuAuthReq) error {
	var model = models.SysRole{}
	var mlist = make([]models.SysMenu, 0)
	var err error

	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	tx.Preload("SysMenu").First(&model, c.GetId())
	tx.Where("menu_id in ?", c.MenuIds).Find(&mlist)
	err = tx.Model(&model).Association("SysMenu").Delete(model.SysMenu)
	if err != nil {
		e.Log.Error("delete policy error: " + err.Error())
		return err
	}
	model.SysMenu = &mlist
	// 更新关联的数据，使用 FullSaveAssociations 模式
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if err = db.Error; err != nil {
		e.Log.Error("db error: " + err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return err
}

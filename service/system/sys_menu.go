// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 17:46:00
// @Desc
package system

import (
	"errors"
	"gorm.io/gorm"
	"meng-admin-gin/common/dto"
	"meng-admin-gin/common/models"
	"meng-admin-gin/core/service"
	"strconv"
)

type SysMenuService struct {
	service.Service
}

// Insert 创建SysMenu对象
func (e *SysMenuService) Insert(c *dto.SysMenuInsertReq) *SysMenuService {
	var err error
	var data models.SysMenu
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if !errors.Is(e.Orm.Where("menu_name = ? and path = ?", c.MenuName, c.Path).First(&models.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		//return errors.New("存在重复name，请修改name")
		_ = e.AddError(errors.New("存在重复的menu_name和path，请修改菜单名称和路由地址"))
		return e
	}

	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		e.Log.Error(err.Error())
		_ = e.AddError(err)
	}
	c.MenuId = data.MenuId
	err = e.initPaths(tx, &data)
	if err != nil {
		tx.Rollback()
		e.Log.Error(err.Error())
		_ = e.AddError(err)
	}
	tx.Commit()
	return e
}

func (e *SysMenuService) initPaths(tx *gorm.DB, menu *models.SysMenu) error {
	var err error
	var data models.SysMenu
	parentMenu := new(models.SysMenu)
	if menu.ParentId != 0 {
		err = tx.Model(&data).First(parentMenu, menu.ParentId).Error
		if err != nil {
			return err
		}
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return err
		}
		menu.Paths = parentMenu.Paths + "/" + strconv.Itoa(menu.MenuId)
	} else {
		menu.Paths = "/0/" + strconv.Itoa(menu.MenuId)
	}
	err = tx.Model(&data).Where("menu_id = ?", menu.MenuId).Update("paths", menu.Paths).Error
	return err
}

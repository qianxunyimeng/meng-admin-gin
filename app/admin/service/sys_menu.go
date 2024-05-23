// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 17:46:00
// @Desc
package service

import (
	"errors"
	"gorm.io/gorm"
	"meng-admin-gin/app/admin/Vo"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	cDto "meng-admin-gin/common/dto"
	cModels "meng-admin-gin/common/model"
	"meng-admin-gin/common/service"
	"meng-admin-gin/utils"
	"sort"
	"strconv"
)

type SysMenuService struct {
	service.Service
}

func (e *SysMenuService) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) *SysMenuService {

	var err error
	var data models.SysMenu

	err = e.Orm.Model(&data).
		Scopes(
			cDto.OrderDest("sort", false),
			cDto.MakeCondition(c.GetNeedSearch()),
		).Preload("SysApi").
		Find(menus).Error
	if err != nil {
		e.Log.Error(err.Error())
		_ = e.AddError(err)
		return e
	}
	return e
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
		_ = e.AddError(errors.New("当前层级存在相同的菜单，请修改菜单名称和路由地址"))
		return e
	}

	//if !e.checkMenuNameUniqueOfParent(c) {
	//	_ = e.AddError(errors.New("当前层级存在相同的菜单，请修改菜单名称"))
	//	return e
	//}

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

func (e *SysMenuService) checkMenuNameUniqueOfParent(c *dto.SysMenuInsertReq) bool {
	if !errors.Is(e.Orm.Where("menu_name = ? and parent_id = ?", c.MenuName, c.ParentId).First(&models.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (e *SysMenuService) GetMenuById(d *dto.SysMenuGetReq, result *models.SysMenu) *SysMenuService {
	var err error
	var data models.SysMenu

	db := e.Orm.Model(&data).
		First(result, d.GetId())
	err = db.Error
	if err != nil {
		e.Log.Error(err.Error())
		_ = e.AddError(err)
		return e
	}

	return e
}

// Insert 创建SysMenu对象
func (e *SysMenuService) Update(c *dto.SysMenuUpdatetReq) *SysMenuService {
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

	if !errors.Is(e.Orm.Where("menu_name = ? and path = ? and menu_id != ?", c.MenuName, c.Path, c.MenuId).First(&models.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		_ = e.AddError(errors.New("当前层级存在相同的菜单，请修改菜单名称和路由地址"))
		return e
	}
	var model = models.SysMenu{}
	// 查询当前最新的信息
	tx.First(&model, c.GetId())
	//oldPath := model.Paths
	c.Generate(&model)
	//fmt.Println("paths:", model.Paths)
	db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if err = db.Error; err != nil {
		e.Log.Error(err.Error())
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}
	tx.Commit()
	return e
}

// Remove 删除SysMenu
func (e *SysMenuService) Remove(d *dto.SysMenuDeleteReq) *SysMenuService {
	var err error
	var data models.SysMenu

	db := e.Orm.Model(&data).Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Error(err.Error())
		_ = e.AddError(err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		_ = e.AddError(err)
	}
	return e
}

// GetRouterByUserId 获取左侧菜单树使用
func (e *SysMenuService) GetRouterByUserId(roleCode string) (m []models.SysMenu, err error) {
	data := make([]models.SysMenu, 0)
	var role models.SysRole
	if utils.IsAdminOfRoleCode(roleCode) {
		err = e.Orm.Where(" menu_type in ('M','C') and deleted_at is null").
			Order("sort").
			Find(&data).
			Error
	} else {
		err = e.Orm.Model(&role).Where("role_code = ? ", roleCode).Preload("SysMenu").First(&role).Error
		if role.SysMenu != nil {
			for _, menu := range *role.SysMenu {
				data = append(data, menu)
			}
		}
		sort.Sort(models.SysMenuSlice(data))
	}
	m = make([]models.SysMenu, 0)
	for i := 0; i < len(data); i++ {
		if data[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&data, data[i])
		m = append(m, menusInfo)
	}
	return
}
func (e *SysMenuService) getByRoleName(roleName string) ([]models.SysMenu, error) {
	var err error
	data := make([]models.SysMenu, 0)
	if roleName == "admin" {
		err = e.Orm.Where(" menu_type in ('M','C') and deleted_at is null").
			Order("sort").
			Find(&data).
			Error
	}
	sort.Sort(models.SysMenuSlice(data))
	return data, err
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.ViewType = list[j].ViewType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.SysApi = list[j].SysApi
		mi.Children = []models.SysMenu{}

		if mi.MenuType != cModels.BUTTON {
			ms := menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

// BuildMenus 构造前端需要的路由树
func (e *SysMenuService) BuildMenus(menus []models.SysMenu) ([]vo.RouterVo, error) {
	//var err error
	//data := make([]vo.RouterVo, 0)

	//for i := 0; i < len(menus); i++ {
	//	router := new(vo.RouterVo)
	//	router.SetPath(getRouterPath(menus[i]))
	//	router.SetHandle(vo.RouterHandleVo{
	//		Component: getRouterComponent(menus[i]),
	//		Title:     menus[i].MenuName,
	//		Icon:      menus[i].Icon,
	//	})
	//
	//	child := menus[i].Children
	//	if child != nil && len(child) > 0 && cModels.DIRECTORY == menus[i].MenuType {
	//		router.SetChildren(BUild(child))
	//	}
	//}
	return buildMenu(menus, nil)
}

func buildMenu(menus []models.SysMenu, parent *vo.RouterVo) ([]vo.RouterVo, error) {
	data := make([]vo.RouterVo, 0)

	for i := 0; i < len(menus); i++ {
		router := vo.RouterVo{}
		router.SetPath(getRouterPath(menus[i], parent))
		router.SetHandle(vo.RouterHandleVo{
			Component:   getRouterComponent(menus[i]),
			Title:       menus[i].MenuName,
			Icon:        menus[i].Icon,
			IsKeepAlive: !menus[i].NoCache,
		})

		child := menus[i].Children
		if child != nil && len(child) > 0 && cModels.DIRECTORY == menus[i].MenuType {
			newChild, err := buildMenu(child, &router)
			if err != nil {
				return data, err
			}
			router.SetChildren(newChild)
		}

		data = append(data, router)
	}

	return data, nil
}

func getRouterPath(menu models.SysMenu, parent *vo.RouterVo) string {
	var path = menu.Path
	//if menu.ParentId == 0 && cModels.DIRECTORY == menu.MenuType {
	//	path = "/" + menu.Path
	//
	//	if parent != nil {
	//		path = parent.Path + path
	//	}
	//}
	path = "/" + menu.Path
	if parent != nil {
		path = parent.Path + path
	}

	return path
}

// getRouterComponent 获取菜单组件信息
func getRouterComponent(menu models.SysMenu) string {
	var component = cModels.LAYOUT
	//if menu.Component != "" && !isMenuFrame(menu) {
	//	component = menu.Component
	//}
	if menu.Component != "" && menu.ViewType == cModels.NORMAL_PAGE {
		component = menu.Component
	}

	return component
}

// 是否为菜单内部跳转
func isMenuFrame(menu models.SysMenu) bool {
	return menu.ParentId == 0 && cModels.MENU == menu.MenuType
}

package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
	"errors"
)

type MenuService struct {}

//
// MenuTreeDtoList
//  @Description: 获取菜单树列表，支持查找
//  @receiver a
//  @param menuVo
//  @return []*response.MenuTreeDto
//  @return error
//
func (a *MenuService) MenuTreeDtoList(menuVo request.MenuVo) ([]*response.MenuTreeDto, error) {
	// 动态拼接查询条件
	tx := db
	if menuVo.Code != "" {
		tx = tx.Where("code LIKE ?", "%" + menuVo.Code + "%")
	}
	if menuVo.Name != "" {
		tx = tx.Where("name LIKE ?", "%" + menuVo.Name + "%")
	}
	if menuVo.Path != "" {
		tx = tx.Where("path LIKE ?", "%" + menuVo.Path + "%")
	}
	// 查询完整数据
	var menuList []*model.Menu
	err := tx.Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	// 获取顶层菜单列表
	levelTopList := a.getLevelTopList(menuList)
	// 根据父级id，汇总子菜单列表
	childrenMap := a.aggregateChildren(menuList)
	// 将所有菜单列表生成为树形结构
	var menuTreeDtoList []*response.MenuTreeDto
	for _, item := range levelTopList {
		// 初始化树形节点，并进行属性赋值
		var menuTreeDto response.MenuTreeDto
		err = utils.CopyFields(&menuTreeDto, item)
		if err != nil {
			return nil, err
		}
		// 添加对应的子节点列表
		childrenList := childrenMap[menuTreeDto.ID]
		if childrenList != nil {
			for _, child := range childrenList {
				var childTreeDto response.MenuTreeDto
				err = utils.CopyFields(&childTreeDto, child)
				if err != nil {
					return nil, err
				}
				menuTreeDto.Children = append(menuTreeDto.Children, &childTreeDto)
			}
		}
		menuTreeDtoList = append(menuTreeDtoList, &menuTreeDto)
		// 添加完成后，在map中删除该项
		delete(childrenMap, menuTreeDto.ID)
	}
	// 如还存在子节点未处理，则统一追加
	if len(childrenMap) != 0 {
		for _, restChildrenList := range childrenMap {
			for _, restChild := range restChildrenList {
				var restChildTreeDto response.MenuTreeDto
				err = utils.CopyFields(&restChildTreeDto, restChild)
				if err != nil {
					return nil, err
				}
				menuTreeDtoList = append(menuTreeDtoList, &restChildTreeDto)
			}
		}
	}

	return menuTreeDtoList, nil
}

//
// MenuTreeOptionDtoList
//  @Description: 获取资源树形选项列表
//  @receiver a
//  @return []*response.MenuTreeOptionDto
//  @return error
//
func (a *MenuService) MenuTreeOptionDtoList() ([]*response.MenuTreeOptionDto, error) {
	// 查询完整数据
	var menuList []*model.Menu
	err := db.Select("id, name, parent_id").Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	// 获取顶层菜单列表
	levelTopList := a.getLevelTopList(menuList)
	// 根据父级id，汇总子菜单列表
	childrenMap := a.aggregateChildren(menuList)
	// 将所有菜单列表生成为树形结构
	var menuTreeOptionDtoList []*response.MenuTreeOptionDto
	for _, item := range levelTopList {
		// 初始化树形节点，并进行属性赋值
		var menuTreeOptionDto response.MenuTreeOptionDto
		err = utils.CopyFields(&menuTreeOptionDto, item)
		if err != nil {
			return nil, err
		}
		// 添加对应的子节点列表
		childrenList := childrenMap[menuTreeOptionDto.ID]
		if childrenList != nil {
			for _, child := range childrenList {
				var childTreeOptionDto response.MenuTreeOptionDto
				err = utils.CopyFields(&childTreeOptionDto, child)
				if err != nil {
					return nil, err
				}
				menuTreeOptionDto.Children = append(menuTreeOptionDto.Children, &childTreeOptionDto)
			}
		}
		menuTreeOptionDtoList = append(menuTreeOptionDtoList, &menuTreeOptionDto)
	}

	return menuTreeOptionDtoList, nil
}

//
// UserMenuTreeDtoList
//  @Description: 获取用户菜单树列表，本质上通过用户对应的角色列表，查找与之绑定的菜单
//  @receiver a
//  @param id
//  @return []*response.UserMenuTreeDto
//  @return error
//
func (a *MenuService) UserMenuTreeDtoList(userId int) ([]*response.UserMenuTreeDto, error) {
	// 查询用户所属角色对应的所有菜单列表，要求菜单未被禁用
	// 如父节点被禁用，而子节点未被禁用，则子节点无法获取
	// 但是，不建议如此设置，请优先禁用子节点
	var menuList []*model.Menu
	tx := db
	tx = tx.Where("menu.id in ( " +
			"select distinct rrm.menu_id " +
			"from relation_role_menu rrm " +
			"where rrm.role_id in ( " +
				"select distinct rur.role_id " +
				"from relation_user_role rur " +
				"where rur.user_id = ? " +
			")" +
		")", userId)
	tx = tx.Where("deleted = ?", model.MODEL_ACTIVED)
	err := tx.Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	// 获取顶层菜单列表
	levelTopList := a.getLevelTopList(menuList)
	// 根据父级id，汇总子菜单列表
	childrenMap := a.aggregateChildren(menuList)
	// 将所有菜单列表生成为树形结构
	var userMenuTreeDtoList []*response.UserMenuTreeDto
	for _, item := range levelTopList {
		// 初始化树形节点，并进行属性赋值
		var userMenuTreeDto response.UserMenuTreeDto
		err = utils.CopyFields(&userMenuTreeDto, item)
		if err != nil {
			return nil, err
		}
		// 添加对应的子节点列表
		childrenList := childrenMap[userMenuTreeDto.ID]
		if childrenList != nil {
			for _, child := range childrenList {
				var childTreeDto response.UserMenuTreeDto
				err = utils.CopyFields(&childTreeDto, child)
				if err != nil {
					return nil, err
				}
				userMenuTreeDto.Children = append(userMenuTreeDto.Children, &childTreeDto)
			}
		}
		userMenuTreeDtoList = append(userMenuTreeDtoList, &userMenuTreeDto)
	}

	return userMenuTreeDtoList, nil
}

//
// SaveMenuItem
//  @Description: 添加菜单项
//  @receiver a
//  @param menuSaveVo
//  @return bool
//  @return error
//
func (a *MenuService) SaveMenuItem(menuSaveVo request.MenuSaveVo) (bool, error) {
	// 构造菜单
	var menu model.Menu
	err := utils.CopyFields(&menu, menuSaveVo)
	if err != nil {
		return false, err
	}
	// 数据合法性检查
	if menu.Level > model.MENU_LEVEL_TOP && menu.ParentId == nil {
		return false, errors.New("非顶层菜单项必须有父级！")
	}
	menu.Type = model.MENU_TYPE_ITEM
	if menu.Hidden == nil {
		hidden := model.MENU_HIDDEN_DISABLE
		menu.Hidden = &hidden
	}
	// 新增菜单
	err = db.Create(&menu).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// SaveMenuGroup
//  @Description: 添加菜单组
//  @receiver a
//  @param menuSaveVo
//  @return bool
//  @return error
//
func (a *MenuService) SaveMenuGroup(menuSaveVo request.MenuSaveVo) (bool, error) {
	// 构造菜单
	var menu model.Menu
	err := utils.CopyFields(&menu, menuSaveVo)
	if err != nil {
		return false, err
	}
	// 数据合法性检查
	menu.Type = model.MENU_TYPE_GROUP
	menu.Level = model.MENU_LEVEL_TOP
	menu.ParentId = nil
	if menu.Hidden == nil {
		hidden := model.MENU_HIDDEN_DISABLE
		menu.Hidden = &hidden
	}

	// 新增菜单
	err = db.Create(&menu).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// UpdateMenu
//  @Description: 更新菜单，但无法更新类型与层级
//  @receiver a
//  @param menuUpdateVo
//  @return bool
//  @return error
//
func (a *MenuService) UpdateMenu(menuUpdateVo request.MenuUpdateVo) (bool, error) {
	var menu model.Menu
	err := utils.CopyFields(&menu, menuUpdateVo)
	menu.Type = 0
	if err != nil {
		return false, err
	}
	err = db.Model(&menu).Updates(menu).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *MenuService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var menu model.Menu
	menu.ID = modelDeletedVo.ID
	err := db.Model(&menu).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *MenuService) RemoveMenuById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Menu{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// getLevelTopList
//  @Description: 获取给定菜单列表中的顶层部分
//  @receiver a
//  @param menuList
//  @return []*model.Menu
//
func (a *MenuService) getLevelTopList(menuList []*model.Menu) []*model.Menu {
	var levelTopList []*model.Menu
	for _, item := range menuList {
		if item.ParentId == nil {
			levelTopList = append(levelTopList, item)
		}
	}
	return levelTopList
}

//
// aggregateChildren
//  @Description: 汇总子菜单那列表，以父级id为索引
//  @receiver a
//  @param menuList
//  @return map[int][]*model.Menu
//
func (a *MenuService) aggregateChildren(menuList []*model.Menu) map[int][]*model.Menu {
	childrenMap := map[int][]*model.Menu{}
	for _, item := range menuList {
		if item.ParentId == nil {
			continue
		}

		childrenMap[*item.ParentId] = append(childrenMap[*item.ParentId], item)
	}
	return childrenMap
}
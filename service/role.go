package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

type RoleService struct {}

//
// RoleAdminDtoPage
//  @Description: 分页获取后台角色
//  @receiver a
//  @param pageVo
//  @param roleVo
//  @return *response.Page
//  @return error
//
func (a *RoleService) RoleAdminDtoPage(pageVo request.PageVo, roleVo request.RoleVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载，其实只需要加载id
	tx := db.Preload("MenuList").Preload("PermissionList")
	// 动态拼接查询条件
	if roleVo.Code != "" {
		tx = tx.Where("code LIKE ?", "%" + roleVo.Code + "%")
	}
	if roleVo.Name != "" {
		tx = tx.Where("name LIKE ?", "%" + roleVo.Name + "%")
	}
	if roleVo.Description != "" {
		tx = tx.Where("description LIKE ?", "%" + roleVo.Description + "%")
	}

	// 查询分页信息
	var roleList []*model.Role
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&roleList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var roleAdminDtoList []*response.RoleAdminDto
	for _, item := range roleList {
		var roleAdmin *response.RoleAdminDto
		roleAdmin, err = item.CopyToRoleAdminDto()
		roleAdminDtoList = append(roleAdminDtoList, roleAdmin)
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(roleAdminDtoList)),
		Total:   count,
		Records: roleAdminDtoList,
	}
	return &page, nil
}

//
// RoleOptionDtoList
//  @Description: 获取全部角色选项列表
//  @receiver a
//  @return []*response.RoleOptionDto
//  @return error
//
func (a *RoleService) RoleOptionDtoList() ([]*response.RoleOptionDto, error) {
	var roleList []*model.Role
	err := db.Select("id, code, name").Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	var roleOptionDtoList []*response.RoleOptionDto
	for _, item := range roleList {
		var roleOption response.RoleOptionDto
		err = utils.CopyFields(&roleOption, item)
		if err != nil {
			return nil, err
		}
		roleOptionDtoList = append(roleOptionDtoList, &roleOption)
	}
	
	return roleOptionDtoList, nil
}

func (a *RoleService) RoleDashboardDtoList() ([]*response.RoleDashboardDto, error) {
	var roleDashboardDtoList []*response.RoleDashboardDto
	err := db.Table("role").
		Select("id, name, count(*) as count").
		Joins("left join relation_user_role on role.id=relation_user_role.role_id").
		Group("id").
		Scan(&roleDashboardDtoList).Error
	if err != nil {
		return nil, err
	}
	return roleDashboardDtoList, nil
}

//
// RemoveRoleById
//  @Description: 通过id删除角色
//  @receiver a
//  @param id
//  @return bool
//  @return error
//
func (a *RoleService) RemoveRoleById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Role{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *RoleService) SaveRole(roleSaveVo request.RoleSaveVo) (bool, error) {
	// 构造角色
	var role model.Role
	err := utils.CopyFields(&role, roleSaveVo)
	if err != nil {
		return false, err
	}
	// 新增角色
	err = db.Create(&role).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *RoleService) UpdateRole(roleUpdateVo request.RoleUpdateVo) (bool, error) {
	var role model.Role
	err := utils.CopyFields(&role, roleUpdateVo)
	if err != nil {
		return false, err
	}

	err = db.Model(&role).Updates(role).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *RoleService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var role model.Role
	role.ID = modelDeletedVo.ID
	err := db.Model(&role).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// UpdateRoleMenu
//  @Description: 更新角色分配菜单列表
//  @receiver a
//  @param roleMenuVo
//  @return bool
//  @return error
//
func (a *RoleService) UpdateRoleMenu(roleMenuVo request.RoleMenuVo) (bool, error) {
	var role model.Role
	role.ID = roleMenuVo.ID
	var err error
	// 删除原有菜单列表
	tx := db.Begin()
	err = tx.Model(&role).Association("MenuList").Clear().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	// 插入新菜单列表
	if len(roleMenuVo.MenuIdList) != 0 {
		var menuList []model.Menu
		for _, menuId := range roleMenuVo.MenuIdList {
			menu := model.Menu{Model: model.Model{ ID: menuId }}
			menuList = append(menuList, menu)
		}
		err = tx.Model(&role).Association("MenuList").Append(menuList).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

//
// UpdateRolePermission
//  @Description: 更新角色分配权限列表
//  @receiver a
//  @param rolePermissionVo
//  @return bool
//  @return error
//
func (a *RoleService) UpdateRolePermission(rolePermissionVo request.RolePermissionVo) (bool, error) {
	var role model.Role
	role.ID = rolePermissionVo.ID
	var err error
	// 删除原有权限列表
	tx := db.Begin()
	err = tx.Model(&role).Association("PermissionList").Clear().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	// 插入新权限列表
	if len(rolePermissionVo.PermissionIdList) != 0 {
		var permissionList []model.Permission
		for _, permissionId := range rolePermissionVo.PermissionIdList {
			permission := model.Permission{Model: model.Model{ ID: permissionId }}
			permissionList = append(permissionList, permission)
		}
		err = tx.Model(&role).Association("PermissionList").Append(permissionList).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}


package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
	"errors"
)

type PermissionService struct {}

//
// PermissionRoleDtoList
//  @Description: 获取权限-角色关联对象列表，用于授权
//  @receiver a
//  @return []*response.PermissionRoleDto
//  @return error
//
func (a *PermissionService) PermissionRoleDtoList() ([]*response.PermissionRoleDto, error) {
	var permissionList []*model.Permission
	tx := db.Preload("RoleList")
	// 仅查询具体权限，忽略权限组，忽略禁用权限
	tx = tx.Where("type = ? AND url IS NOT NULL", model.PERMISSION_TYPE_ITEM)
	tx = tx.Where("deleted = ?", model.MODEL_ACTIVED)
	err := tx.Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	var permissionRoleDtoList []*response.PermissionRoleDto
	for _, permission := range permissionList {
		var permissionRoleDto *response.PermissionRoleDto
		permissionRoleDto, err = permission.CopyToPermissionRoleDto()
		if err != nil {
			return nil, err
		}
		permissionRoleDtoList = append(permissionRoleDtoList, permissionRoleDto)
	}
	return permissionRoleDtoList, nil
}

//
// PermissionTreeDtoList
//  @Description: 获取资源树形结构列表，支持查找
//  @receiver a
//  @param permissionVo
//  @return []*response.PermissionTreeDto
//  @return error
//
func (a *PermissionService) PermissionTreeDtoList(permissionVo request.PermissionVo) ([]*response.PermissionTreeDto, error) {
	// 动态拼接查询条件
	tx := db
	if permissionVo.Url != nil && *permissionVo.Url != "" {
		tx = tx.Where("url LIKE ?", "%" + *permissionVo.Url + "%")
	}
	if permissionVo.Name != "" {
		tx = tx.Where("name LIKE ?", "%" + permissionVo.Name + "%")
	}
	// 查询完整数据
	var permissionList []*model.Permission
	err := tx.Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	// 获取顶层权限列表
	levelTopList := a.getLevelTopList(permissionList)
	// 根据父级id，汇总子权限列表
	childrenMap := a.aggregateChildren(permissionList)
	// 将所有权限列表生成为树形结构
	var permissionTreeDtoList []*response.PermissionTreeDto
	for _, item := range levelTopList {
		// 初始化树形节点，并进行属性赋值
		var permissionTreeDto response.PermissionTreeDto
		err = utils.CopyFields(&permissionTreeDto, item)
		if err != nil {
			return nil, err
		}
		// 添加对应的子节点列表
		childrenList := childrenMap[permissionTreeDto.ID]
		if childrenList != nil {
			for _, child := range childrenList {
				var childTreeDto response.PermissionTreeDto
				err = utils.CopyFields(&childTreeDto, child)
				if err != nil {
					return nil, err
				}
				permissionTreeDto.Children = append(permissionTreeDto.Children, &childTreeDto)
			}
		}
		permissionTreeDtoList = append(permissionTreeDtoList, &permissionTreeDto)
		// 添加完成后，在map中删除该项
		delete(childrenMap, permissionTreeDto.ID)
	}
	// 如还存在子节点未处理，则统一追加
	if len(childrenMap) != 0 {
		for _, restChildrenList := range childrenMap {
			for _, restChild := range restChildrenList {
				var restChildTreeDto response.PermissionTreeDto
				err = utils.CopyFields(&restChildTreeDto, restChild)
				if err != nil {
					return nil, err
				}
				permissionTreeDtoList = append(permissionTreeDtoList, &restChildTreeDto)
			}
		}
	}

	return permissionTreeDtoList, nil
}

//
// PermissionTreeOptionDtoList
//  @Description: 获取资源树形选项列表
//  @receiver a
//  @return []*response.PermissionTreeOptionDto
//  @return error
//
func (a *PermissionService) PermissionTreeOptionDtoList() ([]*response.PermissionTreeOptionDto, error) {
	// 查询完整数据
	var permissionList []*model.Permission
	err := db.Select("id, name, parent_id").Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	// 获取顶层权限列表
	levelTopList := a.getLevelTopList(permissionList)
	// 根据父级id，汇总子权限列表
	childrenMap := a.aggregateChildren(permissionList)
	// 将所有权限列表生成为树形结构
	var permissionTreeOptionDtoList []*response.PermissionTreeOptionDto
	for _, item := range levelTopList {
		// 初始化树形节点，并进行属性赋值
		var permissionTreeOptionDto response.PermissionTreeOptionDto
		err = utils.CopyFields(&permissionTreeOptionDto, item)
		if err != nil {
			return nil, err
		}
		// 添加对应的子节点列表
		childrenList := childrenMap[permissionTreeOptionDto.ID]
		if childrenList != nil {
			for _, child := range childrenList {
				var childTreeOptionDto response.PermissionTreeOptionDto
				err = utils.CopyFields(&childTreeOptionDto, child)
				if err != nil {
					return nil, err
				}
				permissionTreeOptionDto.Children = append(permissionTreeOptionDto.Children, &childTreeOptionDto)
			}
		}
		permissionTreeOptionDtoList = append(permissionTreeOptionDtoList, &permissionTreeOptionDto)
	}

	return permissionTreeOptionDtoList, nil
}

//
// SavePermissionItem
//  @Description: 添加权限项
//  @receiver a
//  @param permissionSaveVo
//  @return bool
//  @return error
//
func (a *PermissionService) SavePermissionItem(permissionSaveVo request.PermissionSaveVo) (bool, error) {
	// 构造标签
	var permission model.Permission
	err := utils.CopyFields(&permission, permissionSaveVo)
	if err != nil {
		return false, err
	}
	// 数据合法性检查
	if permissionSaveVo.Url == nil || *permissionSaveVo.Url == "" {
		return false, errors.New("权限项的路径不能为空！")
	}
	if permissionSaveVo.Level == model.PERMISSION_LEVEL_TOP && permissionSaveVo.ParentId != nil {
		return false, errors.New("顶级权限没有父级！")
	}
	if permissionSaveVo.Level != model.PERMISSION_LEVEL_TOP &&
		(permissionSaveVo.ParentId == nil || *permissionSaveVo.ParentId <= 0) {
		return false, errors.New("非顶级权限应有父级！")
	}
	permission.Type = model.PERMISSION_TYPE_ITEM
	if permission.Anonymous == nil {
		anonymous := model.PERMISSION_ANONYMOUS_DISABLE
		permission.Anonymous = &anonymous
	}
	// 新增标签
	err = db.Create(&permission).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// SavePermissionGroup
//  @Description: 添加权限组
//  @receiver a
//  @param permissionSaveVo
//  @return bool
//  @return error
//
func (a *PermissionService) SavePermissionGroup(permissionSaveVo request.PermissionSaveVo) (bool, error) {
	// 构造标签
	var permission model.Permission
	err := utils.CopyFields(&permission, permissionSaveVo)
	if err != nil {
		return false, err
	}
	// 数据合法性检查
	permission.Type = model.PERMISSION_TYPE_GROUP
	permission.Level = model.PERMISSION_LEVEL_TOP
	permission.ParentId = nil
	if permission.Anonymous == nil {
		anonymous := model.PERMISSION_ANONYMOUS_DISABLE
		permission.Anonymous = &anonymous
	}

	// 新增标签
	err = db.Create(&permission).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// UpdatePermissionItem
//  @Description: 更新权限项，允许更新多个字段
//  @receiver a
//  @param permissionUpdateVo
//  @return bool
//  @return error
//
func (a *PermissionService) UpdatePermissionItem(permissionUpdateVo request.PermissionUpdateVo) (bool, error) {
	var permission model.Permission
	err := utils.CopyFields(&permission, permissionUpdateVo)
	if err != nil {
		return false, err
	}

	err = db.Model(&permission).Where("type = ?", model.PERMISSION_TYPE_ITEM).Updates(permission).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// UpdatePermissionGroup
//  @Description: 更新权限组，只允许更新名称
//  @receiver a
//  @param permissionUpdateVo
//  @return bool
//  @return error
//
func (a *PermissionService) UpdatePermissionGroup(permissionUpdateVo request.PermissionUpdateVo) (bool, error) {
	var permission model.Permission
	permission.ID = permissionUpdateVo.ID
	permission.Name = permissionUpdateVo.Name

	err := db.Model(&permission).Where("type = ?", model.PERMISSION_TYPE_GROUP).Updates(permission).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *PermissionService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var permission model.Permission
	permission.ID = modelDeletedVo.ID
	err := db.Model(&permission).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *PermissionService) RemovePermissionById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Permission{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// getLevelTopList
//  @Description: 获取给定权限列表中的顶层部分
//  @receiver a
//  @param permissionList
//  @return []*model.Permission
//
func (a *PermissionService) getLevelTopList(permissionList []*model.Permission) []*model.Permission {
	var levelTopList []*model.Permission
	for _, item := range permissionList {
		if item.ParentId == nil {
			levelTopList = append(levelTopList, item)
		}
	}
	return levelTopList
}

//
// aggregateChildren
//  @Description: 汇总子权限列表，以父级id为索引
//  @receiver a
//  @param permissionList
//  @return map[int][]*model.Permission
//
func (a *PermissionService) aggregateChildren(permissionList []*model.Permission) map[int][]*model.Permission {
	childrenMap := map[int][]*model.Permission{}
	for _, item := range permissionList {
		if item.ParentId == nil {
			continue
		}

		childrenMap[*item.ParentId] = append(childrenMap[*item.ParentId], item)
	}
	return childrenMap
}
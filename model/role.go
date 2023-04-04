package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// Role
//  @Description: 角色
//
type Role struct {
	Model
	// 角色代码
	Code                string              `json:"code"            gorm:"unique;not null"`
	// 角色名
	Name			    string		        `json:"name"            gorm:"not null;size:255"`
	// 角色描述
	Description		    string		        `json:"description"     gorm:"size:255"`
	// 菜单列表，多对多关联
	MenuList            []*Menu             `json:"menuList"        gorm:"many2many:relation_role_menu"`
	// 权限列表，多对多关联
	PermissionList      []*Permission       `json:"permissionList"  gorm:"many2many:relation_role_permission"`
}

const (
	ROLE_ADMIN_ID    int8 = 1
	ROLE_TEST_ID     int8 = 2
	ROLE_USER_ID     int8 = 3
)

func (a *Role) CopyToRoleAdminDto() (*response.RoleAdminDto, error) {
	var roleAdminDto response.RoleAdminDto
	err := utils.CopyFields(&roleAdminDto, a)
	if err != nil {
		return nil, err
	}
	for _, item := range a.MenuList {
		roleAdminDto.MenuIdList = append(roleAdminDto.MenuIdList, item.ID)
	}
	for _, item := range a.PermissionList {
		roleAdminDto.PermissionIdList = append(roleAdminDto.PermissionIdList, item.ID)
	}
	return &roleAdminDto, nil
}





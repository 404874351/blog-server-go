package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// Permission
//  @Description: 权限
//
type Permission struct {
	Model
	// 权限路径，权限组通常为空
	Url			    *string		`json:"url"         gorm:"unique;size:255"`
	// 权限名称
	Name		    string		`json:"name"        gorm:"not null;size:255"`
	// 权限类型，0具体权限，1权限组，默认0
	Type		    int8        `json:"type"        gorm:"not null"`
	// 权限层级，0顶层，正数代表具体层级，默认0
	Level		    int8		`json:"level"       gorm:"not null"`
	// 父级id，null没有父级，即处于顶层
	ParentId		*int		`json:"parentId"    gorm:""`
	// 是否支持匿名访问，0否，1是，默认0
	Anonymous		*int8		`json:"anonymous"   gorm:"not null"`
	// 角色列表，多对多关联
	RoleList        []*Role     `json:"roleList"    gorm:"many2many:relation_role_permission"`
}

const (
	PERMISSION_TYPE_ITEM           int8 = 0
	PERMISSION_TYPE_GROUP          int8 = 1
	PERMISSION_LEVEL_TOP           int8 = 0
	PERMISSION_ANONYMOUS_DISABLE   int8 = 0
	PERMISSION_ANONYMOUS_ENABLE    int8 = 1
)

//
// CopyToPermissionRoleDto
//  @Description: 复制Permission到PermissionRoleDto
//  @receiver a
//  @return *response.PermissionRoleDto
//  @return error
//
func (a *Permission) CopyToPermissionRoleDto() (*response.PermissionRoleDto, error) {
	var permissionRole response.PermissionRoleDto
	err := utils.CopyFields(&permissionRole, a)
	if err != nil {
		return nil, err
	}
	for _, item := range a.RoleList {
		var roleOption response.RoleOptionDto
		err = utils.CopyFields(&roleOption, item)
		if err != nil {
			return nil, err
		}
		permissionRole.RoleList = append(permissionRole.RoleList, &roleOption)
	}
	return &permissionRole, nil
}




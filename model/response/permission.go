package response

import (
	"blog-server-go/utils"
)

//
// PermissionTreeDto
//  @Description: 权限树，用于列表展示
//
type PermissionTreeDto struct {
	// id
	ID        			int 					`json:"id"`
	// 权限路径
	Url					*string					`json:"url" `
	// 权限名
	Name				string					`json:"name"`
	// 权限类型
	Type		    	int8        			`json:"type"`
	// 是否支持匿名访问
	Anonymous			*int8					`json:"anonymous"`
	// 创建时间
	CreateTime 			utils.SystemTime		`json:"createTime"`
	// 数据禁用
	Deleted 			int8 					`json:"deleted"`
	// 子权限列表
	Children			[]*PermissionTreeDto 	`json:"children"`
}

//
// PermissionTreeOptionDto
//  @Description: 权限树选项，用于快速增减
//
type PermissionTreeOptionDto struct {
	// id
	ID        			int 						`json:"id"`
	// 权限名
	Name				string						`json:"name"`
	// 子权限列表
	Children			[]*PermissionTreeOptionDto 	`json:"children" `
}

//
// PermissionRoleDto
//  @Description: 权限与关联角色列表，用于授权
//
type PermissionRoleDto struct {
	// id
	ID        			int 					`json:"id"`
	// 权限路径
	Url					*string					`json:"url" `
	// 是否支持匿名访问
	Anonymous			*int8					`json:"anonymous"  `
	// 数据禁用
	Deleted 			int8 					`json:"deleted" `
	// 角色列表
	RoleList			[]*RoleOptionDto		`json:"roleList" `
}


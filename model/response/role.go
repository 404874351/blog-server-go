package response

import (
	"blog-server-go/utils"
)

//
// RoleOptionDto
//  @Description: 角色选项，用于快速增减
//
type RoleOptionDto struct {
	// id
	ID        			int 		        `json:"id"`
	// 角色代码
	Code				string		        `json:"code" `
	// 角色名
	Name    			string		        `json:"name"`
}

//
// RoleDashboardDto
//  @Description: 仪表盘角色信息
//
type RoleDashboardDto struct {
	// id
	ID        			int 		        `json:"id"`
	// 角色名
	Name    			string		        `json:"name"`
	// 计数
	Count				int			        `json:"count"`
}

//
// RoleAdminDto
//  @Description: 后台角色
//
type RoleAdminDto struct {
	// id
	ID        			int 		        `json:"id"`
	// 角色代码
	Code				string		        `json:"code" `
	// 角色名
	Name				string		        `json:"name"`
	// 角色描述
	Description			string		        `json:"description"`
	// 创建时间
	CreateTime 			utils.SystemTime	`json:"createTime"`
	// 数据禁用
	Deleted 			int8 		        `json:"deleted" `
	// 权限id列表
	PermissionIdList    []int    	        `json:"permissionIdList"`
	// 菜单id列表
	MenuIdList    		[]int    	        `json:"menuIdList"`
}

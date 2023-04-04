package response

import (
	"blog-server-go/utils"
)

//
// MenuTreeDto
//  @Description: 菜单树，用于列表展示
//
type MenuTreeDto struct {
	// id
	ID        			int 					`json:"id"`
	// 菜单代码
	Code				string					`json:"code" `
	// 菜单名
	Name				string					`json:"name"`
	// 菜单路径
	Path		    	string					`json:"path"    `
	// 菜单组件
	Component			string					`json:"component"  `
	// 菜单图标
	Icon				string					`json:"icon"    `
	// 菜单类型
	Type		   	 	int8        			`json:"type"      `
	// 父级id
	ParentId			*int		    		`json:"parentId" `
	// 是否隐藏
	Hidden				*int8					`json:"hidden"  `
	// 创建时间
	CreateTime 			utils.SystemTime		`json:"createTime"`
	// 数据禁用
	Deleted 			int8 					`json:"deleted" `
	// 子菜单列表
	Children			[]*MenuTreeDto 			`json:"children" `
}

//
// MenuTreeOptionDto
//  @Description: 菜单树选项，用于快速增减
//
type MenuTreeOptionDto struct {
	// id
	ID        			int 						`json:"id"`
	// 菜单名
	Name				string						`json:"name"`
	// 子菜单列表
	Children			[]*MenuTreeOptionDto 		`json:"children" `
}

//
// UserMenuTreeDto
//  @Description: 用户菜单树，用于后台动态菜单配置
//
type UserMenuTreeDto struct {
	// id
	ID        			int 					`json:"id"`
	// 菜单代码
	Code				string					`json:"code" `
	// 菜单名
	Name				string					`json:"name"`
	// 菜单路径
	Path		    	string					`json:"path"    `
	// 菜单组件
	Component			string					`json:"component"  `
	// 菜单图标
	Icon				string					`json:"icon"    `
	// 是否隐藏
	Hidden				*int8					`json:"hidden"  `
	// 子菜单列表
	Children			[]*UserMenuTreeDto 		`json:"children" `
}
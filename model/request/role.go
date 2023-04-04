package request

//
// RoleVo
//  @Description: 角色 请求对象
//
type RoleVo struct {
	// 角色代码
	Code				string		`json:"code"                form:"code"`
	// 角色名
	Name				string		`json:"name"                form:"name"`
	// 角色描述
	Description			string		`json:"description"         form:"description"`
}

//
// RoleUpdateVo
//  @Description: 角色更新 请求对象
//
type RoleUpdateVo struct {
	// id
	ID        			int			`json:"id"                  form:"id"`
	// 角色代码
	Code				string		`json:"code"                form:"code"`
	// 角色名
	Name				string		`json:"name"                form:"name"`
	// 角色描述
	Description			string		`json:"description"         form:"description"`
}

//
// RoleSaveVo
//  @Description: 角色新增 请求对象
//
type RoleSaveVo struct {
	// 角色代码
	Code				string		`json:"code"                form:"code"                 binding:"required"`
	// 角色名
	Name				string		`json:"name"                form:"name"                 binding:"required"`
	// 角色描述
	Description			string		`json:"description"         form:"description"`
}

//
// RoleMenuVo
//  @Description: 角色分配菜单 请求对象
//
type RoleMenuVo struct {
	// id
	ID        			int			`json:"id"                  form:"id"`
	// 菜单id列表
	MenuIdList    		[]int    	`json:"menuIdList"          form:"menuIdList"           binding:"required"`
}

type RolePermissionVo struct {
	// id
	ID        			int			`json:"id"                  form:"id"`
	// 权限id列表
	PermissionIdList 	[]int    	`json:"permissionIdList"    form:"permissionIdList"     binding:"required"`
}

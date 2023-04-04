package request

//
// PermissionVo
//  @Description: 权限 请求对象
//
type PermissionVo struct {
	// 权限路径
	Url			    *string		`json:"url"         form:"url"`
	// 权限名称
	Name		    string		`json:"name"        form:"name"`
}

//
// PermissionUpdateVo
//  @Description: 权限更新 请求对象
//
type PermissionUpdateVo struct {
	// id
	ID        		int			`json:"id"          form:"id"`
	// 权限路径
	Url			    *string		`json:"url"         form:"url"`
	// 权限名称
	Name		    string		`json:"name"        form:"name"`
	// 权限类型
	Type		    int8        `json:"type"        form:"type"         binding:"ValidatePermissionType"`
	// 是否支持匿名访问
	Anonymous		*int8		`json:"anonymous"   form:"anonymous"    binding:"omitempty,ValidatePermissionAnonymous"`
}

type PermissionSaveVo struct {
	// 权限路径
	Url			    *string		`json:"url"         form:"url"`
	// 权限名称
	Name		    string		`json:"name"        form:"name"         binding:"required"`
	// 权限类型
	Type		    int8        `json:"type"        form:"type"         binding:"ValidatePermissionType"`
	// 权限层级
	Level		    int8		`json:"level"       form:"level"        binding:"ValidatePermissionLevel"`
	// 父级id
	ParentId		*int		`json:"parentId"    form:"parentId"`
	// 是否支持匿名访问
	Anonymous		*int8		`json:"anonymous"   form:"anonymous"    binding:"omitempty,ValidatePermissionAnonymous"`
}
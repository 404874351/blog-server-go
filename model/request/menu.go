package request

//
// MenuVo
//  @Description: 菜单 请求对象
//
type MenuVo struct {
	// 菜单代码
	Code			string		`json:"code"        form:"code"`
	// 菜单名称
	Name		    string		`json:"name"        form:"name"`
	// 菜单路径
	Path		    string		`json:"path"        form:"path"`
}

//
// MenuUpdateVo
//  @Description: 菜单更新 请求对象
//
type MenuUpdateVo struct {
	// id
	ID        		int			`json:"id"          form:"id"`
	// 菜单代码
	Code			string		`json:"code"        form:"code"`
	// 菜单名称
	Name		    string		`json:"name"        form:"name"`
	// 菜单路径
	Path		    string		`json:"path"        form:"path"`
	// 菜单组件
	Component		string		`json:"component"   form:"component"`
	// 菜单图标
	Icon			string		`json:"icon"        form:"icon"`
	// 菜单类型
	Type		    int8        `json:"type"        form:"type"             binding:"ValidateMenuType"`
	// 是否隐藏
	Hidden			*int8		`json:"hidden"      form:"hidden"           binding:"omitempty,ValidateMenuHidden"`
}

type MenuSaveVo struct {
	// 菜单代码
	Code			string		`json:"code"        form:"code"             binding:"required"`
	// 菜单名称
	Name		    string		`json:"name"        form:"name"             binding:"required"`
	// 菜单路径
	Path		    string		`json:"path"        form:"path"`
	// 菜单组件
	Component		string		`json:"component"   form:"component"        binding:"required"`
	// 菜单图标
	Icon			string		`json:"icon"        form:"icon"`
	// 菜单类型
	Type		    int8        `json:"type"        form:"type"             binding:"ValidateMenuType"`
	// 菜单层级
	Level		    int8		`json:"level"       form:"level"            binding:"omitempty,ValidateMenuLevel"`
	// 父级id
	ParentId		*int		`json:"parentId"    form:"parentId"`
	// 是否隐藏
	Hidden			*int8		`json:"hidden"      form:"hidden"           binding:"omitempty,ValidateMenuHidden"`
}

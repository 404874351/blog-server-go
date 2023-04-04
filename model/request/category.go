package request

//
// CategoryVo
//  @Description: 分类 请求对象
//
type CategoryVo struct {
	// 分类名称
	Name			string		`json:"name"        form:"name"`
}

//
// CategoryUpdateVo
//  @Description: 分类更新 请求对象
//
type CategoryUpdateVo struct {
	// id
	ID        	    int		    `json:"id"          form:"id"`
	// 分类名称
	Name			string		`json:"name"        form:"name"         binding:"required"`
}

//
// CategorySaveVo
//  @Description: 分类新增 请求对象
//
type CategorySaveVo struct {
	// 分类名称
	Name			string		`json:"name"        form:"name"         binding:"required"`
}


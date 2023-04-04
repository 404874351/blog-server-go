package request

//
// TagVo
//  @Description: 标签 请求对象
//
type TagVo struct {
	// 标签名称
	Name			string		`json:"name"        form:"name"`
}

//
// TagUpdateVo
//  @Description: 标签更新 请求对象
//
type TagUpdateVo struct {
	// id
	ID        	    int		    `json:"id"          form:"id"`
	// 标签名称
	Name			string		`json:"name"        form:"name"         binding:"required"`
}

//
// TagSaveVo
//  @Description: 标签新增 请求对象
//
type TagSaveVo struct {
	// 标签名称
	Name			string		`json:"name"        form:"name"         binding:"required"`
}
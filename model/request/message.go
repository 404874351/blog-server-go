package request

//
// MessageVo
//  @Description: 留言 请求对象
//
type MessageVo struct {
	// 内容
	Content			string		`json:"content"        form:"content"`
	// 用户昵称
	Nickname		string		`json:"nickname"       form:"nickname"`
}

//
// MessageSaveVo
//  @Description: 新增留言 请求对象
//
type MessageSaveVo struct {
	// 内容
	Content			string		`json:"content"        form:"content"         binding:"required"`
	// 用户id
	UserId		    *int 		`json:"userId"         form:"userId"`
}

package request

//
// CommentVo
//  @Description: 评论 前台请求对象
//
type CommentVo struct {
	// 用户id，从登录态获取，也可为空
	UserId		    int		    `json:"userId"          form:"userId"`
	// 文章id
	ArticleId		int		    `json:"articleId"       form:"articleId"        binding:"required"`
	// 父评论id
	ParentId		int		    `json:"parentId"        form:"parentId"`
	// 排序字段
	SortBy			string		`json:"sortBy"          form:"sortBy"`
}

//
// CommentAdminVo
//  @Description: 评论 后台请求对象
//
type CommentAdminVo struct {
	// 内容
	Content			string		`json:"content"         form:"content"`
	// 用户昵称
	Nickname		string		`json:"nickname"        form:"nickname"`
	// 文章标题
	ArticleTitle	string		`json:"articleTitle"    form:"articleTitle"`
	// 置顶
	Top		        *int8		`json:"top"             form:"top"`
}

//
// CommentUpdateVo
//  @Description: 评论更新 请求对象
//
type CommentUpdateVo struct {
	// id
	ID        	    int		    `json:"id"              form:"id"`
	// 文章id，用于查验信息，不作为更新内容
	ArticleId		int		    `json:"articleId"       form:"articleId"        binding:"required"`
	// 置顶
	Top		        *int8		`json:"top"             form:"top"              binding:"omitempty,ValidateCommentTop"`
}

//
// CommentSaveVo
//  @Description: 评论新增 请求对象
//
type CommentSaveVo struct {
	// 内容
	Content			string		`json:"content"         form:"content"          binding:"required"`
	// 文章id
	ArticleId		int		    `json:"articleId"       form:"articleId"        binding:"required"`
	// 用户id，从登录态获取
	UserId		    int		    `json:"userId"          form:"userId"`
	// 父评论id，顶级评论为空
	ParentId		*int		`json:"parentId"        form:"parentId"`
	// 回复用户id，顶级评论为空
	ReplyUserId		*int	 	`json:"replyUserId"     form:"replyUserId"`
}

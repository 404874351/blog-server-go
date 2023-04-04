package request

//
// ArticleVo
//  @Description: 文章 前台请求对象
//
type ArticleVo struct {
	// 搜索关键词
	Key			    string		`json:"key"             form:"key"`
	// 分类id
	CategoryId		int		    `json:"categoryId"      form:"categoryId"`
	// 降序字段
	SortBy			string		`json:"sortBy"          form:"sortBy"`
}

//
// ArticleAdminVo
//  @Description: 文章 后台请求对象
//
type ArticleAdminVo struct {
	// 标题
	Title			string		`json:"title"           form:"title"`
	// 作者昵称
	Nickname		string		`json:"nickname"        form:"nickname"`
	// 分类id
	CategoryId		int		    `json:"categoryId"      form:"categoryId"`
	// 标签id
	TagId			int		    `json:"tagId"           form:"tagId"`
}

//
// ArticleUpdateVo
//  @Description: 分类更新 请求对象
//
type ArticleUpdateVo struct {
	// id
	ID        	    int		    `json:"id"              form:"id"`
	// 标题
	Title			string		`json:"title"           form:"title"`
	// 简介
	Description		string		`json:"description"     form:"description"`
	// 标题
	CoverUrl		string		`json:"coverUrl"        form:"coverUrl"`
	// 分类id
	CategoryId		int		    `json:"categoryId"      form:"categoryId"`
	// 标签id列表
	TagIdList    	[]int    	`json:"tagIdList"       form:"tagIdList"`
	// 置顶
	Top		        *int8		`json:"top"             form:"top"                  binding:"omitempty,ValidateArticleTop"`
	// 置顶
	CloseComment    *int8		`json:"closeComment"    form:"closeComment"         binding:"omitempty,ValidateArticleCloseComment"`
}

//
// ArticleSaveVo
//  @Description: 文章新增 请求对象
//
type ArticleSaveVo struct {
	// 标题
	Title			string		`json:"title"           form:"title"                binding:"required"`
	// 简介
	Description		string		`json:"description"     form:"description"`
	// 标题
	CoverUrl		string		`json:"coverUrl"        form:"coverUrl"`
	// 原文链接
	ContentUrl		string		`json:"contentUrl"      form:"contentUrl"           binding:"required"`
	// 作者id
	UserId		    int		    `json:"userId"          form:"userId"               binding:"required"`
	// 分类id
	CategoryId		int		    `json:"categoryId"      form:"categoryId"           binding:"required"`
	// 标签id列表
	TagIdList    	[]int    	`json:"tagIdList"       form:"tagIdList"`
}
package response

import "blog-server-go/utils"

//
// ArticleAdminDto
//  @Description: 后台文章
//
type ArticleAdminDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标题
	Title		    string				`json:"title"`
	// 封面图链接
	CoverUrl		string				`json:"coverUrl"`
	// 原文链接
	ContentUrl      string				`json:"contentUrl"`
	// 浏览量
	ViewCount       int                 `json:"viewCount"`
	// 点赞量
	PraiseCount     int                 `json:"praiseCount"`
	// 评论量
	CommentCount    int                 `json:"commentCount"`
	// 作者
	User            *UserBaseInfoDto    `json:"user"`
	// 分类
	Category        *CategoryOptionDto  `json:"category"`
	// 标签列表
	TagList    	    []*TagOptionDto    	`json:"tagList"`
	// 置顶
	Top		        *int8		        `json:"top"`
	// 置顶
	CloseComment	*int8		        `json:"closeComment"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted"`
}

//
// ArticleDto
//  @Description: 前台文章
//
type ArticleDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标题
	Title		    string				`json:"title"`
	// 标题
	Description		string				`json:"description"`
	// 封面图链接
	CoverUrl		string				`json:"coverUrl"`
	// 原文链接
	ContentUrl      string				`json:"contentUrl"`
	// 浏览量
	ViewCount       int                 `json:"viewCount"`
	// 浏览量
	PraiseCount     int                 `json:"praiseCount"`
	// 浏览量
	CommentCount    int                 `json:"commentCount"`
	// 作者
	User            *UserBaseInfoDto    `json:"user"`
	// 分类
	Category        *CategoryOptionDto  `json:"category"`
	// 标签列表
	TagList    	    []*TagOptionDto    	`json:"tagList"`
	// 置顶
	Top		        *int8		        `json:"top"`
	// 置顶
	CloseComment	*int8		        `json:"closeComment"`
	// 是否已点赞
	PraiseStatus    bool                `json:"praiseStatus"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
}

//
// ArticleUpdateDto
//  @Description: 后台文章更新的展示数据
//
type ArticleUpdateDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标题
	Title		    string				`json:"title"`
	// 标题
	Description		string				`json:"description"`
	// 封面图链接
	CoverUrl		string				`json:"coverUrl"`
	// 原文链接
	ContentUrl      string				`json:"contentUrl"`
	// 分类
	Category        *CategoryOptionDto  `json:"category"`
	// 标签列表
	TagList    	    []*TagOptionDto    	`json:"tagList"`
	// 置顶
	Top		        *int8		        `json:"top"`
	// 置顶
	CloseComment	*int8		        `json:"closeComment"`
}

//
// ArticleDashboardDto
//  @Description: 仪表盘文章信息
//
type ArticleDashboardDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标题
	Title		    string				`json:"title"`
	// 浏览量
	ViewCount       int                 `json:"viewCount"`
}

//
// ArticleExtra
//  @Description: ArticleDto的附加信息，单独抽离便于业务处理
//
type ArticleExtra struct {
	// id
	ID        		int                 `json:"id"`
	// 点赞量
	PraiseCount     int                 `json:"praiseCount"`
	// 评论量
	CommentCount    int                 `json:"commentCount"`
	// 是否已点赞
	PraiseStatus    bool                `json:"praiseStatus"`
}
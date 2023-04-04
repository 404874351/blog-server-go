package response

import "blog-server-go/utils"

//
// CommentAdminDto
//  @Description: 后台评论
//
type CommentAdminDto struct {
	// id
	ID        		int                 `json:"id"`
	// 内容
	Content		    string				`json:"content"`
	// 置顶
	Top		        *int8		        `json:"top"`
	// 文章id
	ArticleId		int		            `json:"articleId"`
	// 文章标题
	ArticleTitle	string		        `json:"articleTitle"`
	// 用户基本信息
	User            *UserBaseInfoDto    `json:"user"`
	// 回复用户基本信息
	ReplyUser       *UserBaseInfoDto    `json:"replyUser"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted"`
}

//
// CommentDto
//  @Description: 前台评论
//
type CommentDto struct {
	// id
	ID        		int                 `json:"id"`
	// 内容
	Content		    string				`json:"content"`
	// 父评论id
	ParentId		*int		        `json:"parentId"`
	// 置顶
	Top		        *int8		        `json:"top"`
	// 用户基本信息
	User            *UserBaseInfoDto    `json:"user"`
	// 回复用户基本信息
	ReplyUser       *UserBaseInfoDto    `json:"replyUser"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 点赞数
	PraiseCount     int64               `json:"praiseCount"`
	// 是否已点赞
	PraiseStatus    bool                `json:"praiseStatus"`
	// 子评论总数
	ChildrenCount   int64               `json:"childrenCount"`
	// 子评论部分列表，默认最多加载3条
	Children        []*CommentDto       `json:"children"`
}

//
// CommentExtra
//  @Description: CommentDto的附加信息，单独抽离便于业务处理
//
type CommentExtra struct {
	// id
	ID        		int                 `json:"id"`
	// 点赞数
	PraiseCount     int64               `json:"praiseCount"`
	// 是否已点赞
	PraiseStatus    bool                `json:"praiseStatus"`
	// 子评论总数
	ChildrenCount   int64               `json:"childrenCount"`
	// 子评论部分列表，默认最多加载3条
	Children        []*CommentDto       `json:"children"`
}

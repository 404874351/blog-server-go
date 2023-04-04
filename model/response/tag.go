package response

import "blog-server-go/utils"

//
// TagDto
//  @Description: 文章标签
//
type TagDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标签名称
	Name			string		        `json:"name"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted"`
	// 标签下的文章数
	ArticleCount    int                 `json:"articleCount"`
}

//
// TagOptionDto
//  @Description: 文章标签选项
//
type TagOptionDto struct {
	// id
	ID        		int                 `json:"id"`
	// 标签名称
	Name			string		        `json:"name"`
}

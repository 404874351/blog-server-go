package response

import "blog-server-go/utils"

//
// CategoryDto
//  @Description: 文章分类
//
type CategoryDto struct {
	// id
	ID        		int                 `json:"id"`
	// 分类名称
	Name			string		        `json:"name"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted"`
	// 分类下的文章数
	ArticleCount    int                 `json:"articleCount"`
}

//
// CategoryOptionDto
//  @Description: 文章分类选项
//
type CategoryOptionDto struct {
	// id
	ID        		int                 `json:"id"`
	// 分类名称
	Name			string		        `json:"name"`
}

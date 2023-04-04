package response

import "blog-server-go/utils"

//
// MessageDto
//  @Description: 留言
//
type MessageDto struct {
	// id
	ID        		int                 `json:"id"`
	// 内容
	Content			string		        `json:"content"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 用户基本信息
	User            *UserBaseInfoDto    `json:"user"`
}

//
// MessageAdminDto
//  @Description: 后台留言
//
type MessageAdminDto struct {
	// id
	ID        		int                 `json:"id"`
	// 内容
	Content			string		        `json:"content"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted"`
	// 用户基本信息
	User            *UserBaseInfoDto    `json:"user"`
}

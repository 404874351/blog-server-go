package response

import (
	"blog-server-go/utils"
)

//
// UserInfoDto
//  @Description: 用户信息
//
type UserInfoDto struct {
	// id
	ID        		int                 `json:"id"`
	// 昵称
	Nickname		string		        `json:"nickname"`
	// 头像
	AvatarUrl		string		        `json:"avatarUrl"`
	// 用户对应的角色列表
	RoleList        []*RoleOptionDto    `json:"roleList"`
}

//
// UserBaseInfoDto
//  @Description: 用户基本信息，常用于展示
//
type UserBaseInfoDto struct {
	// id
	ID        		int                 `json:"id"`
	// 昵称
	Nickname		string		        `json:"nickname"`
	// 头像
	AvatarUrl		string		        `json:"avatarUrl"`
}

//
// UserAdminDto
//  @Description: 后台用户
//
type UserAdminDto struct {
	// id
	ID        		int                 `json:"id"`
	// 角色名
	Username		string				`json:"username"`
	// 昵称
	Nickname		string				`json:"nickname"`
	// 头像
	AvatarUrl		string				`json:"avatarUrl"`
	// 手机号
	Phone			string				`json:"phone"`
	// 创建时间
	CreateTime 		utils.SystemTime	`json:"createTime"`
	// 逻辑删除 数据禁用标记
	Deleted 		int8 				`json:"deleted" `
	// 用户对应的角色列表
	RoleList        []*RoleOptionDto    `json:"roleList"`
}
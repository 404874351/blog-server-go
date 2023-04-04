package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// User
//  @Description: 用户
//
type User struct {
	Model
	// 角色名
	Username		string		`json:"username"  gorm:"unique;not null;size:255"`
	// 密码
	Password		string		`json:"password"  gorm:"not null;size:255"`
	// 昵称
	Nickname		string		`json:"nickname"  gorm:"size:255"`
	// 头像
	AvatarUrl		string		`json:"avatarUrl" gorm:"size:255"`
	// 手机号
	Phone			string		`json:"phone"     gorm:"size:11"`
	// 角色列表，多对多关联
	RoleList        []*Role     `json:"roleList"  gorm:"many2many:relation_user_role"`
}

//
// CopyToUserInfoDto
//  @Description: 复制User到UserInfoDto
//  @receiver a
//  @return *response.UserInfoDto
//  @return error
//
func (a *User) CopyToUserInfoDto() (*response.UserInfoDto, error) {
	var userInfo response.UserInfoDto
	err := utils.CopyFields(&userInfo, a)
	if err != nil {
		return nil, err
	}
	for _, item := range a.RoleList {
		var roleOption response.RoleOptionDto
		err = utils.CopyFields(&roleOption, item)
		if err != nil {
			return nil, err
		}
		userInfo.RoleList = append(userInfo.RoleList, &roleOption)
	}
	return &userInfo, nil
}

//
// CopyToUserAdminDto
//  @Description: 复制User到UserAdminDto
//  @receiver a
//  @return *response.UserAdminDto
//  @return error
//
func (a *User) CopyToUserAdminDto() (*response.UserAdminDto, error) {
	var userAdmin response.UserAdminDto
	err := utils.CopyFields(&userAdmin, a)
	if err != nil {
		return nil, err
	}
	for _, item := range a.RoleList {
		var roleOption response.RoleOptionDto
		err = utils.CopyFields(&roleOption, item)
		if err != nil {
			return nil, err
		}
		userAdmin.RoleList = append(userAdmin.RoleList, &roleOption)
	}
	return &userAdmin, nil
}

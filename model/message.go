package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// Message
//  @Description: 留言
//
type Message struct {
	Model
	// 留言内容
	Content			string		`json:"content"         gorm:"not null;size:1023"`
	// 留言用户id
	UserId			*int		`json:"userId"          gorm:""`
	// 留言用户，游客为空
	User			*User       `json:"user"            gorm:"foreignkey:UserId"`
}

func (a *Message) CopyToMessageAdminDto() (*response.MessageAdminDto, error) {
	var messageAdminDto response.MessageAdminDto
	err := utils.CopyFields(&messageAdminDto, a)
	if err != nil {
		return nil, err
	}
	if a.User != nil {
		messageAdminDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(messageAdminDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	return &messageAdminDto, nil
}

func (a *Message) CopyToMessageDto() (*response.MessageDto, error) {
	var messageDto response.MessageDto
	err := utils.CopyFields(&messageDto, a)
	if err != nil {
		return nil, err
	}
	if a.User != nil {
		messageDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(messageDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	return &messageDto, nil
}

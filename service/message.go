package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

type MessageService struct {}

func (a *MessageService) MessageAdminDtoPage(pageVo request.PageVo, messageVo request.MessageVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载用户
	tx := db.Preload("User")
	// 设置动态查询条件
	if messageVo.Content != "" {
		tx = tx.Where("content LIKE ?", "%" + messageVo.Content + "%")
	}
	if messageVo.Nickname != "" {
		var userIdList []int
		err := db.Table("user").
			Select("id").
			Where("nickname LIKE ?", "%" + messageVo.Nickname + "%").
			Pluck("id", &userIdList).Error
		if err != nil {
			return nil, err
		}
		tx = tx.Where("user_id in (?)", userIdList)
	}
	// 查询分页信息
	var messageList []*model.Message
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&messageList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var messageAdminDtoList []*response.MessageAdminDto
	for _, item := range messageList {
		var messageAdminDto *response.MessageAdminDto
		messageAdminDto, err = item.CopyToMessageAdminDto()
		if err != nil {
			return nil, err
		}
		messageAdminDtoList = append(messageAdminDtoList, messageAdminDto)
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(messageAdminDtoList)),
		Total:   count,
		Records: messageAdminDtoList,
	}
	return &page, nil
}

func (a *MessageService) MessageDtoPage(pageVo request.PageVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载用户
	tx := db.Preload("User")
	// 设置查询条件和排序条件
	tx = tx.Where("deleted = ?", model.MODEL_ACTIVED).Order("create_time desc")
	// 查询分页信息
	var messageList []*model.Message
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&messageList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var messageDtoList []*response.MessageDto
	for _, item := range messageList {
		var messageDto *response.MessageDto
		messageDto, err = item.CopyToMessageDto()
		if err != nil {
			return nil, err
		}
		messageDtoList = append(messageDtoList, messageDto)
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(messageDtoList)),
		Total:   count,
		Records: messageDtoList,
	}
	return &page, nil
}

func (a *MessageService) CountMessage() (int, error) {
	var count int
	err := db.Model(&model.Message{}).Count(&count).Error
	return count, err
}

func (a *MessageService) SaveMessage(messageSaveVo request.MessageSaveVo) (bool, error) {
	// 构造留言
	var message model.Message
	err := utils.CopyFields(&message, messageSaveVo)
	if err != nil {
		return false, err
	}
	// 新增留言
	err = db.Create(&message).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *MessageService) RemoveMessageById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Message{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *MessageService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var message model.Message
	message.ID = modelDeletedVo.ID
	err := db.Model(&message).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
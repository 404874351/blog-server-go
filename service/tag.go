package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

type TagService struct {}

func (a *TagService) TagOptionDtoList() ([]*response.TagOptionDto, error) {
	var tagOptionDtoList []*response.TagOptionDto
	err := db.Table("tag").Select("id, name").Scan(&tagOptionDtoList).Error
	if err != nil {
		return nil, err
	}
	return tagOptionDtoList, nil
}

func (a *TagService) QueryTagByName(name string) ([]*response.TagOptionDto, error) {
	// 动态拼接查询条件
	tx := db
	if name != "" {
		tx = tx.Where("name LIKE ?", "%" + name + "%")
	}
	// 只查询10条记录
	var tagOptionDtoList []*response.TagOptionDto
	err := tx.Table("tag").
		Select("id, name").
		Limit(10).Offset(0).
		Scan(&tagOptionDtoList).Error
	if err != nil {
		return nil, err
	}
	return tagOptionDtoList, nil
}

func (a *TagService) TagDtoPage(pageVo request.PageVo, tagVo request.TagVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 动态拼接查询条件
	tx := db
	if tagVo.Name != "" {
		tx = tx.Where("name LIKE ?", "%" + tagVo.Name + "%")
	}
	// 查询分页信息
	var tagDtoList []*response.TagDto
	var count int64
	err := tx.Table("tag").
		Select("id, name, create_time, deleted, " +
			"( select count(1) from relation_article_tag rat where rat.tag_id=tag.id ) as article_count").
		Limit(limit).Offset(offset).
		Scan(&tagDtoList).
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(tagDtoList)),
		Total:   count,
		Records: tagDtoList,
	}
	return &page, nil
}

func (a *TagService) SaveTag(tagSaveVo request.TagSaveVo) (*response.TagOptionDto, error) {
	// 构造标签
	var tag model.Tag
	err := utils.CopyFields(&tag, tagSaveVo)
	if err != nil {
		return nil, err
	}
	// 新增标签
	err = db.Create(&tag).Error
	if err != nil {
		return nil, err
	}
	var tagOptionDto response.TagOptionDto
	err = utils.CopyFields(&tagOptionDto, tag)
	if err != nil {
		return nil, err
	}
	return &tagOptionDto, nil
}

func (a *TagService) RemoveTagById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *TagService) UpdateTag(tagUpdateVo request.TagUpdateVo) (bool, error) {
	var tag model.Tag
	err := utils.CopyFields(&tag, tagUpdateVo)
	if err != nil {
		return false, err
	}

	err = db.Model(&tag).Updates(tag).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *TagService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var tag model.Tag
	tag.ID = modelDeletedVo.ID
	err := db.Model(&tag).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

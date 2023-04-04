package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

type CategoryService struct {}

func (a *CategoryService) CategoryOptionDtoList() ([]*response.CategoryOptionDto, error) {
	var categoryOptionDtoList []*response.CategoryOptionDto
	err := db.Table("category").Select("id, name").Scan(&categoryOptionDtoList).Error
	if err != nil {
		return nil, err
	}
	return categoryOptionDtoList, nil
}

func (a *CategoryService) QueryCategoryByName(name string) ([]*response.CategoryOptionDto, error) {
	// 动态拼接查询条件
	tx := db
	if name != "" {
		tx = tx.Where("name LIKE ?", "%" + name + "%")
	}
	// 只查询10条记录
	var categoryOptionDtoList []*response.CategoryOptionDto
	err := tx.Table("category").
		Select("id, name").
		Limit(10).Offset(0).
		Scan(&categoryOptionDtoList).Error
	if err != nil {
		return nil, err
	}
	return categoryOptionDtoList, nil
}

func (a *CategoryService) CategoryDtoPage(pageVo request.PageVo, categoryVo request.CategoryVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 动态拼接查询条件
	tx := db
	if categoryVo.Name != "" {
		tx = tx.Where("name LIKE ?", "%" + categoryVo.Name + "%")
	}
	// 查询分页信息
	var categoryDtoList []*response.CategoryDto
	var count int64
	err := tx.Table("category").
		Select("id, name, create_time, deleted, " +
			"( select count(1) from article where article.category_id=category.id ) as article_count").
		Limit(limit).Offset(offset).
		Scan(&categoryDtoList).
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(categoryDtoList)),
		Total:   count,
		Records: categoryDtoList,
	}
	return &page, nil
}

func (a *CategoryService) SaveCategory(categorySaveVo request.CategorySaveVo) (*response.CategoryOptionDto, error) {
	// 构造标签
	var category model.Category
	err := utils.CopyFields(&category, categorySaveVo)
	if err != nil {
		return nil, err
	}
	// 新增标签
	err = db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	var categoryOptionDto response.CategoryOptionDto
	err = utils.CopyFields(&categoryOptionDto, category)
	if err != nil {
		return nil, err
	}
	return &categoryOptionDto, nil
}

func (a *CategoryService) RemoveCategoryById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Category{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CategoryService) UpdateCategory(categoryUpdateVo request.CategoryUpdateVo) (bool, error) {
	var category model.Category
	err := utils.CopyFields(&category, categoryUpdateVo)
	if err != nil {
		return false, err
	}

	err = db.Model(&category).Updates(category).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *CategoryService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var category model.Category
	category.ID = modelDeletedVo.ID
	err := db.Model(&category).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

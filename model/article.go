package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// Article
//  @Description: 文章标签
//
type Article struct {
	Model
	// 标题
	Title			string		`json:"title"         gorm:"not null;size:255"`
	// 简介
	Description		string		`json:"description"   gorm:"not null;size:1023"`
	// 封面图链接
	CoverUrl		string		`json:"coverUrl"      gorm:"not null;size:255"`
	// 原文链接
	ContentUrl		string		`json:"contentUrl"    gorm:"not null;size:255"`
	// 浏览量
	ViewCount		int		    `json:"viewCount"     gorm:"not null"`
	// 用户id
	UserId		    int		    `json:"userId"        gorm:"not null"`
	// 分类id
	CategoryId		int		    `json:"categoryId"    gorm:"not null"`
	// 置顶，0否，1是，默认0
	Top			    *int8		`json:"top"           gorm:"not null"`
	// 关闭评论，0否，1是，默认0
	CloseComment	*int8		`json:"closeComment"  gorm:"not null"`

	// 用户
	User			*User       `json:"user"          gorm:"foreignkey:UserId"`
	// 分类
	Category		*Category   `json:"category"      gorm:"foreignkey:CategoryId"`
	// 标签列表，多对多关联
	TagList         []*Tag      `json:"tagList"       gorm:"many2many:relation_article_tag"`
}

const (
	ARTICLE_TOP_DISABLE    				int8 = 0
	ARTICLE_TOP_ENABLE     				int8 = 1
	ARTICLE_CLOSE_COMMENT_DISABLE    	int8 = 0
	ARTICLE_CLOSE_COMMENT_ENABLE     	int8 = 1
)

func (a *Article) CopyToArticleAdminDto() (*response.ArticleAdminDto, error) {
	var articleAdminDto response.ArticleAdminDto
	err := utils.CopyFields(&articleAdminDto, a)
	if err != nil {
		return nil, err
	}
	if a.User != nil {
		articleAdminDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(articleAdminDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	if a.Category != nil {
		articleAdminDto.Category = &response.CategoryOptionDto{}
		err = utils.CopyFields(articleAdminDto.Category, a.Category)
		if err != nil {
			return nil, err
		}
	}
	for _, item := range a.TagList {
		tagOption := &response.TagOptionDto{}
		err = utils.CopyFields(tagOption, item)
		if err != nil {
			return nil, err
		}
		articleAdminDto.TagList = append(articleAdminDto.TagList, tagOption)
	}
	return &articleAdminDto, nil
}

func (a *Article) CopyToArticleDto() (*response.ArticleDto, error) {
	var articleDto response.ArticleDto
	err := utils.CopyFields(&articleDto, a)
	if err != nil {
		return nil, err
	}
	if a.User != nil {
		articleDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(articleDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	if a.Category != nil {
		articleDto.Category = &response.CategoryOptionDto{}
		err = utils.CopyFields(articleDto.Category, a.Category)
		if err != nil {
			return nil, err
		}
	}
	for _, item := range a.TagList {
		tagOption := &response.TagOptionDto{}
		err = utils.CopyFields(tagOption, item)
		if err != nil {
			return nil, err
		}
		articleDto.TagList = append(articleDto.TagList, tagOption)
	}
	return &articleDto, nil
}

func (a *Article) CopyToArticleUpdateDto() (*response.ArticleUpdateDto, error) {
	var articleUpdateDto response.ArticleUpdateDto
	err := utils.CopyFields(&articleUpdateDto, a)
	if err != nil {
		return nil, err
	}
	if a.Category != nil {
		articleUpdateDto.Category = &response.CategoryOptionDto{}
		err = utils.CopyFields(articleUpdateDto.Category, a.Category)
		if err != nil {
			return nil, err
		}
	}
	for _, item := range a.TagList {
		tagOption := &response.TagOptionDto{}
		err = utils.CopyFields(tagOption, item)
		if err != nil {
			return nil, err
		}
		articleUpdateDto.TagList = append(articleUpdateDto.TagList, tagOption)
	}
	return &articleUpdateDto, nil
}
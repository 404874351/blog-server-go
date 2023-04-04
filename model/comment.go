package model

import (
	"blog-server-go/model/response"
	"blog-server-go/utils"
)

//
// Comment
//  @Description: 文章评论
//
type Comment struct {
	Model
	// 评论内容
	Content			string		`json:"content"         gorm:"not null;size:1023"`
	// 置顶，0否，1是，默认0，一篇文章只能有一个顶级评论被置顶
	Top			    *int8		`json:"top"             gorm:"not null"`
	// 用户id
	UserId		    int		    `json:"userId"          gorm:"not null"`
	// 文章id
	ArticleId		int		    `json:"articleId"       gorm:"not null"`
	// 父评论id，顶级评论为空，用于检索二级评论
	ParentId		*int		`json:"parentId"        gorm:""`
	// 回复用户id，顶级评论为空，用于检索三级评论，但统一显示为二级
	ReplyUserId		*int		`json:"replyUserId"     gorm:""`

	// 留言用户
	User			*User       `json:"user"            gorm:"foreignkey:UserId"`
	// 回复用户
	ReplyUser		*User       `json:"replyUser"       gorm:"foreignkey:ReplyUserId"`
	// 关联文章
	Article			*Article    `json:"article"         gorm:"foreignkey:ArticleId"`
}

const (
	COMMENT_TOP_DISABLE    int8 = 0
	COMMENT_TOP_ENABLE     int8 = 1
)

func (a *Comment) CopyToCommentAdminDto() (*response.CommentAdminDto, error) {
	var commentAdminDto response.CommentAdminDto
	err := utils.CopyFields(&commentAdminDto, a)
	if err != nil {
		return nil, err
	}
	commentAdminDto.ArticleTitle = a.Article.Title
	if a.User != nil {
		commentAdminDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(commentAdminDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	if a.ReplyUser != nil {
		commentAdminDto.ReplyUser = &response.UserBaseInfoDto{}
		err = utils.CopyFields(commentAdminDto.ReplyUser, a.ReplyUser)
		if err != nil {
			return nil, err
		}
	}
	return &commentAdminDto, nil
}

func (a *Comment) CopyToCommentDto() (*response.CommentDto, error) {
	var commentDto response.CommentDto
	err := utils.CopyFields(&commentDto, a)
	if err != nil {
		return nil, err
	}

	if a.User != nil {
		commentDto.User = &response.UserBaseInfoDto{}
		err = utils.CopyFields(commentDto.User, a.User)
		if err != nil {
			return nil, err
		}
	}
	if a.ReplyUser != nil {
		commentDto.ReplyUser = &response.UserBaseInfoDto{}
		err = utils.CopyFields(commentDto.ReplyUser, a.ReplyUser)
		if err != nil {
			return nil, err
		}
	}
	return &commentDto, nil
}
package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
	"errors"
	"strconv"
)

type CommentService struct {}

func (a *CommentService) CommentAdminDtoPage(pageVo request.PageVo, commentAdminVo request.CommentAdminVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载
	tx := db
	tx = tx.Preload("User").Preload("ReplyUser").Preload("Article")
	// 动态拼接查询条件
	if commentAdminVo.Content != "" {
		tx = tx.Where("content LIKE ?", "%" + commentAdminVo.Content + "%")
	}
	if commentAdminVo.Top != nil {
		tx = tx.Where("top = ?", *commentAdminVo.Top)
	}
	if commentAdminVo.Nickname != "" {
		var userIdList []int
		err := db.Table("user").
			Select("id").
			Where("nickname LIKE ?", "%" + commentAdminVo.Nickname + "%").
			Pluck("id", &userIdList).Error
		if err != nil {
			return nil, err
		}
		tx = tx.Where("user_id in (?)", userIdList)
	}
	if commentAdminVo.ArticleTitle != "" {
		var articleIdList []int
		err := db.Table("article").
			Select("id").
			Where("title LIKE ?", "%" + commentAdminVo.ArticleTitle + "%").
			Pluck("id", &articleIdList).Error
		if err != nil {
			return nil, err
		}
		tx = tx.Where("article_id in (?)", articleIdList)
	}

	// 查询分页信息
	var commentList []*model.Comment
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&commentList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var commentAdminDtoList []*response.CommentAdminDto
	for _, item := range commentList {
		var commentAdminDto *response.CommentAdminDto
		commentAdminDto, err = item.CopyToCommentAdminDto()
		if err != nil {
			return nil, err
		}
		commentAdminDtoList = append(commentAdminDtoList, commentAdminDto)
	}

	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(commentAdminDtoList)),
		Total:   count,
		Records: commentAdminDtoList,
	}
	return &page, nil
}

func (a *CommentService) CommentDtoPage(pageVo request.PageVo, commentVo request.CommentVo) (*response.Page, error) {
	if commentVo.ParentId != 0 {
		return a.ChildrenCommentDtoPage(pageVo, commentVo)
	}
	// 查询父评论
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	tx := db
	tx = tx.Preload("User").Preload("ReplyUser").Preload("Article")
	// 动态拼接查询条件
	tx = tx.Where("article_id = ?", commentVo.ArticleId)
	tx = tx.Where("parent_id is null")
	tx = tx.Where("deleted = ?", model.MODEL_ACTIVED)
	tx = tx.Order("top desc")
	if commentVo.SortBy != "" {
		tx = tx.Order( commentVo.SortBy + " desc")
	}
	var commentList []*model.Comment
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&commentList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var commentDtoList []*response.CommentDto
	var parentCommentIdList []int
	for _, item := range commentList {
		var commentDto *response.CommentDto
		commentDto, err = item.CopyToCommentDto()
		if err != nil {
			return nil, err
		}
		commentDtoList = append(commentDtoList, commentDto)
		parentCommentIdList = append(parentCommentIdList, commentDto.ID)
	}
	// 查询父评论的附加信息
	var parentCommentExtraList []*response.CommentExtra
	var sql = "id, " +
		"( select count(1) from relation_user_comment ruc where ruc.comment_id=comment.id ) as praise_count, " +
		"( select count(1) from comment c1 where c1.parent_id=comment.id and c1.deleted = 0 ) as children_count"
	if commentVo.UserId != 0 {
		sql += ", ( select count(1) > 0 from relation_user_comment ruc where ruc.comment_id=comment.id and ruc.user_id=" +
			strconv.FormatInt(int64(commentVo.UserId), 10) +
				" ) as praise_status"
	}
	err = db.Table("comment").
		Select(sql).
		Where("id in (?)", parentCommentIdList).
		Scan(&parentCommentExtraList).Error
	if err != nil {
		return nil, err
	}
	// 查询每个父评论下的部分子评论
	childrenMap := map[int]*response.CommentExtra{}
	var childrenPageVo = request.PageVo{Current: 1, Size: 3}
	for _, item := range parentCommentExtraList {
		commentVo.ParentId = item.ID
		var childrenPage *response.Page
		childrenPage, err = a.ChildrenCommentDtoPage(childrenPageVo, commentVo)
		if err != nil {
			return nil, err
		}
		var ok bool
		item.Children, ok = childrenPage.Records.([]*response.CommentDto)
		if !ok {
			return nil, err
		}
		childrenMap[item.ID] = item
	}
	// 封装子评论
	for _, item := range commentDtoList {
		if childrenMap[item.ID] == nil {
			continue
		}
		err = utils.CopyFields(item, childrenMap[item.ID])
		if err != nil {
			return nil, err
		}
	}

	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(commentDtoList)),
		Total:   count,
		Records: commentDtoList,
	}
	return &page, nil
}


func (a *CommentService) ChildrenCommentDtoPage(pageVo request.PageVo, commentVo request.CommentVo) (*response.Page, error) {
	if commentVo.ParentId == 0 {
		return a.CommentDtoPage(pageVo, commentVo)
	}
	// 查询评论
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	tx := db
	tx = tx.Preload("User").Preload("ReplyUser").Preload("Article")
	// 拼接查询条件
	tx = tx.Where("article_id = ?", commentVo.ArticleId)
	tx = tx.Where("parent_id = ?", commentVo.ParentId)
	tx = tx.Where("deleted = ?", model.MODEL_ACTIVED)
	tx = tx.Order("create_time")

	var commentList []*model.Comment
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&commentList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var commentDtoList []*response.CommentDto
	var commentIdList []int
	for _, item := range commentList {
		var commentDto *response.CommentDto
		commentDto, err = item.CopyToCommentDto()
		if err != nil {
			return nil, err
		}
		commentDtoList = append(commentDtoList, commentDto)
		commentIdList = append(commentIdList, commentDto.ID)
	}
	// 查询评论的附加信息
	var commentExtraList []*response.CommentExtra
	var sql = "id, " +
		"( select count(1) from relation_user_comment ruc where ruc.comment_id=comment.id ) as praise_count, " +
		"( select count(1) from comment c1 where c1.parent_id=comment.id and c1.deleted = 0 ) as children_count"
	if commentVo.UserId != 0 {
		sql += ", ( select count(1) > 0 from relation_user_comment ruc where ruc.comment_id=comment.id and ruc.user_id=" +
			strconv.FormatInt(int64(commentVo.UserId), 10) +
			" ) as praise_status"
	}
	err = db.Table("comment").
		Select(sql).
		Where("id in (?)", commentIdList).
		Scan(&commentExtraList).Error
	if err != nil {
		return nil, err
	}
	// 封装附加信息
	extraMap := map[int]*response.CommentExtra{}
	for _, item := range commentExtraList {
		extraMap[item.ID] = item
	}
	for _, item := range commentDtoList {
		if extraMap[item.ID] == nil {
			continue
		}
		err = utils.CopyFields(item, extraMap[item.ID])
		if err != nil {
			return nil, err
		}
	}

	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(commentDtoList)),
		Total:   count,
		Records: commentDtoList,
	}
	return &page, nil
}

func (a *CommentService) CountComment() (int64, error) {
	var count int64
	err := db.Table("comment").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *CommentService) SaveComment(commentSaveVo request.CommentSaveVo) (*response.CommentDto, error) {
	// 构造评论
	var comment model.Comment
	err := utils.CopyFields(&comment, commentSaveVo)
	if err != nil {
		return nil, err
	}
	top := model.COMMENT_TOP_DISABLE
	comment.Top = &top
	// 新增评论
	err = db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	// 查询完整信息
	tx := db.Preload("User").Preload("ReplyUser").Preload("Article")
	err = tx.First(&comment, comment.ID).Error
	if err != nil {
		return nil, err
	}
	var commentDto *response.CommentDto
	commentDto, err = comment.CopyToCommentDto()
	if err != nil {
		return nil, err
	}
	// 查询附加信息
	var commentExtra response.CommentExtra
	var sql = "id, " +
		"( select count(1) from relation_user_comment ruc where ruc.comment_id=comment.id ) as praise_count, " +
		"( select count(1) from comment c1 where c1.parent_id=comment.id and c1.deleted = 0 ) as children_count"
	err = db.Table("comment").
		Select(sql).
		Where("id = ?", comment.ID).
		Scan(&commentExtra).Error
	if err != nil {
		return nil, err
	}
	err = utils.CopyFields(commentDto, commentExtra)
	if err != nil {
		return nil, err
	}

	return commentDto, nil
}

func (a *CommentService) PraiseComment(commentId int, userId int)  (bool, error) {
	err := db.Exec("insert into relation_user_comment(`user_id`, `comment_id`) values(?, ?)", userId, commentId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CommentService) CancelPraiseComment(commentId int, userId int)  (bool, error) {
	err := db.Exec("delete from relation_user_comment where user_id = ? and comment_id = ?", userId, commentId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CommentService) UpdateComment(commentUpdateVo request.CommentUpdateVo) (bool, error) {
	// 仅仅更新top字段，其他字段不生效
	var comment model.Comment
	err := utils.CopyFields(&comment, commentUpdateVo)
	if err != nil {
		return false, err
	}
	// 如需更新置顶字段，则必须保证该文章下只有唯一顶级评论可置顶
	if comment.Top != nil && *comment.Top == model.COMMENT_TOP_ENABLE {
		var commentCheck model.Comment
		var count int64
		// 置顶评论必唯一
		err = db.Table("comment").Where("article_id = ? and top = ?", comment.ArticleId, model.COMMENT_TOP_ENABLE).Count(&count).Error
		if err != nil {
			return false, err
		}
		if count != 0 {
			return false, errors.New("该文章的置顶评论已存在且唯一")
		}
		// 置顶评论必顶级
		err = db.Where("id = ?", comment.ID).First(&commentCheck).Error
		if err != nil {
			return false, err
		}
		if commentCheck.ParentId != nil {
			return false, errors.New("该评论不是顶级评论")
		}
	}

	err = db.Model(&comment).Updates(comment).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *CommentService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var comment model.Comment
	comment.ID = modelDeletedVo.ID
	err := db.Model(&comment).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CommentService) RemoveCommentById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CommentService) RemoveComment(commentId int, userId int)  (bool, error) {
	err := db.Where("id = ? and user_id = ?", commentId, userId).Delete(&model.Comment{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}


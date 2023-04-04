package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ArticleService struct {}

func (a *ArticleService) ArticleAdminDtoPage(pageVo request.PageVo, articleAdminVo request.ArticleAdminVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载
	tx := db.Preload("User").Preload("Category").Preload("TagList")
	// 动态拼接查询条件
	if articleAdminVo.Title != "" {
		tx = tx.Where("title LIKE ?", "%" + articleAdminVo.Title + "%")
	}
	if articleAdminVo.CategoryId != 0 {
		tx = tx.Where("category_id = ?", articleAdminVo.CategoryId)
	}
	if articleAdminVo.Nickname != "" {
		tx = tx.Where("user_id in ( select id from user where nickname Like ? )", "%" + articleAdminVo.Nickname + "%")
	}
	if articleAdminVo.TagId != 0 {
		tx = tx.Where("id in ( select article_id from relation_article_tag where tag_id = ? )", articleAdminVo.TagId)
	}
	// 查询分页信息
	var articleList []*model.Article
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&articleList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var articleAdminDtoList []*response.ArticleAdminDto
	var articleIdList []int
	for _, item := range articleList {
		var articleAdmin *response.ArticleAdminDto
		articleAdmin, err = item.CopyToArticleAdminDto()
		articleAdminDtoList = append(articleAdminDtoList, articleAdmin)
		articleIdList = append(articleIdList, item.ID)
	}
	// 查询附加信息
	var articleExtraList []*response.ArticleExtra
	var sql = "id, " +
		"( select count(1) from relation_user_article rua where article.id=rua.article_id ) as praise_count, " +
		"( select count(1) from comment where article.id=comment.article_id and comment.deleted=0 ) as comment_count"
	err = db.Table("article").
		Select(sql).
		Where("id in (?)", articleIdList).
		Scan(&articleExtraList).Error
	if err != nil {
		return nil, err
	}
	// 封装附加信息
	articleMap := map[int]*response.ArticleExtra{}
	for _, item := range articleExtraList {
		articleMap[item.ID] = item
	}
	for _, item := range articleAdminDtoList {
		if articleMap[item.ID] == nil {
			continue
		}
		err = utils.CopyFields(item, articleMap[item.ID])
		if err != nil {
			return nil, err
		}
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(articleAdminDtoList)),
		Total:   count,
		Records: articleAdminDtoList,
	}
	return &page, nil
}

func (a *ArticleService) ArticleDtoPage(pageVo request.PageVo, articleVo request.ArticleVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载
	tx := db.Preload("User").Preload("Category").Preload("TagList")
	// 动态拼接查询条件
	tx = tx.Order("top desc")
	if articleVo.SortBy != "" {
		tx = tx.Order(articleVo.SortBy + " desc")
	}
	if articleVo.CategoryId != 0 {
		tx = tx.Where("category_id = ?", articleVo.CategoryId)
	}
	if articleVo.Key != "" {
		tx = tx.Where("( " +
				"title LIKE ? " +
				"OR description LIKE ? " +
				"OR category_id in ( select id from category where name LIKE ? ) " +
				"OR id in ( select article_id from relation_article_tag where tag_id in ( select id from tag where name LIKE ? ) )" +
			" )", "%" + articleVo.Key + "%", "%" + articleVo.Key + "%", "%" + articleVo.Key + "%", "%" + articleVo.Key + "%")
	}
	// 查询分页信息
	var articleList []*model.Article
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&articleList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var articleDtoList []*response.ArticleDto
	var articleIdList []int
	for _, item := range articleList {
		var articleDto *response.ArticleDto
		articleDto, err = item.CopyToArticleDto()
		articleDtoList = append(articleDtoList, articleDto)
		articleIdList = append(articleIdList, item.ID)
	}
	// 查询附加信息
	var articleExtraList []*response.ArticleExtra
	var sql = "id, " +
		"( select count(1) from relation_user_article rua where article.id=rua.article_id ) as praise_count, " +
		"( select count(1) from comment where article.id=comment.article_id and comment.deleted=0 ) as comment_count"
	err = db.Table("article").
		Select(sql).
		Where("id in (?)", articleIdList).
		Scan(&articleExtraList).Error
	if err != nil {
		return nil, err
	}
	// 封装附加信息
	articleMap := map[int]*response.ArticleExtra{}
	for _, item := range articleExtraList {
		articleMap[item.ID] = item
	}
	for _, item := range articleDtoList {
		if articleMap[item.ID] == nil {
			continue
		}
		err = utils.CopyFields(item, articleMap[item.ID])
		if err != nil {
			return nil, err
		}
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(articleDtoList)),
		Total:   count,
		Records: articleDtoList,
	}
	return &page, nil
}

func (a *ArticleService) ArticleDashboardDtoPage(pageVo request.PageVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载
	tx := db
	// 查询分页信息
	var articleList []*model.Article
	var count int64
	err := tx.Select("id, title, view_count").Order("view_count desc").Limit(limit).Offset(offset).Find(&articleList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var articleDashboardDtoList []*response.ArticleDashboardDto
	for _, item := range articleList {
		articleDashboardDto := &response.ArticleDashboardDto{}
		err = utils.CopyFields(articleDashboardDto, item)
		if err != nil {
			return nil, err
		}
		articleDashboardDtoList = append(articleDashboardDtoList, articleDashboardDto)
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(articleDashboardDtoList)),
		Total:   count,
		Records: articleDashboardDtoList,
	}
	return &page, nil
}

func (a *ArticleService) GetArticleUpdateDtoById(id int) (*response.ArticleUpdateDto, error) {
	// 构造查询条件
	tx := db.Preload("Category").Preload("TagList")
	tx = tx.Where("id = ?", id)
	// 查询信息
	var article model.Article
	err := tx.Find(&article).Error
	if err != nil {
		return nil, err
	}

	var articleUpdate *response.ArticleUpdateDto
	articleUpdate, err = article.CopyToArticleUpdateDto()
	if err != nil {
		return nil, err
	}
	return articleUpdate, nil
}

func (a *ArticleService) GetArticleDtoById(id int, userId int) (*response.ArticleDto, error) {
	// 构造查询条件
	tx := db.Preload("User").Preload("Category").Preload("TagList")
	tx = tx.Where("id = ?", id)
	// 查询信息
	var article model.Article
	err := tx.Find(&article).Error
	if err != nil {
		return nil, err
	}
	var articleDto *response.ArticleDto
	articleDto, err = article.CopyToArticleDto()
	if err != nil {
		return nil, err
	}
	// 查询附加信息
	var articleExtra response.ArticleExtra
	var sql = "id, " +
		"( select count(1) from relation_user_article rua where article.id=rua.article_id ) as praise_count, " +
		"( select count(1) from comment where article.id=comment.article_id and comment.deleted=0 ) as comment_count, " +
		"( select count(1) > 0 from relation_user_article rua where article.id=rua.article_id and rua.user_id = " + strconv.Itoa(userId) + " ) as praise_status"
	err = db.Table("article").
		Select(sql).
		Where("id = ?", id).
		Scan(&articleExtra).Error
	if err != nil {
		return nil, err
	}
	err = utils.CopyFields(articleDto, articleExtra)
	if err != nil {
		return nil, err
	}

	return articleDto, nil
}

func (a *ArticleService) SumArticleViewCount() (int64, error) {
	var sum []int64
	err := db.Table("article").Select("sum(article.view_count) as sum").Pluck("sum", &sum).Error
	if err != nil {
		return 0, err
	}
	return sum[0], nil
}

func (a *ArticleService) CountArticle() (int64, error) {
	var count int64
	err := db.Table("article").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleService) CountArticlePraise() (int64, error) {
	var count int64
	err := db.Table("relation_user_article").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleService) SaveArticle(articleSaveVo request.ArticleSaveVo) (bool, error) {
	// 构造文章
	var article model.Article
	err := utils.CopyFields(&article, articleSaveVo)
	if err != nil {
		return false, err
	}
	top := model.ARTICLE_TOP_DISABLE
	article.Top = &top
	closeComment := model.ARTICLE_CLOSE_COMMENT_DISABLE
	article.CloseComment = &closeComment
	// 新增文章
	tx := db.Begin()
	err = tx.Create(&article).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 插入标签列表
	var tagList []model.Tag
	for _, tagId := range articleSaveVo.TagIdList {
		tag := model.Tag{Model: model.Model{ ID: tagId}}
		tagList = append(tagList, tag)
	}
	err = tx.Model(&article).Association("TagList").Append(tagList).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (a *ArticleService) ViewArticle(id int) (bool, error) {
	var article model.Article
	article.ID = id
	err := db.Model(&article).Update("view_count", gorm.Expr("view_count + 1")).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleService) PraiseArticle(articleId int, userId int) (bool, error) {
	err := db.Exec("insert into relation_user_article(`user_id`, `article_id`) values(?, ?)", userId, articleId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleService) CancelPraiseArticle(articleId int, userId int) (bool, error) {
	err := db.Exec("delete from relation_user_article where user_id = ? and article_id = ?", userId, articleId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleService) UpdateArticle(articleUpdateVo request.ArticleUpdateVo) (bool, error) {
	var article model.Article
	err := utils.CopyFields(&article, articleUpdateVo)
	if err != nil {
		return false, err
	}
	// 更新基本信息
	tx := db.Begin()
	err = tx.Model(&article).Updates(article).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 更新标签列表
	if len(articleUpdateVo.TagIdList) != 0 {
		// 清空原有列表
		err = tx.Model(&article).Association("TagList").Clear().Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
		// 插入标签列表
		var tagList []model.Tag
		for _, tagId := range articleUpdateVo.TagIdList {
			menu := model.Tag{Model: model.Model{ ID: tagId}}
			tagList = append(tagList, menu)
		}
		err = tx.Model(&article).Association("TagList").Append(tagList).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (a *ArticleService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var article model.Article
	article.ID = modelDeletedVo.ID
	err := db.Model(&article).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleService) RemoveArticleById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.Article{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
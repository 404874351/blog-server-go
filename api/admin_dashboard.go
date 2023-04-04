package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
)

func AdminDashboardIndex(c *gin.Context) {
	resMap := map[string]interface{}{}
	viewCount, err := articleService.SumArticleViewCount()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["view"] = viewCount

	articleCount, err := articleService.CountArticle()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["article"] = articleCount

	userCount, err := userService.CountUser()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["user"] = userCount

	messageCount, err := messageService.CountMessage()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["message"] = messageCount

	middleware.SetData(c, resMap)
}

func AdminDashboardView(c *gin.Context) {
	var pageVo request.PageVo
	pageVo.New()
	page, err := articleService.ArticleDashboardDtoPage(pageVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page.Records)
}

func AdminDashboardRole(c *gin.Context) {
	list, err := roleService.RoleDashboardDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminDashboardCategory(c *gin.Context) {
	list, err := categoryService.CategoryOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminDashboardTag(c *gin.Context) {
	list, err := tagService.TagOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}
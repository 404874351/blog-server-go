package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)


func AdminArticlePage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var articleAdminVo request.ArticleAdminVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&articleAdminVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询分页
	var page *response.Page
	page, err = articleService.ArticleAdminDtoPage(pageVo, articleAdminVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func AdminArticleDetail(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询记录
	var articleUpdateDto *response.ArticleUpdateDto
	articleUpdateDto, err = articleService.GetArticleUpdateDtoById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, articleUpdateDto)
}

func AdminArticleSave(c *gin.Context) {
	// 绑定并检查参数
	var articleSaveVo request.ArticleSaveVo
	var err error
	err = c.ShouldBind(&articleSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 新增数据
	var res bool
	res, err = articleService.SaveArticle(articleSaveVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

func AdminArticleRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = articleService.RemoveArticleById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminArticleUpdate(c *gin.Context) {
	// 绑定并检查参数
	var articleUpdateVo request.ArticleUpdateVo
	var err error
	err = c.ShouldBind(&articleUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	articleUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新数据
	var res bool
	res, err = articleService.UpdateArticle(articleUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminArticleUpdateDeleted(c *gin.Context) {
	// 绑定并检查参数
	var deletedVo request.ModelDeletedVo
	var err error
	err = c.ShouldBind(&deletedVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	deletedVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = articleService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
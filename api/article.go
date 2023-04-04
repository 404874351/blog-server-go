package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ArticleStatistic(c *gin.Context) {
	resMap := map[string]interface{}{}
	articleCount, err := articleService.CountArticle()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["count"] = articleCount

	viewCount, err := articleService.SumArticleViewCount()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["viewCount"] = viewCount

	praiseCount, err := articleService.CountArticlePraise()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["praiseCount"] = praiseCount

	commentCount, err := commentService.CountComment()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	resMap["commentCount"] = commentCount

	middleware.SetData(c, resMap)
}

func ArticlePage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var articleVo request.ArticleVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&articleVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询分页
	var page *response.Page
	page, err = articleService.ArticleDtoPage(pageVo, articleVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func ArticleDetail(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 从登录状态中获取用户id
	var userId int
	claims := middleware.GetClaims(c)
	if claims != nil {
		userId = claims.ID
	}
	// 更新浏览记录
	var res bool
	res, err = articleService.ViewArticle(id)
	if err != nil || !res {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	// 查询记录
	var articleUpdateDto *response.ArticleDto
	articleUpdateDto, err = articleService.GetArticleDtoById(id, userId)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, articleUpdateDto)
}

func ArticlePraise(c *gin.Context) {
	// 获取文章id
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	var res bool
	res, err = articleService.PraiseArticle(articleId, claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
func ArticleCancelPraise(c *gin.Context) {
	// 获取文章id
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	var res bool
	res, err = articleService.CancelPraiseArticle(articleId, claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
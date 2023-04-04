package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CommentPage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var commentVo request.CommentVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&commentVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims != nil {
		commentVo.UserId = claims.ID
	}
	// 查询分页
	var page *response.Page
	page, err = commentService.CommentDtoPage(pageVo, commentVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func CommentSave(c *gin.Context) {
	// 绑定并检查参数
	var commentSaveVo request.CommentSaveVo
	var err error
	err = c.ShouldBind(&commentSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	commentSaveVo.UserId = claims.ID
	// 新增数据
	var commentDto *response.CommentDto
	commentDto, err = commentService.SaveComment(commentSaveVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, commentDto)
}

func CommentPraise(c *gin.Context) {
	// 获取评论id
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	var res bool
	res, err = commentService.PraiseComment(commentId, claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)

}

func CommentCancelPraise(c *gin.Context) {
	// 获取评论id
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	var res bool
	res, err = commentService.CancelPraiseComment(commentId, claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func CommentRemove(c *gin.Context) {
	// 获取评论id
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 获取登录态
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	// 更新记录
	var res bool
	res, err = commentService.RemoveComment(commentId, claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

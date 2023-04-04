package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminCommentPage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var commentAdminVo request.CommentAdminVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&commentAdminVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询分页
	var page *response.Page
	page, err = commentService.CommentAdminDtoPage(pageVo, commentAdminVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func AdminCommentRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = commentService.RemoveCommentById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminCommentUpdate(c *gin.Context) {
	// 绑定并检查参数
	var commentUpdateVo request.CommentUpdateVo
	var err error
	err = c.ShouldBind(&commentUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	commentUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}

	// 更新记录
	var res bool
	res, err = commentService.UpdateComment(commentUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminCommentUpdateDeleted(c *gin.Context) {
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
	res, err = commentService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
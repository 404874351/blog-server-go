package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
)

func MessageCount(c *gin.Context) {
	count, err := messageService.CountMessage()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, count)
}

func MessagePage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	// 查询分页
	var page *response.Page
	page, err = messageService.MessageDtoPage(pageVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func MessageSave(c *gin.Context) {
	// 绑定并检查参数
	var messageSaveVo request.MessageSaveVo
	var err error
	err = c.ShouldBind(&messageSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 从登录状态中获取用户id
	claims := middleware.GetClaims(c)
	if claims != nil {
		messageSaveVo.UserId = &claims.ID
	} else {
		messageSaveVo.UserId = nil
	}
	// 新增数据
	var res bool
	res, err = messageService.SaveMessage(messageSaveVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}
package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminTagPage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var tagVo request.TagVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&tagVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询分页
	var page *response.Page
	page, err = tagService.TagDtoPage(pageVo, tagVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func AdminTagQuery(c *gin.Context) {
	name := c.Query("name")
	list, err := tagService.QueryTagByName(name)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminTagOption(c *gin.Context) {
	list, err := tagService.TagOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminTagSave(c *gin.Context) {
	// 绑定并检查参数
	var tagSaveVo request.TagSaveVo
	var err error
	err = c.ShouldBind(&tagSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 新增数据
	var tagOptionDto *response.TagOptionDto
	tagOptionDto, err = tagService.SaveTag(tagSaveVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, tagOptionDto)
}

func AdminTagRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	var res bool
	// 更新记录
	res, err = tagService.RemoveTagById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminTagUpdate(c *gin.Context) {
	// 绑定并检查参数
	var tagUpdateVo request.TagUpdateVo
	var err error
	err = c.ShouldBind(&tagUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	tagUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	var res bool
	// 更新记录
	res, err = tagService.UpdateTag(tagUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminTagUpdateDeleted(c *gin.Context) {
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
	res, err = tagService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
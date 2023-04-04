package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminCategoryPage(c *gin.Context) {
	// 绑定并检查参数
	var pageVo request.PageVo
	var categoryVo request.CategoryVo
	var err error
	err = c.ShouldBind(&pageVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	if res, _ := request.ValidatePageVo(pageVo); !res {
		pageVo.New()
	}
	err = c.ShouldBind(&categoryVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询分页
	var page *response.Page
	page, err = categoryService.CategoryDtoPage(pageVo, categoryVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, page)
}

func AdminCategoryQuery(c *gin.Context) {
	name := c.Query("name")
	list, err := categoryService.QueryCategoryByName(name)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminCategoryOption(c *gin.Context) {
	list, err := categoryService.CategoryOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminCategorySave(c *gin.Context) {
	// 绑定并检查参数
	var categorySaveVo request.CategorySaveVo
	var err error
	err = c.ShouldBind(&categorySaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 新增数据
	var categoryOptionDto *response.CategoryOptionDto
	categoryOptionDto, err = categoryService.SaveCategory(categorySaveVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, categoryOptionDto)
}

func AdminCategoryRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	var res bool
	// 更新记录
	res, err = categoryService.RemoveCategoryById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminCategoryUpdate(c *gin.Context) {
	// 绑定并检查参数
	var categoryUpdateVo request.CategoryUpdateVo
	var err error
	err = c.ShouldBind(&categoryUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	categoryUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}

	// 更新记录
	var res bool
	res, err = categoryService.UpdateCategory(categoryUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}

func AdminCategoryUpdateDeleted(c *gin.Context) {
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
	res, err = categoryService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, res)
}
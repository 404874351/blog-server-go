package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
)

func CategoryOption(c *gin.Context) {
	list, err := categoryService.CategoryOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}
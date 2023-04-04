package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitCategoryRouter
//  @Description: 初始化分类路由
//  @param r
//
func InitCategoryRouter(r *gin.Engine) {
	categoryGroup := r.Group("/category")

	categoryGroup.GET("/option", api.CategoryOption)
}

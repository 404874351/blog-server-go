package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminCategoryRouter
//  @Description: 初始化控制台分类路由
//  @param r
//
func InitAdminCategoryRouter(r *gin.Engine) {
	adminCategoryGroup := r.Group("/admin/category")

	adminCategoryGroup.GET("/page", api.AdminCategoryPage)
	adminCategoryGroup.GET("/query", api.AdminCategoryQuery)
	adminCategoryGroup.GET("/option", api.AdminCategoryOption)
	adminCategoryGroup.POST("/save", api.AdminCategorySave)
	adminCategoryGroup.POST("/remove/:id", api.AdminCategoryRemove)
	adminCategoryGroup.POST("/update/:id", api.AdminCategoryUpdate)
	adminCategoryGroup.POST("/update/:id/deleted", api.AdminCategoryUpdateDeleted)
}

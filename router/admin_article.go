package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminArticleRouter
//  @Description: 初始化控制台文章路由
//  @param r
//
func InitAdminArticleRouter(r *gin.Engine) {
	adminArticleGroup := r.Group("/admin/article")

	adminArticleGroup.GET("/page", api.AdminArticlePage)
	adminArticleGroup.GET("/:id", api.AdminArticleDetail)
	adminArticleGroup.POST("/save", api.AdminArticleSave)
	adminArticleGroup.POST("/remove/:id", api.AdminArticleRemove)
	adminArticleGroup.POST("/update/:id", api.AdminArticleUpdate)
	adminArticleGroup.POST("/update/:id/deleted", api.AdminArticleUpdateDeleted)
}

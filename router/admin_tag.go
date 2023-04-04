package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminTagRouter
//  @Description: 初始化控制台标签路由
//  @param r
//
func InitAdminTagRouter(r *gin.Engine) {
	adminTagGroup := r.Group("/admin/tag")

	adminTagGroup.GET("/page", api.AdminTagPage)
	adminTagGroup.GET("/query", api.AdminTagQuery)
	adminTagGroup.GET("/option", api.AdminTagOption)
	adminTagGroup.POST("/save", api.AdminTagSave)
	adminTagGroup.POST("/remove/:id", api.AdminTagRemove)
	adminTagGroup.POST("/update/:id", api.AdminTagUpdate)
	adminTagGroup.POST("/update/:id/deleted", api.AdminTagUpdateDeleted)
}

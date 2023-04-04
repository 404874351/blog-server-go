package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminMessageRouter
//  @Description: 初始化控制台留言路由
//  @param r
//
func InitAdminMessageRouter(r *gin.Engine) {
	adminMessageGroup := r.Group("/admin/message")

	adminMessageGroup.GET("/page", api.AdminMessagePage)
	adminMessageGroup.POST("/remove/:id", api.AdminMessageRemove)
	adminMessageGroup.POST("/update/:id/deleted", api.AdminMessageUpdateDeleted)
}

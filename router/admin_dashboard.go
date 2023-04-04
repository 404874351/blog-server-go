package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminDashboardRouter
//  @Description: 初始化控制台仪表路由
//  @param r
//
func InitAdminDashboardRouter(r *gin.Engine) {
	adminDashboardGroup := r.Group("/admin/dashboard")

	adminDashboardGroup.GET("/index", api.AdminDashboardIndex)
	adminDashboardGroup.GET("/view", api.AdminDashboardView)
	adminDashboardGroup.GET("/role", api.AdminDashboardRole)
	adminDashboardGroup.GET("/category", api.AdminDashboardCategory)
	adminDashboardGroup.GET("/tag", api.AdminDashboardTag)

}

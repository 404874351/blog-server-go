package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminUserRouter
//  @Description: 初始化控制台用户路由
//  @param r
//
func InitAdminUserRouter(r *gin.Engine) {
	adminUserGroup := r.Group("/admin/user")

	adminUserGroup.GET("/page", api.AdminUserPage)
	adminUserGroup.POST("/update/:id", api.AdminUserUpdate)
	adminUserGroup.POST("/update/:id/deleted", api.AdminUserUpdateDeleted)
	adminUserGroup.POST("/remove/:id", api.AdminUserRemove)
}
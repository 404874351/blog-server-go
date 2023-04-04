package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminPermissionRouter
//  @Description: 初始化控制台权限路由
//  @param r
//
func InitAdminPermissionRouter(r *gin.Engine) {
	adminPermissionGroup := r.Group("/admin/permission")

	adminPermissionGroup.GET("/list", api.AdminPermissionList)
	adminPermissionGroup.GET("/option", api.AdminPermissionOption)
	adminPermissionGroup.POST("/save", api.AdminPermissionSave)
	adminPermissionGroup.POST("/remove/:id", api.AdminPermissionRemove)
	adminPermissionGroup.POST("/update/:id", api.AdminPermissionUpdate)
	adminPermissionGroup.POST("/update/:id/deleted", api.AdminPermissionUpdateDeleted)
}

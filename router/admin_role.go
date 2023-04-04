package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminRoleRouter
//  @Description: 初始化控制台角色路由
//  @param r
//
func InitAdminRoleRouter(r *gin.Engine) {
	adminRoleGroup := r.Group("/admin/role")

	adminRoleGroup.GET("/page", api.AdminRolePage)
	adminRoleGroup.GET("/option", api.AdminRoleOption)
	adminRoleGroup.POST("/save", api.AdminRoleSave)
	adminRoleGroup.POST("/remove/:id", api.AdminRoleRemove)
	adminRoleGroup.POST("/update/:id", api.AdminRoleUpdate)
	adminRoleGroup.POST("/update/:id/deleted", api.AdminRoleUpdateDeleted)
	adminRoleGroup.POST("/update/:id/menu", api.AdminRoleUpdateMenu)
	adminRoleGroup.POST("/update/:id/permission", api.AdminRoleUpdatePermission)
}

package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminMenuRouter
//  @Description: 初始化控制台菜单路由
//  @param r
//
func InitAdminMenuRouter(r *gin.Engine) {
	adminMenuGroup := r.Group("/admin/menu")

	adminMenuGroup.GET("/list", api.AdminMenuList)
	adminMenuGroup.GET("/option", api.AdminMenuOption)
	adminMenuGroup.GET("/user", api.AdminMenuUserTree)
	adminMenuGroup.POST("/save", api.AdminMenuSave)
	adminMenuGroup.POST("/remove/:id", api.AdminMenuRemove)
	adminMenuGroup.POST("/update/:id", api.AdminMenuUpdate)
	adminMenuGroup.POST("/update/:id/deleted", api.AdminMenuUpdateDeleted)
}

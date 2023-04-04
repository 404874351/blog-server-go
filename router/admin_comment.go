package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitAdminCommentRouter
//  @Description: 初始化控制台评论路由
//  @param r
//
func InitAdminCommentRouter(r *gin.Engine) {
	adminCommentGroup := r.Group("/admin/comment")

	adminCommentGroup.GET("/page", api.AdminCommentPage)
	adminCommentGroup.POST("/remove/:id", api.AdminCommentRemove)
	adminCommentGroup.POST("/update/:id", api.AdminCommentUpdate)
	adminCommentGroup.POST("/update/:id/deleted", api.AdminCommentUpdateDeleted)

}

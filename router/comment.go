package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitCommentRouter
//  @Description: 初始化评论路由
//  @param r
//
func InitCommentRouter(r *gin.Engine) {
	commentGroup := r.Group("/comment")

	commentGroup.GET("/page", api.CommentPage)
	commentGroup.POST("/save", api.CommentSave)
	commentGroup.POST("/praise/:id", api.CommentPraise)
	commentGroup.POST("/cancel_praise/:id", api.CommentCancelPraise)
	commentGroup.POST("/remove/:id", api.CommentRemove)
}

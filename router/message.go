package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitMessageRouter
//  @Description: 初始化留言路由
//  @param r
//
func InitMessageRouter(r *gin.Engine) {
	messageGroup := r.Group("/message")

	messageGroup.GET("/count", api.MessageCount)
	messageGroup.GET("/page", api.MessagePage)
	messageGroup.POST("/save", api.MessageSave)

}

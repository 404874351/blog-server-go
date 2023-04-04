package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitLoginRouter
//  @Description: 初始化登录路由
//  @param r
//
func InitLoginRouter(r *gin.Engine) {
	userGroup := r.Group("/login")

	userGroup.POST("", api.UsernamePasswordLogin)
	userGroup.POST("/code", api.AuthCodeLogin)

	//userGroup.POST("/test", func (c *gin.Context) {
	//
	//})
}

//
// InitLogoutRouter
//  @Description: 初始化登出路由
//  @param r
//
func InitLogoutRouter(r *gin.Engine) {
	r.POST("/logout", api.Logout)
}

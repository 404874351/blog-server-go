package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitUserRouter
//  @Description: 初始化用户路由
//  @param r
//
func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")

	userGroup.GET("/info", api.UserInfo)
	userGroup.POST("/code", api.UserAuthCode)
	userGroup.POST("/register", api.UserRegister)
	userGroup.POST("/password", api.UserUpdatePassword)
	userGroup.POST("/update", api.UserUpdate)
}

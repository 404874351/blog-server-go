package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitLoginRouter
//  @Description: 初始化七牛云路由
//  @param r
//
func InitQiniuRouter(r *gin.Engine) {
	userGroup := r.Group("/qiniu")

	userGroup.GET("/token", api.QiniuToken)
}

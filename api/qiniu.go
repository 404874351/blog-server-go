package api

import (
	"blog-server-go/middleware"
	"github.com/gin-gonic/gin"
)

//
// QiniuToken
//  @Description: 获取七牛云文件上传凭证
//  @param c
//
func QiniuToken(c *gin.Context) {
	key := c.Query("key")
	token := qiniuService.CreateToken(key)
	middleware.SetData(c, token)
}

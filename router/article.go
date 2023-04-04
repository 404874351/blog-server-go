package router

import (
	"blog-server-go/api"
	"github.com/gin-gonic/gin"
)

//
// InitArticleRouter
//  @Description: 初始化文章路由
//  @param r
//
func InitArticleRouter(r *gin.Engine) {
	articleGroup := r.Group("/article")

	articleGroup.GET("/statistic", api.ArticleStatistic)
	articleGroup.GET("/page", api.ArticlePage)
	articleGroup.GET("/:id", api.ArticleDetail)
	articleGroup.POST("/praise/:id", api.ArticlePraise)
	articleGroup.POST("/cancel_praise/:id", api.ArticleCancelPraise)
}

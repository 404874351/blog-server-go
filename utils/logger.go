package utils

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func BindZapLogger(r *gin.Engine) {
	// 获取config对象，区分开发和生产模式
	conf := zap.NewDevelopmentConfig()
	//conf := zap.NewProductionConfig()
	// 自定义修改config对象的属性
	//conf.Encoding = "json"
	// 通过config对象得到logger对象指针
	logger, _ := conf.Build()
	// 替换掉全局的logger，以后都使用zap.L()
	zap.ReplaceGlobals(logger)
	// 替换gin默认的Logger中间件
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

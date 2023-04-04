package main

import (
	"blog-server-go/conf"
	"blog-server-go/router"
	"go.uber.org/zap"
	"strconv"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

var serverConfig = conf.GlobalConfig.Server

func main() {
	// 绑定路由
	r := router.InitRouter()
	// 启动服务
	err := r.Run(":" + strconv.FormatInt(int64(serverConfig.Port), 10))

	if err != nil {
		zap.L().Error(err.Error())
		zap.L().Error("server start failed")
	}
}


package router

import (
	"blog-server-go/middleware"
	"blog-server-go/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置工作模式，默认debug模式
	// 部署上线时务必开启release模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化引擎
	r := gin.New()

	// 使用zap输出日志
	utils.BindZapLogger(r)

	// 注册全局中间件，注意有先后顺序
	r.Use(middleware.Recovery())
	r.Use(middleware.Jwt())
	// 测试接口时，可关闭权限控制
	r.Use(middleware.Casbin())

	// 添加不同模块的路由
	// 登录模块
	InitLoginRouter(r)
	InitLogoutRouter(r)
	// 控制台模块
	InitAdminArticleRouter(r)
	InitAdminCategoryRouter(r)
	InitAdminCommentRouter(r)
	InitAdminDashboardRouter(r)
	InitAdminMenuRouter(r)
	InitAdminMessageRouter(r)
	InitAdminPermissionRouter(r)
	InitAdminRoleRouter(r)
	InitAdminTagRouter(r)
	InitAdminUserRouter(r)
	// 客户端模块
	InitArticleRouter(r)
	InitCategoryRouter(r)
	InitCommentRouter(r)
	InitMessageRouter(r)
	InitTagRouter(r)
	InitUserRouter(r)
	// 其他服务模块
	InitQiniuRouter(r)

	return r
}

package service

import (
	"blog-server-go/model"
)

// 打包业务层
var GlobalService = struct {
	ArticleService
	AuthCodeService
	CategoryService
	CommentService
	JwtService
	MenuService
	MessageService
	PermissionService
	QiniuService
	RoleService
	TagService
	UserService
}{}

// 统一获取数据库对象
var db = model.DB

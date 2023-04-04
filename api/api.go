package api

import "blog-server-go/service"

var articleService      = service.GlobalService.ArticleService
var authCodeService     = service.GlobalService.AuthCodeService
var categoryService     = service.GlobalService.CategoryService
var commentService      = service.GlobalService.CommentService
var jwtService          = service.GlobalService.JwtService
var menuService         = service.GlobalService.MenuService
var messageService      = service.GlobalService.MessageService
var permissionService   = service.GlobalService.PermissionService
var qiniuService        = service.GlobalService.QiniuService
var roleService         = service.GlobalService.RoleService
var tagService          = service.GlobalService.TagService
var userService         = service.GlobalService.UserService
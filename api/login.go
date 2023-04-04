package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//
// UsernamePasswordLogin
//  @Description: 普通用户名密码登录
//  @param c
//
func UsernamePasswordLogin(c *gin.Context) {
	// 绑定登录参数
	var loginVo request.UsernamePasswordLoginVo
	var err error
	err = c.ShouldBind(&loginVo)
	// 响应参数错误
	if err != nil {
		if loginVo.Username == "" {
			middleware.ReportError(c, response.USERNAME_NULL, err)
		} else if loginVo.Password == "" {
			middleware.ReportError(c, response.USERNAME_OR_PASSWORD_ERROR, err)
		} else {
			middleware.ReportError(c, response.LOGIN_FAIL, err)
		}
	}
	// 获取用户数据，核对密码和停用状态
	var user *model.User
	user, err = userService.GetUserByUsername(loginVo.Username)
	if err != nil || user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVo.Password)) != nil {
		middleware.ReportError(c, response.USERNAME_OR_PASSWORD_ERROR, err)
	}
	if user.Deleted == model.MODEL_DEACTIVED {
		middleware.ReportError(c, response.USER_DEACTIVATED, err)
	}
	// 认证通过，登录成功，开始后续处理
	onLoginSuccess(c, user)
}

//
// AuthCodeLogin
//  @Description: 短信验证码登录
//  @param c
//
func AuthCodeLogin(c *gin.Context) {
	// 绑定登录参数
	var loginVo request.AuthCodeLoginVo
	var err error
	err = c.ShouldBind(&loginVo)
	// 响应参数错误
	if err != nil {
		if loginVo.Username == "" {
			middleware.ReportError(c, response.USERNAME_NULL, err)
		} else if loginVo.Code == "" {
			middleware.ReportError(c, response.AUTH_CODE_ERROR, err)
		} else {
			middleware.ReportError(c, response.LOGIN_FAIL, err)
		}
	}
	// 获取用户数据，核对验证码和停用状态
	var user *model.User
	user, err = userService.GetUserByUsername(loginVo.Username)
	if err != nil || user == nil {
		middleware.ReportError(c, response.USER_NOT_EXIST, err)
	}
	var codeInRedis string
	codeInRedis, err = authCodeService.GetAuthCodeInRedis(loginVo.Username)
	if err != nil || codeInRedis != loginVo.Code {
		middleware.ReportError(c, response.AUTH_CODE_ERROR, err)
	}
	if user.Deleted == model.MODEL_DEACTIVED {
		middleware.ReportError(c, response.USER_DEACTIVATED, err)
	}
	// 认证通过，清空验证码redis缓存
	err = authCodeService.DelAuthCodeInRedis(loginVo.Username)
	if err != nil {
		middleware.ReportError(c, response.LOGIN_FAIL, err)
	}
	// 登录成功，开始后续处理
	onLoginSuccess(c, user)
}

//
// Logout
//  @Description: 用户登出
//  @param c
//
func Logout(c *gin.Context) {
	// 通过token，获取username
	claims := middleware.GetClaims(c)
	// 删除授权的用户-角色关系
	username := claims.Subject
	err := middleware.ClearUserGroup(username)
	if err != nil {
		middleware.ReportError(c, response.ACCESS_FAILED, err)
	}
	// redis删除token
	err = jwtService.DelTokenInRedis(username)
	if err != nil {
		middleware.ReportError(c, response.ACCESS_FAILED, err)
	}

	middleware.SetData(c, true)
}

//
// onLoginSuccess
//  @Description: 登录成功的后续处理
//  @param c
//
func onLoginSuccess(c *gin.Context, user *model.User) {
	// 生成token
	token, err := jwtService.CreateToken(*user)
	if err != nil {
		middleware.ReportError(c, response.LOGIN_FAIL, err)
	}
	// 存储用户-角色关系，用于授权
	var roles []string
	for _, role := range user.RoleList {
		roles = append(roles, role.Code)
	}
	err = middleware.AddUserGroup(user.Username, roles)
	if err != nil {
		middleware.ReportError(c, response.LOGIN_FAIL, err)
	}
	// redis存储token
	err = jwtService.SetTokenInRedis(user.Username, token)
	if err != nil {
		middleware.ReportError(c, response.LOGIN_FAIL, err)
	}

	middleware.SetData(c, gin.H{
		"token": token,
	})
}


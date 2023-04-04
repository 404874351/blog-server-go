package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
)

//
// UserInfo
//  @Description: 获取用户个人信息
//  @param c
//
func UserInfo(c *gin.Context) {
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	userInfo, err := userService.GetUserInfoDtoById(claims.ID)
	if err != nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, err)
	}
	middleware.SetData(c, userInfo)
}

//
// UserAuthCode
//  @Description: 用户获取验证码
//  @param c
//
func UserAuthCode(c *gin.Context) {
	// 绑定并检查参数
	var authCodeVo request.UserAuthCodeVo
	var err error
	err = c.ShouldBind(&authCodeVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 查询该用户是否存在
	username := authCodeVo.Phone
	_, err = userService.GetUserByUsername(username)
	if err != nil {
		middleware.ReportError(c, response.USER_NOT_EXIST, err)
	}
	// 查询是否可重新发送
	if !authCodeService.CanSend(username) {
		middleware.ReportError(c, response.AUTH_CODE_SEND_TOO_FAST, err)
	}
	// 发送验证码，并存入redis
	code := authCodeService.CreateAuthCode()
	err = authCodeService.SetAuthCodeInRedis(username, code)
	if err != nil {
		middleware.ReportError(c, response.AUTH_CODE_SEND_ERROR, err)
	}
	err = authCodeService.Send(username, code)
	if err != nil {
		middleware.ReportError(c, response.AUTH_CODE_SEND_ERROR, err)
	}
	middleware.SetData(c, true)
}

//
// UserRegister
//  @Description: 用户注册
//  @param c
//
func UserRegister(c *gin.Context) {
	// 绑定并检查参数
	var registerVo request.UserRegisterVo
	var err error
	err = c.ShouldBind(&registerVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 检查用户是否已注册
	var user *model.User
	user, err = userService.GetUserByUsername(registerVo.Phone)
	if err != nil || user != nil {
		middleware.ReportError(c, response.PHONE_EXIST, err)
	}
	// 核实并删除验证码
	var codeInRedis string
	codeInRedis, err = authCodeService.GetAuthCodeInRedis(registerVo.Phone)
	if err != nil || codeInRedis != registerVo.Code {
		middleware.ReportError(c, response.AUTH_CODE_INVALID, err)
	}
	err = authCodeService.DelAuthCodeInRedis(registerVo.Phone)
	if err != nil {
		middleware.ReportError(c, response.ACCESS_FAILED, err)
	}
	// 完成注册
	var res bool
	res, err = userService.UserRegister(registerVo)
	if err != nil || !res {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

//
// UserUpdatePassword
//  @Description: 修改密码
//  @param c
//
func UserUpdatePassword(c *gin.Context) {
	// 绑定并检查参数
	var userPasswordVo request.UserPasswordVo
	var err error
	err = c.ShouldBind(&userPasswordVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 从登录状态中获取用户id
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	userPasswordVo.ID = claims.ID
	// 更新记录
	var res bool
	res, err = userService.UpdatePassword(userPasswordVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

//
// UserUpdate
//  @Description: 更新用户信息
//  @param c
//
func UserUpdate(c *gin.Context) {
	// 绑定并检查参数
	var userUpdateVo request.UserUpdateVo
	var err error
	err = c.ShouldBind(&userUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 从登录状态中获取用户id
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	userUpdateVo.ID = claims.ID
	// 更新记录
	var res bool
	res, err = userService.UpdateUser(userUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

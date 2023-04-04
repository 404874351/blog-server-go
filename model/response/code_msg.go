package response

type StateCode int

const (
	// 1xxx 请求成功
	SUCCESS 					StateCode = 1000

	// 2xxx 登录控制
	LOGIN_FAIL 					StateCode = 2000
	USERNAME_NULL 				StateCode = 2001
	USER_DEACTIVATED 			StateCode = 2002
	USERNAME_OR_PASSWORD_ERROR 	StateCode = 2003
	AUTH_CODE_ERROR 			StateCode = 2004

	// 3xxx 用户认证异常
	AUTHENTICATION_FAILED 		StateCode = 3000
	TOKEN_NOT_EXIST 			StateCode = 3001
	TOKEN_ILLEGAL 				StateCode = 3002
	TOKEN_INVALID 				StateCode = 3003

	// 4xxx 业务访问异常
	ACCESS_FAILED 				StateCode = 4000
	ACCESS_DENIED 				StateCode = 4001
	PARAMETER_ILLEGAL 			StateCode = 4002
	SQL_FAILED 					StateCode = 4003
	SQL_UNIQUE_ERROR 			StateCode = 4004
	SQL_INTEGRITY_ERROR 		StateCode = 4005
	RAW_PASSWORD_ERROR 			StateCode = 4006
	USER_NOT_EXIST				StateCode = 4007
	PHONE_EXIST 				StateCode = 4008
	COMMENT_TOP_EXIST 			StateCode = 4009
	COMMENT_NOT_LEVEL_TOP 		StateCode = 4010
	AUTH_CODE_SEND_TOO_FAST 	StateCode = 4011
	AUTH_CODE_INVALID 			StateCode = 4012
	AUTH_CODE_SEND_ERROR 		StateCode = 4013
)

var MsgMap = map[StateCode]string {
	// 1xxx 请求成功
	SUCCESS 					: "请求成功",

	// 2xxx 登录控制
	LOGIN_FAIL 					: "登录流程异常，请联系管理员",
	USERNAME_NULL 				: "用户名为空",
	USER_DEACTIVATED 			: "该账号已停用，请联系管理员",
	USERNAME_OR_PASSWORD_ERROR 	: "用户名或密码不正确",
	AUTH_CODE_ERROR 			: "用户名或验证码不正确",

	// 3xxx 用户认证异常
	AUTHENTICATION_FAILED 		: "用户认证流程异常，请联系管理员",
	TOKEN_NOT_EXIST 			: "Token不存在，请先登录",
	TOKEN_ILLEGAL 				: "Token非法或过期，请重新登录",
	TOKEN_INVALID 				: "Token已失效，请重新登录",

	// 4xxx 业务访问异常
	ACCESS_FAILED 				: "处理流程异常，访问失败",
	ACCESS_DENIED 				: "权限不足，拒绝访问",
	PARAMETER_ILLEGAL 			: "请求参数非法，请检查数据合法性",
	SQL_FAILED 					: "数据库操作失败，请检查逻辑错误",
	SQL_UNIQUE_ERROR 			: "数据库字段唯一性错误",
	SQL_INTEGRITY_ERROR 		: "数据库字段完整性错误",
	RAW_PASSWORD_ERROR 			: "原密码不正确",
	USER_NOT_EXIST				: "用户不存在",
	PHONE_EXIST 				: "该手机号已被他人注册",
	COMMENT_TOP_EXIST 			: "该文章的置顶评论已存在且唯一",
	COMMENT_NOT_LEVEL_TOP 		: "该评论不是顶级评论",
	AUTH_CODE_SEND_TOO_FAST 	: "验证码发送过于频繁",
	AUTH_CODE_INVALID 			: "验证码错误或已失效",
	AUTH_CODE_SEND_ERROR 		: "验证码发送失败，请稍后再试",
}


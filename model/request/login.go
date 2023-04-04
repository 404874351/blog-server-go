package request

//
// UsernamePasswordLoginVo
//  @Description: 用户名密码登录 请求对象
//
type UsernamePasswordLoginVo struct {
	// 用户名
	Username	string	`json:"username"    form:"username"    binding:"required"`
	// 密码
	Password	string	`json:"password"    form:"password"    binding:"required"`
}

//
// AuthCodeLoginVo
//  @Description: 验证码登录 请求对象
//
type AuthCodeLoginVo struct {
	// 用户名
	Username	string	`json:"username"    form:"username"    binding:"required"`
	// 验证码
	Code	    string	`json:"code"        form:"code"        binding:"required"`
}


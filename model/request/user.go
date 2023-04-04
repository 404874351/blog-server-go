package request

//
// UserVo
//  @Description: 用户 请求对象
//
type UserVo struct {
	// 昵称
	Nickname	string	`json:"nickname"    form:"nickname"`
	// 手机号
	Phone 		string	`json:"phone"       form:"phone"`
}

//
// UserUpdateVo
//  @Description: 用户更新 请求对象
//
type UserUpdateVo struct {
	// id，只允许用户修改自己的密码，因此登录后自动获取
	ID        	int		`json:"id"          form:"id"`
	// 昵称
	Nickname	string	`json:"nickname"    form:"nickname"`
	// 手机号
	Phone 		string	`json:"phone"       form:"phone"`
	// 头像
	AvatarUrl	string	`json:"avatarUrl"   form:"avatarUrl"`
}

//
// UserRegisterVo
//  @Description: 用户注册 请求对象
//
type UserRegisterVo struct {
	// 昵称
	Nickname	string	`json:"nickname"    form:"nickname"      binding:"required"`
	// 手机号
	Phone 		string	`json:"phone"       form:"phone"         binding:"required,ValidatePhone"`
	// 密码
	Password	string	`json:"password"    form:"password"      binding:"required,ValidatePassword"`
	// 验证码
	Code 		string	`json:"code"        form:"code"          binding:"required"`
}

//
// UserPasswordVo
//  @Description: 用户密码 请求对象
//
type UserPasswordVo struct {
	// id，只允许用户修改自己的密码，因此登录后自动获取
	ID        	int		`json:"id"          form:"id"`
	// 密码
	Password	string	`json:"password"    form:"password"      binding:"required"`
	// 密码
	NewPassword	string	`json:"newPassword" form:"newPassword"   binding:"required,ValidatePassword"`
}

//
// UserAuthCodeVo
//  @Description: 用户请求验证码 请求对象
//
type UserAuthCodeVo struct {
	// 手机号
	Phone 		string	`json:"phone"       form:"phone"         binding:"required,ValidatePhone"`
}

//
// UserAdminUpdateVo
//  @Description: 用户后台更新 请求对象
//
type UserAdminUpdateVo struct {
	// id
	ID        	int		`json:"id"          form:"id"`
	// 昵称
	Nickname	string	`json:"nickname"    form:"nickname"`
	// 角色id列表
	RoleIdList  []int   `json:"roleIdList"  form:"roleIdList"    binding:"required"`
}

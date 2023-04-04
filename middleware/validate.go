package middleware

import (
	"blog-server-go/model"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var phonePattern *regexp2.Regexp
var passwordPattern *regexp2.Regexp

func init() {
	var err error
	phonePattern, err = regexp2.Compile("^1[3-9]\\d{9}$", 0)
	if err != nil {
		panic(err)
	}
	passwordPattern, err = regexp2.Compile("^(?=.*[0-9])(?=.*[a-zA-Z])(?=.*[!@#$_&*+-])[0-9a-zA-Z!@#$_&*+-]{8,18}$", 0)
	if err != nil {
		panic(err)
	}

	// 注册自定义验证器
	InitValidate()
}

func InitValidate() {
	var err error
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("ValidatePhone", ValidatePhone)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidatePassword", ValidatePassword)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateDeleted", ValidateDeleted)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateCommentTop", ValidateCommentTop)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidatePermissionType", ValidatePermissionType)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidatePermissionLevel", ValidatePermissionLevel)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidatePermissionAnonymous", ValidatePermissionAnonymous)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateMenuType", ValidateMenuType)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateMenuLevel", ValidateMenuLevel)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateMenuHidden", ValidateMenuHidden)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateArticleTop", ValidateArticleTop)
		if err != nil {
			panic(err)
		}
		err = v.RegisterValidation("ValidateArticleCloseComment", ValidateArticleCloseComment)
		if err != nil {
			panic(err)
		}
	}
}


//
// ValidatePhone
//  @Description: 验证手机号格式，暂时仅验证长度等少量信息
//  @param fl
//  @return bool
//
func ValidatePhone(fl validator.FieldLevel) bool {
	phone, ok := fl.Field().Interface().(string)
	if ok {
		res, err := phonePattern.MatchString(phone)
		if err == nil {
			return res
		}
	}
	return false
}

//
// ValidatePassword
//  @Description: 验证密码格式，密码必须包含数字、字母和特殊字符，长度为8-18位
//  @param fl
//  @return bool
//
func ValidatePassword(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		res, err := passwordPattern.MatchString(password)
		if err == nil {
			return res
		}
	}
	return false
}

func ValidateDeleted(fl validator.FieldLevel) bool {
	deleted, ok := fl.Field().Interface().(int8)
	if ok {
		return deleted == model.MODEL_ACTIVED || deleted == model.MODEL_DEACTIVED
	}
	return false
}

func ValidateCommentTop(fl validator.FieldLevel) bool {
	top, ok := fl.Field().Interface().(int8)
	if ok {
		return top == model.COMMENT_TOP_DISABLE || top == model.COMMENT_TOP_ENABLE
	}
	return false
}

func ValidatePermissionType(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(int8)
	if ok {
		return t == model.PERMISSION_TYPE_ITEM || t == model.PERMISSION_TYPE_GROUP
	}
	return false
}

func ValidatePermissionLevel(fl validator.FieldLevel) bool {
	level, ok := fl.Field().Interface().(int8)
	if ok {
		return level >= model.PERMISSION_LEVEL_TOP
	}
	return false
}

func ValidatePermissionAnonymous(fl validator.FieldLevel) bool {
	anonymous, ok := fl.Field().Interface().(int8)
	if ok {
		return anonymous == model.PERMISSION_ANONYMOUS_DISABLE || anonymous == model.PERMISSION_ANONYMOUS_ENABLE
	}
	return false
}

func ValidateMenuType(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(int8)
	if ok {
		return t == model.MENU_TYPE_ITEM || t == model.MENU_TYPE_GROUP
	}
	return false
}

func ValidateMenuLevel(fl validator.FieldLevel) bool {
	level, ok := fl.Field().Interface().(int8)
	if ok {
		return level >= model.MENU_LEVEL_TOP
	}
	return false
}

func ValidateMenuHidden(fl validator.FieldLevel) bool {
	hidden, ok := fl.Field().Interface().(int8)
	if ok {
		return hidden == model.MENU_HIDDEN_DISABLE || hidden == model.MENU_HIDDEN_ENABLE
	}
	return false
}

func ValidateArticleTop(fl validator.FieldLevel) bool {
	top, ok := fl.Field().Interface().(int8)
	if ok {
		return top == model.ARTICLE_TOP_DISABLE || top == model.ARTICLE_TOP_ENABLE
	}
	return false
}

func ValidateArticleCloseComment(fl validator.FieldLevel) bool {
	closeComment, ok := fl.Field().Interface().(int8)
	if ok {
		return closeComment == model.ARTICLE_CLOSE_COMMENT_DISABLE || closeComment == model.ARTICLE_CLOSE_COMMENT_ENABLE
	}
	return false
}

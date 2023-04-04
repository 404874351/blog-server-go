package service

import "blog-server-go/utils"

type AuthCodeService struct {}

const (
	AUTH_CODE_EXPIRED_TIME = 300
	AUTH_CODE_RETRY_TIME   = 60
)

//
// CreateAuthCode
//  @Description: 随机生成6位验证码
//  @receiver a
//  @return string
//
func (a *AuthCodeService) CreateAuthCode() string {
	return utils.RandomNumbers(6)
}

//
// Send
//  @Description: 发送验证码
//  @receiver a
//  @param phone
//  @param code
//  @return error
//
func (a *AuthCodeService) Send(phone string, code string) error {
	return utils.SmsSendAuthCode(phone, code)
}

//
// CanSend
//  @Description: 判断是否可重发验证码
//  @receiver a
//  @param username
//  @return bool
//
func (a *AuthCodeService) CanSend(username string) bool {
	value, err := utils.RedisGet(username + ":code-retry")
	if err != nil {
		return false
	}
	return value != "1"
}

//
// GetAuthCodeInRedis
//  @Description: redis获取验证码
//  @receiver a
//  @param username
//  @return string
//  @return error
//
func (a *AuthCodeService) GetAuthCodeInRedis(username string) (string, error) {
	return utils.RedisGet(username + ":code")
}

//
// SetAuthCodeInRedis
//  @Description: redis设置验证码
//  @receiver a
//  @param username
//  @param code
//  @return error
//
func (a *AuthCodeService) SetAuthCodeInRedis(username string, code string) error {
	var err error
	err = utils.RedisSet(username + ":code", code, AUTH_CODE_EXPIRED_TIME)
	if err != nil {
		return err
	}
	err = utils.RedisSet(username + ":code-retry", "1", AUTH_CODE_RETRY_TIME)
	if err != nil {
		return err
	}
	return nil
}

//
// DelAuthCodeInRedis
//  @Description: redis删除验证码
//  @receiver a
//  @param username
//  @return error
//
func (a *AuthCodeService) DelAuthCodeInRedis(username string) error {
	var err error
	err = utils.RedisDel(username + ":code")
	if err != nil {
		return err
	}
	err = utils.RedisDel(username + ":code-retry")
	if err != nil {
		return err
	}
	return nil
}

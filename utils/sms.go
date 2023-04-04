package utils

import (
	"blog-server-go/conf"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var smsConfig = conf.GlobalConfig.SMS

// sms客户端
var smsClient *dysmsapi20170525.Client

func init() {
	config := &openapi.Config{
		AccessKeyId: &smsConfig.AccessKeyId,
		AccessKeySecret: &smsConfig.AccessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	smsClient = &dysmsapi20170525.Client{}
	var err error
	smsClient, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		panic(err)
	}
}

func SmsSendAuthCode(phone string, code string) (_err error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:	tea.String(phone),
		SignName:     	tea.String(smsConfig.SignName),
		TemplateCode: 	tea.String(smsConfig.TemplateCode),
		TemplateParam:	tea.String(`{"code":"` + code + `"}`),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = smsClient.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}
		return nil
	}()
	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(err.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

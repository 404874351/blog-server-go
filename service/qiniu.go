package service

import (
	"blog-server-go/conf"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuService struct {}

var qiniuConfig = conf.GlobalConfig.Qiniu

const EXPIRE_SECONDS = 3600

//
// CreateToken
//  @Description: 构造文件上传token
//  @receiver a
//  @param key 覆盖上传所需的key
//  @return string
//  @return error
//
func (a *QiniuService) CreateToken(key string) string {
	var putPolicy storage.PutPolicy
	if key != "" {
		putPolicy = storage.PutPolicy{
			Scope: fmt.Sprintf("%s:%s", qiniuConfig.Bucket, key),
		}
	} else {
		putPolicy = storage.PutPolicy{
			Scope: qiniuConfig.Bucket,
		}
	}
	mac := qbox.NewMac(qiniuConfig.AccessKey, qiniuConfig.SecretKey)
	token := putPolicy.UploadToken(mac)
	return token
}
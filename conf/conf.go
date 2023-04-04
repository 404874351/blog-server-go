package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server 		Server 	`mapstructure:"server"      json:"server"      yaml:"server"`
	Mysql 		Mysql 	`mapstructure:"mysql"       json:"mysql"       yaml:"mysql"`
	Redis 		Redis 	`mapstructure:"redis"       json:"redis"       yaml:"redis"`
	Jwt         JWT     `mapstructure:"jwt"         json:"jwt"         yaml:"jwt"`
	Qiniu      	Qiniu   `mapstructure:"qiniu"       json:"qiniu"       yaml:"qiniu"`
	SMS      	SMS     `mapstructure:"aliyun-sms"  json:"aliyun-sms"  yaml:"aliyun-sms"`
}

type Server struct {
	Name 		string 	`mapstructure:"name"        json:"name"        yaml:"name"`
	Port 		int		`mapstructure:"port"        json:"port"        yaml:"port"`
}

type Mysql struct {
	Url 		string 	`mapstructure:"url"         json:"url"         yaml:"url"`
	Username 	string 	`mapstructure:"username"    json:"username"    yaml:"username"`
	Password 	string 	`mapstructure:"password"    json:"password"    yaml:"password"`
}

type Redis struct {
	Host        string  `mapstructure:"host"        json:"host"        yaml:"host"`
	Port        string  `mapstructure:"port"        json:"port"        yaml:"port"`
	Password    string  `mapstructure:"password"    json:"password"    yaml:"password"`
}

type JWT struct {
	MaxAge      int 	`mapstructure:"max-age"     json:"max-age"     yaml:"max-age"`
	Secret      string 	`mapstructure:"secret"      json:"secret"      yaml:"secret"`
}

type Qiniu struct {
	AccessKey	string  `mapstructure:"access-key"  json:"access-key"  yaml:"access-key"`
	SecretKey	string  `mapstructure:"secret-key"  json:"secret-key"  yaml:"secret-key"`
	Bucket		string  `mapstructure:"bucket"      json:"bucket"      yaml:"bucket"`
}

type SMS struct {
	AccessKeyId	    string  `mapstructure:"access-key-id"      json:"access-key-id"      yaml:"access-key-id"`
	AccessKeySecret	string  `mapstructure:"access-key-secret"  json:"access-key-secret"  yaml:"access-key-secret"`
	SignName		string  `mapstructure:"sign-name"          json:"sign-name"          yaml:"sign-name"`
	TemplateCode	string  `mapstructure:"template-code"      json:"template-code"      yaml:"template-code"`
}

var GlobalConfig *Config

func init() {
	// 导入配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/conf.yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	// 将配置文件读到结构体中
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Println(err)
	}

}



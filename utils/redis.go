package utils

import (
	"blog-server-go/conf"
	"github.com/gomodule/redigo/redis"
)

var redisConfig = conf.GlobalConfig.Redis

// redis客户端
var redisClient redis.Conn

func init() {
	address := redisConfig.Host + ":" + redisConfig.Port
	c, err := redis.Dial("tcp", address, redis.DialPassword(redisConfig.Password))
	if err != nil {
		panic(err)
	}
	redisClient = c
}

//
// RedisGet
//  @Description: 获取键值
//  @param key
//  @return string 键值不存在则返回空串，存在则返回具体值
//  @return error
//
func RedisGet(key string) (string, error) {
	value, err := redisClient.Do("get", key)
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", nil
	}
	return redis.String(value, nil)
}

//
// RedisSet
//  @Description: 设置键值
//  @param key
//  @param value
//  @param expire 过期时间，单位s，大于0生效
//  @return error
//
func RedisSet(key string, value string, expire int) error {
	if expire <= 0 {
		_, err := redisClient.Do("set", key, value)
		return err
	} else {
		_, err := redisClient.Do("set", key, value, "EX", expire)
		return err
	}
}

//
// RedisDel
//  @Description: 删除键值
//  @param key
//  @return error
//
func RedisDel(key string) error {
	_, err := redisClient.Do("del", key)
	return err
}

//
// RedisExists
//  @Description: 判断键值是否存在
//  @param key
//  @return bool
//  @return error
//
func RedisExists(key string) (bool, error) {
	exists, err := redisClient.Do("exists", key)
	if err != nil {
		return false, err
	}
	return redis.Bool(exists, nil)
}
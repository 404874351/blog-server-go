package service

import (
	"blog-server-go/conf"
	"blog-server-go/model"
	"blog-server-go/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type JwtService struct {}

var jwtConfig = conf.GlobalConfig.Jwt

// token前缀
const TOKEN_PREFIX = "Bearer"

//
// Claims
//  @Description: 自定义jwt载荷
//
type Claims struct {
	jwt.StandardClaims
	ID          int         `json:"id"`
}

//
// CreateToken
//  @Description: 构建token
//  @receiver a
//  @param user
//  @return string
//  @return error
//
func (a *JwtService) CreateToken(user model.User) (string, error) {
	// 设置签发和过期时间
	currentTime := time.Now()
	expireTime := currentTime.Add(time.Second * time.Duration(jwtConfig.MaxAge))
	// 生成载荷
	claims := Claims{
		ID:             user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  currentTime.Unix(),
			Issuer:    "blog-server-go",
			Subject:   user.Username,
		},
	}
	// 使用指定加密方法和密钥加密并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

//
// ParseToken
//  @Description: 解析token
//  @receiver a
//  @param token
//  @return *Claims
//  @return error
//
func (a *JwtService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})
	if err != nil  {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}

//
// GetJwtByAuthorization
//  @Description: 从请求头的authorization获取token
//  @receiver a
//  @param authorization
//  @return string
//  @return error
//
func (a *JwtService) GetJwtByAuthorization(authorization string) (string, error) {
	if authorization == "" || !strings.HasPrefix(authorization, TOKEN_PREFIX){
		return "", errors.New("authorization not exists")
	}
	token := strings.ReplaceAll(authorization, TOKEN_PREFIX, "")
	token = strings.TrimSpace(token)
	return token, nil
}

//
// SetTokenInRedis
//  @Description: redis存储token
//  @receiver a
//  @param username
//  @param token
//  @return error
//
func (a *JwtService) SetTokenInRedis(username string, token string) error {
	return utils.RedisSet(username + ":token", token, jwtConfig.MaxAge)
}

//
// GetTokenInRedis
//  @Description: redis获取token
//  @receiver a
//  @param username
//  @return string
//  @return error
//
func (a *JwtService) GetTokenInRedis(username string) (string, error) {
	return utils.RedisGet(username + ":token")
}

//
// DelTokenInRedis
//  @Description: redis删除token
//  @receiver a
//  @param username
//  @return error
//
func (a *JwtService) DelTokenInRedis(username string) error {
	return utils.RedisDel(username + ":token")
}

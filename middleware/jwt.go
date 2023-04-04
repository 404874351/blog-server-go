package middleware

import (
	"blog-server-go/model/response"
	"blog-server-go/service"
	"github.com/gin-gonic/gin"
)

var jwtService = service.GlobalService.JwtService

//
// Jwt
//  @Description: jwt认证
//  @return gin.HandlerFunc
//
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, _ := jwtService.GetJwtByAuthorization(c.GetHeader("Authorization"))
		// 匿名访问，转交权限控制，白名单可自由访问，其他均拦截
		if jwt == "" {
			c.Next()
			return
		}
		// 实名访问，检查token是否非法或过期
		claims, err := jwtService.ParseToken(jwt)
		if err != nil {
			ReportError(c, response.TOKEN_ILLEGAL, err)
		}
		// 判定缓存的token是否为空或不一致
		username := claims.Subject
		token, err := jwtService.GetTokenInRedis(username)
		if err != nil {
			ReportError(c, response.TOKEN_INVALID, err)
		}
		if jwt != token {
			ReportError(c, response.TOKEN_ILLEGAL, err)
		}
		// 将载荷保存到请求域，供后续使用
		SetClaims(c, claims)

		c.Next()
	}
}

func GetClaims(c *gin.Context) *service.Claims {
	data, exists := c.Get("claims")
	if !exists {
		return nil
	}
	claims, ok := data.(*service.Claims)
	if !ok {
		return nil
	}
	return claims
}

func SetClaims(c *gin.Context, data *service.Claims) {
	c.Set("claims", data)
}

package middleware

import (
	"blog-server-go/model/response"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//
// Recovery
//  @Description: 报错恢复函数，请求报错处理，并进行格式化响应数据
//  @return gin.HandlerFunc
//
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 程序报错，panic抛出错误信息
		defer func() {
			if err := recover(); err != nil {
				// 封装提示信息
				systemError, ok := err.(response.SystemError)
				if !ok {
					systemError.Code = response.ACCESS_FAILED
				}
				// 打印并获取错误信息
				info, _ := GetData(c).(error)
				if info != nil {
					zap.L().Error(info.Error())
					//debug.PrintStack()
				} else {
					info = errors.New("error")
				}
				c.JSON(http.StatusOK, response.Fail(systemError.Code, info.Error()))
				// 终止后续中间件和函数的调用
				c.Abort()
			}
		}()
		// 执行后续工作
		c.Next()
		// 程序正常，则存在返回数据，统一数据返回格式
		// 访问无效接口，则不报错且无返回数据，报404
		data := GetData(c)
		if data != nil {
			c.JSON(http.StatusOK, response.Success(data))
		}

	}
}

//
// GetData
//  @Description: 获取上下文数据，数据由接口函数存储
//  @param c 上下文
//  @return interface{} 存储数据
//
func GetData(c *gin.Context) interface{} {
	data, exists := c.Get("data")
	if !exists {
		return nil
	}
	return data
}

//
// SetData
//  @Description: 设置数据到上下文，只在接口函数调用
//  @param c 上下文
//  @param data 存储数据
//
func SetData(c *gin.Context, data interface{}) {
	c.Set("data", data)
}

//
// ReportError
//  @Description: 上报错误，交给恢复函数处理
//  @param c
//  @param code
//  @param err
//
func ReportError(c *gin.Context, code response.StateCode, err error)  {
	SetData(c, err)
	panic(response.SystemError{Code: code})
}

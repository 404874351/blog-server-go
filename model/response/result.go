package response

import "github.com/gin-gonic/gin"

func Success(data interface{}) gin.H {
	return gin.H{
		"code": SUCCESS,
		"msg" : MsgMap[SUCCESS],
		"data": data,
	}
}

func Fail(code StateCode, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg" : MsgMap[code],
		"data": data,
	}
}

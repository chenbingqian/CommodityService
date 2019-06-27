package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const CODE_SUCCES = 10000

const CODE_SUCCES_MSG = "请求成功！"

const CODE_INVALID = 20000
const CODE_INVALID_MSG = "数据不合法"

const CODE_ERROR = 40000
const CODE_ERROR_MSG = "服务异常"

// business error
func Error(code int64, msg string, data interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{
		"result_code": code,
		"result_data": data,
		"msg":         msg,
	}
}

// system error
func ErrorSystem(err error) (int, gin.H) {
	fmt.Println(err)
	return http.StatusOK, gin.H{
		"result_code": CODE_ERROR,
		"result_data": nil,
		"msg":         CODE_ERROR_MSG,
	}
}
func Success(data interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{
		"result_code": CODE_SUCCES,
		"result_data": data,
		"msg":         CODE_SUCCES_MSG,
	}
}

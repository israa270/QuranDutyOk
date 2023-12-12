package response

import (
	"github.com/gin-gonic/gin"
)

// type response struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// 	Msg  string      `json:"msg"`
// }

// Response api
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Type string      `json:"type"`
}

const (
	Success = "success"
	Error   = "error"
	Warning = "warning"
)

const (
	ERROR   = 7
	SUCCESS = 0
	TIMEOUT = 401
)

// Result response
func Result(code int, data interface{}, msg string, statusCode int, typeResult string, c *gin.Context) {
	c.JSON(statusCode, Response{
		code,
		data,
		msg,
		typeResult,
	})
}

// Ok result
func Ok(c *gin.Context, statusCode int, typeResult string) {
	Result(SUCCESS, map[string]interface{}{}, "operations success", statusCode, typeResult, c)
}

// OkWithMessage result
func OkWithMessage(message string, statusCode int, typeResult string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, statusCode, typeResult, c)
}

// OkWithData result
func OkWithData(data interface{}, statusCode int, typeResult string, c *gin.Context) {
	Result(SUCCESS, data, "operations success", statusCode, typeResult, c)
}

// OkWithDetailed result
func OkWithDetailed(data interface{}, message string, statusCode int, typeResult string, c *gin.Context) {
	Result(SUCCESS, data, message, statusCode, typeResult, c)
}

// Fail result
func Fail(c *gin.Context, statusCode int, typeResult string) {
	Result(ERROR, map[string]interface{}{}, "operations failed", statusCode, typeResult, c)
}

// FailWithMessage result
func FailWithMessage(message string, statusCode int, typeResult string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, statusCode, typeResult, c)
}

// FailWithDetailed result
func FailWithDetailed(data interface{}, message string, statusCode int, typeResult string, c *gin.Context) {
	Result(ERROR, data, message, statusCode, typeResult, c)
}

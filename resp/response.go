package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应编号
const (
	OK              = 0
	Err             = 1001
	LoginErr        = 2001
	LoginPassErr    = 2002
	LoginCaptchaErr = 2003
	AuthErr         = 9001
	SessErr         = 9002
)

// Succ 成功响应json格式
func Succ(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": data})
}

// Fail 失败响应json格式
func Fail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": Err, "errMsg": data})
}

// LoginFail 登录失败
func LoginFail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": LoginErr, "errMsg": data})
}

// LoginPassFail 登录密码错误
func LoginPassFail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": LoginPassErr, "errMsg": data})
}

// LoginCaptchaFail 登录验证错误
func LoginCaptchaFail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": LoginCaptchaErr, "errMsg": data})
}

// AuthFail 授权失败
func AuthFail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": AuthErr, "errMsg": data})
}

// SessFail session过期
func SessFail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": SessErr, "errMsg": "Session Expired!"})
}

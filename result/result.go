package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK           = 0
	Err          = 1001
	LoginErr     = 2001
	LoginPassErr = 2002
	AuthErr      = 9001
)

//PrintJSON 响应json格式
func PrintJSON(c *gin.Context, code int, data interface{}) {
	if code != OK {
		c.JSON(http.StatusOK, gin.H{"code": code, "errMsg": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": data})
	}
}

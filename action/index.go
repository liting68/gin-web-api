package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 默认
func Index(c *gin.Context) {
	c.String(http.StatusOK, "index")
}

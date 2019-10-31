package controller

import (
	"app/model"

	"github.com/gin-gonic/gin"
)

//Home 首页控制器
type Home struct{}

//Login 登录
func (h Home) Login(c *gin.Context) {
	var user model.User
	user.Login(c)
}

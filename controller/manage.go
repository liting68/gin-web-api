package controller

import (
	"app/model"

	"github.com/gin-gonic/gin"
)

//Manage 管理员控制器
type Manage struct{}

//Login 管理员登录
func (m Manage) Login(c *gin.Context) {
	var admin model.Admin
	admin.Login(c)
}

//ManagerAccountAdd 管理员账号新增
func (m Manage) ManagerAccountAdd(c *gin.Context) {
	var admin model.Admin
	admin.Create(c)
}

//ManagerPassUpdate 管理员账号密码更新
func (m Manage) ManagerPassUpdate(c *gin.Context) {
	var admin model.Admin
	admin.UpdatePass(c)
}

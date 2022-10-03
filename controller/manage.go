package controller

import (
	"app/model"
	"app/route"

	"github.com/gin-gonic/gin"
)

// Manage 管理员控制器
type Manage struct{}

func (m Manage) Auth() string {
	return model.BEARER.Admin
}

// LoginAdmin 登录者Admin
func (m Manage) LoginAdmin(c *gin.Context) model.Admin {
	headAuth := c.Request.Header.Get("Authorization")
	var admin model.Admin
	if len(headAuth) != 0 {
		claim, err := route.ParseToken(headAuth)
		if err == nil {
			if len(claim.Audience) > 0 {
				admin = admin.FirstByUsername(claim.Audience[0])
			}
		}
	}
	return admin
}

// Login 管理员登录
func (m Manage) Login(c *gin.Context) {
	res, err := model.Admin{}.Login(c)
	if err != nil {
		route.Fail(c, err.Error())
		return
	}
	route.Succ(c, res)
}

// ManagerAccountAdd 管理员账号新增
func (m Manage) ManagerAccountAdd(c *gin.Context) {
	res, err := model.Admin{}.Create(c)
	if err != nil {
		route.Fail(c, err.Error())
		return
	}
	route.Succ(c, res)
}

// ManagerPassUpdate 管理员账号密码更新
func (m Manage) ManagerPassUpdate(c *gin.Context) {
	loginAdmin := m.LoginAdmin(c)
	if loginAdmin.ID == 0 {
		route.SessFail(c)
		return
	}
	res, err := model.Admin{}.UpdatePass(c, loginAdmin)
	if err != nil {
		route.Fail(c, err.Error())
		return
	}
	route.Succ(c, res)

}

// ModifyPass 修改密码
func (m Manage) ModifyPass(c *gin.Context) {
	loginAdmin := m.LoginAdmin(c)
	if loginAdmin.ID == 0 {
		route.SessFail(c)
		return
	}
	var admin model.Admin
	res, err := admin.ModifyPass(c, loginAdmin)
	if err != nil {
		route.Fail(c, err.Error())
		return
	}
	route.Succ(c, res)

}

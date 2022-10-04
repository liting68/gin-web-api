package controller

import (
	"app/model"
	"app/route"

	"github.com/gin-gonic/gin"
)

// User 用户控制器
type User struct{}

// LoginAdmin 登录者Admin
func (u User) LoginUser(c *gin.Context) model.User {
	headAuth := c.Request.Header.Get("Authorization")
	var user model.User
	if len(headAuth) != 0 {
		claim, err := route.ParseToken(headAuth)
		if err == nil {
			if len(claim.Audience) > 0 {
				user = user.FirstByUsername(claim.Audience[0])
			}
		}
	}
	return user
}

// Ping 连通测试
func (u User) Ping(c *gin.Context) {
	route.Succ(c, "pong")
}

// Login 登录
func (u User) Login(c *gin.Context) {
	res, err := model.User{}.Login(c)
	if err != nil {
		if err.Error() == "密码错误" {
			route.LoginPassFail(c, err.Error())
			return
		}
		route.LoginFail(c, err.Error())
		return
	}
	route.Succ(c, res)
}

func (u User) Info(c *gin.Context) {
	login := u.LoginUser(c)
	if login.ID == 0 {
		route.SessFail(c)
		return
	}
	res, err := model.User{}.Info(c, login)
	if err != nil {
		route.Fail(c, err.Error())
		return
	}
	route.Succ(c, res)
}

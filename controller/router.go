package controller

import (
	"app/model"
	"app/route"
	"os"

	"github.com/gin-gonic/gin"
)

// RegisterServer 注册路由服务
func RegisterServer() *gin.Engine {
	g := route.InitRouter()

	crontab := Crontab{}
	crontab.Start()

	basedir, _ := os.Getwd()
	g.Static("/web", basedir+"/web") //静态资源

	user := User{}
	g.POST("/login", user.Login)
	g.POST("/ping", user.Ping)

	manage := Manage{}
	g.POST("/manage/login", manage.Login)
	g.POST("/manage/account", manage.ManagerAccountAdd)
	g.POST("/manage/password-update", manage.ManagerPassUpdate)
	g.POST("/manage/password", route.Auth(model.BEARER.Admin), manage.ModifyPass)

	return g
}

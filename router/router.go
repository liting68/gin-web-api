package router

import (
	"app/action"
	"app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

//SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	// router := gin.Default()
	router := gin.New()
	// 添加自定义的 logger 中间件
	// router.Use(gin.Logger(), gin.Recovery())
	RegisterRouter(router)
	return router
}

//RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	basedir, _ := os.Getwd()
	router.Static("/web", basedir+"/web")

	home := action.HomeAction{}
	router.POST("/login", home.Login)
	router.GET("/init", middleware.Auth(), home.InitInfo)

}

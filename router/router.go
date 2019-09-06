package router

import (
	"app/action"
	"app/middleware"

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
	router.GET("/", action.Index)
	router.POST("/login", action.Login)
	router.POST("/updateManagerPass", middleware.Auth(), action.UpdateManagerPass)
	router.POST("/addManagerAccount", action.AddManagerAccount)
}

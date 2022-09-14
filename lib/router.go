package lib

import (
	"app/config"
	"app/controller"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

// InitGin 初始化路由
func InitGin() *gin.Engine {
	if config.Info.App.Debug != true {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = colorable.NewColorableStdout()
	g := gin.New()
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout
	g.Use(gin.Logger())
	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	g.Use(gin.Recovery())
	//Cors 跨域设置
	g.Use(Cors())

	//Timeout超时控制
	// g.Use(timeout.Timeout(
	// 	timeout.WithTimeout(60*time.Second),
	// 	timeout.WithErrorHttpCode(http.StatusOK),
	// 	timeout.WithDefaultMsg(`{"code": 1001, "errMsg":"Request timeout"}`),
	// 	timeout.WithCallBack(func(r *http.Request) {
	// 		fmt.Println("timeout happen, url:", r.URL.String())
	// 	}),
	// ))

	store := sessions.NewCookieStore([]byte("secret"))
	g.Use(sessions.Sessions("MyProject-session", store))
	RegisterRouter(g)
	return g
}

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	basedir, _ := os.Getwd()
	router.Static("/web", basedir+"/web")

	home := controller.Home{}
	router.POST("/login", home.Login)

	manage := controller.Manage{}
	router.POST("/manage/login", manage.Login)
	router.POST("/manage/account", manage.ManagerAccountAdd)
	router.POST("/manage/password-update", manage.ManagerPassUpdate)
	router.POST("/manage/password", ManageAuth(), manage.ModifyPass)

}

// Cors 跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,X-Requested-With,accept,client-security-token")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
		}
		c.Next()
	}
}

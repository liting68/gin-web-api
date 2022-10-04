package route

import (
	"app/config"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	if !config.Info.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = colorable.NewColorableStdout()
	g := gin.New()
	g.SetTrustedProxies(nil)
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
	g.Use(sessions.Sessions("MyWeb-session", store))
	return g
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

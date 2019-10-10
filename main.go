package main

import (
	"app/db"
	"app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := router.SetupRouter()
	// router.Use(Cors())

	router.Run("127.0.0.1:8081")
	defer db.CloseDB()
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,X-Requested-With,accept,client-security-token")
// 		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")
// 		c.Header("Access-Control-Allow-Credentials", "true")
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "")
// 		}
// 		c.Next()
// 	}
// }

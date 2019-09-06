package main

import (
	"app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := router.SetupRouter()
	router.Run("localhost:80")
}
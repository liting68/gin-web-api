package main

import (
	"app/controller"
	"app/db"
)

func main() {
	router := controller.RegisterRouter()
	router.Run("127.0.0.1:8068")
	defer db.CloseDB()
}

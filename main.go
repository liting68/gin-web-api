package main

import (
	"app/controller"
	"app/db"
)

func main() {
	sev := controller.RegisterServer()
	sev.Run("127.0.0.1:8068")
	defer db.CloseDB()
}

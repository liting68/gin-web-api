package main

import (
	"app/db"
	"app/lib"
)

func main() {
	lib.StartCron()
	router := lib.InitGin()
	router.Run("127.0.0.1:8081")
	defer db.CloseDB()
}

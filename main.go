package main

import (
	"app/db"
	"app/lib"
	"log"
	"syscall"

	"github.com/fvbock/endless"
)

func main() {
	lib.StartCron() //定时任务
	router := lib.InitGin()
	server := endless.NewServer("127.0.0.1:8081", router) // router.Run("127.0.0.1:8081")
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	server.ListenAndServe()
	defer db.CloseDB()
}

package main

import (
	"app/db"
	"app/lib"
)

func main() {
	lib.StartCron() //定时任务
	router := lib.InitGin()
	router.Run("127.0.0.1:8068")
	// s := &http.Server{
	// 	Addr:              ":8091",
	// 	Handler:           router,
	// 	ReadHeaderTimeout: 0,
	// 	ReadTimeout:       60 * time.Second,
	// }
	defer db.CloseDB()
}

package lib

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

//StartCron 初始化定时任务 秒 分 时 天 月 周
func StartCron() {
	sha, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithSeconds(), cron.WithLocation(sha))
	c.AddFunc("1 * * * * *", func() {
		fmt.Println("每分钟定时任务:", time.Now().Format("2006-01-02 15:04:05"))
	})

	c.Start()
}

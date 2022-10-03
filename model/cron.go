package model

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type Crontab struct{}

var CRONTAB *cron.Cron

func init() {
	sha, _ := time.LoadLocation("Asia/Shanghai")
	CRONTAB = cron.New(cron.WithSeconds(), cron.WithLocation(sha))
}

// StartCron 初始化定时任务 秒 分 时 天 月 周
func (ct Crontab) StartCron() {
	CRONTAB.AddFunc("1 * * * * *", func() {
		fmt.Println("每分钟定时任务:", time.Now().Format("2006-01-02 15:04:05"))
	})
	CRONTAB.Start()
}

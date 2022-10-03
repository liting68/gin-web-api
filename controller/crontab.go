package controller

import "app/model"

// Cron 定时任务管理控制器
type Crontab struct{}

// Start 任务常驻并开始
func (ct Crontab) Start() {
	model.Crontab{}.StartCron()
}

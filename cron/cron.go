package cron

import (
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/robfig/cron"
)

func CronInit() *cron.Cron {
	c := cron.New()

	// c.AddFunc("0 0 0 * *", func() {
	// 	models.UpdateTask()
	// })
	c.AddFunc("0 0 1 * *", func() {
		db := setting.MysqlConn()
		db.Table("user").Select("selection_num").Update(5)
	})

	return c
}

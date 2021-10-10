package cron

import (
	"github.com/robfig/cron"
)

func CronInit() *cron.Cron {
	c := cron.New()

	// c.AddFunc("0 0 0 * *", func() {
	// 	models.UpdateTask()
	// })

	return c
}

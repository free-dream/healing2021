package cron

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/robfig/cron"
	"strconv"
)

func CronInit() *cron.Cron {
	c := cron.New()

	// c.AddFunc("0 0 0 * *", func() {
	// 	models.UpdateTask()
	// })
	c.AddFunc("0 0 0 * *", func() {
		db := setting.MysqlConn()
		db.Table("user").Select("selection_num").Update(2)
	})

	c.AddFunc("0 */1 * * *", func() {
		redisCli := setting.RedisConn()
		db := setting.MysqlConn()
		rows, _ := db.Exec("select cover_id,count(user_id) as lieks from praise where is_liked=1 group by cover_id;").Rows()
		like := dao.LikeObj{}
		for rows.Next() {
			err := db.ScanRows(rows, &like)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			redisCli.HSet("healing2021:praise of cover", strconv.Itoa(like.CoverId), like.Likes)
		}
		rows.Close()
	})

	return c
}

package sandwich

//需要更新的主要是以积分和点赞为代表的高频表
//这些表更新的同时还需要更新相关的排序

import (
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type MailBox struct {
	LikesBox  map[string]int
	PointsBox map[string]int
	TaskBox   map[string]int
}

var Sandwich *MailBox

func init() {
	Sandwich = new(MailBox)
	Sandwich.LikesBox = make(map[string]int)
	Sandwich.PointsBox = make(map[string]int)
	Sandwich.TaskBox = make(map[string]int)
}

func UpdatePoints() {
	mysqlDb := setting.MysqlConn()
	redisDb := setting.RedisConn()

}

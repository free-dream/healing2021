package sandwich

//需要更新的主要是以积分和点赞为代表的高频表
//这些表更新的同时还需要更新相关的排序

import (
	"strconv"
	"strings"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type MailBox struct {
	PointsBox map[string]int
	TaskBox   map[string]int
}

var Sandwich *MailBox

func init() {
	Sandwich = new(MailBox)
	Sandwich.PointsBox = make(map[string]int)
	Sandwich.TaskBox = make(map[string]int)
}

//更新points数据
func (box *MailBox) updatePoints() {
	redisDb := setting.RedisConn()

	//判空
	if len(box.PointsBox) == 0 {
		return
	}

	//取值并写回
	var record, point int
	for key, _ := range box.PointsBox {
		value := redisDb.HMGet(key, "record", "point").Val()
		record = value[0].(int)
		point = value[1].(int)

		//解析key中的userid
		userid, _ := strconv.Atoi(strings.Split(key, "/")[0])
		dao.UpdatePoints(userid, record, point)

		delete(box.PointsBox, key)
	}
}

//更新点赞数
func (box *MailBox) updateLikes() {
	redisDb := setting.RedisConn()
	//userid
	userid := 0
	//

	coverdata := redisDb.ZRangeWithScores("cover", 0, -1).Val()
	var likes, member int
	for _, unit := range coverdata {
		likes = int(unit.Score)
		member = unit.Member.(int)
		dao.UpdateLikesByID(userid, member, likes, "cover")
	}

	momentdata := redisDb.ZRangeWithScores("moment", 0, -1).Val()
	for _, unit := range momentdata {
		likes = int(unit.Score)
		member = unit.Member.(int)
		dao.UpdateLikesByID(userid, member, likes, "cover")
	}

	commentdata := redisDb.ZRangeWithScores("comment", 0, -1).Val()
	for _, unit := range commentdata {
		likes = int(unit.Score)
		member = unit.Member.(int)
		dao.UpdateLikesByID(userid, member, likes, "cover")
	}
}

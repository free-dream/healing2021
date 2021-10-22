package sandwich

//需要更新的主要是以积分和点赞为代表的高频表
//这些表更新的同时还需要更新相关的排序
//redis的并发控制比较头疼

import (
	"strconv"
	"strings"
	"sync"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type MailBox struct {
	PointBox map[string]int
	TaskBox  map[string]int
	LikeBox  bool
}

var (
	Sandwich *MailBox
)

func init() {
	Sandwich = new(MailBox)
	Sandwich.PointBox = make(map[string]int)
	Sandwich.TaskBox = make(map[string]int)
}

func (box *MailBox) Update(tube chan int) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		if len(box.PointBox) > 0 {
			box.UpdatePoints()
		}
		wg.Done()
	}()

	go func() {
		if len(box.TaskBox) > 0 {
			box.UpdateTasks()
		}
		wg.Done()
	}()

	go func() {
		if box.LikeBox {
			box.UpdateLikes()
		}
		wg.Done()
	}()

	wg.Wait()
}

//更新points数据
func (box *MailBox) UpdatePoints() {
	redisDb := setting.RedisConn()

	//判空
	if len(box.PointBox) == 0 {
		return
	}

	//取值并写回
	var record, point int
	for key, _ := range box.PointBox {
		value := redisDb.HMGet(key, "record", "point").Val()
		record = value[0].(int)
		point = value[1].(int)

		//解析key中的userid
		userid, _ := strconv.Atoi(strings.Split(key, "/")[0])
		dao.UpdatePoints(userid, record, point)

		delete(box.PointBox, key)
	}
}

//更新点赞数
func (box *MailBox) UpdateLikes() {
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
		dao.UpdateLikesByID(userid, member, likes, "moment")
	}

	commentdata := redisDb.ZRangeWithScores("comment", 0, -1).Val()
	for _, unit := range commentdata {
		likes = int(unit.Score)
		member = unit.Member.(int)
		dao.UpdateLikesByID(userid, member, likes, "comment")
	}
}

//更新任务信息
func (box *MailBox) UpdateTasks() {
	redisDb := setting.RedisConn()

	//判空
	if len(box.PointBox) == 0 {
		return
	}

	//取值并写回
	var process, check int
	for key, _ := range box.TaskBox {
		value := redisDb.HMGet(key, "process", "check").Val()
		process = value[0].(int)
		check = value[1].(int)

		//解析key中的userid
		taskid, _ := strconv.Atoi(strings.Split(key, "/")[2])
		userid, _ := strconv.Atoi(strings.Split(key, "/")[0])
		dao.UpdateTasks(userid, taskid, process, check)

		delete(box.PointBox, key)
	}
}

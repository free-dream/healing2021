package sandwich

import (
	// "git.100steps.top/100steps/healing2021_be/models/statements"
	"math"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//liketype{0:cover;1:moment;comment:2}
//throw回bool,交付接口处理点赞信息

//其它的任务任务对应积分常量
const (
	DRAWCOST int = -200
)

func InitKVs() {
	redisDb := setting.RedisConn()
	redisDb.SAdd("cover", "init")
	redisDb.SAdd("moment", "init")
	redisDb.SAdd("comment", "init")
}

//点赞，可直接调用
func Likes(targetid int, liketype int) bool {
	//此处应有读取用户userid的操作
	var userid int = 0
	//
	redisDb := setting.RedisConn()
	var targettype string
	switch liketype {
	case 0:
		targettype = "cover"
	case 1:
		targettype = "moment"
	case 2:
		targettype = "comment"
	default:
		targettype = ""
		panic("likes error")
	}
	tempkey := "user" + strconv.Itoa(userid) + targettype
	check := redisDb.SIsMember(tempkey, targetid).Val()
	if !check {
		redisDb.ZIncrBy(targettype, 1, strconv.Itoa(targetid))
	}
	return !check
}

//修改用户点数
func Changepoints(point float64) bool {
	//此处应有读取用户userid的操作
	var userid int = 0
	//
	redisDb := setting.RedisConn()
	tempkey := strconv.Itoa(userid) + "point"

	if data, _ := strconv.Atoi(redisDb.HGet(tempkey, "point").Val()); data < int(math.Abs(point)) {
		return false
	} else if point > 0 {
		redisDb.HIncrBy(tempkey, "record", int64(point))
	}
	redisDb.HIncrBy(tempkey, "point", int64(point))
	return true
}

//用户任务进度更新,若任务已完成，return true
func HandleTask(process int, taskid int) bool {
	//此处应有读取用户userid的操作
	var userid int = 0
	//
	redisDb := setting.RedisConn()
	tempkey := strconv.Itoa(userid) + "task" + strconv.Itoa(taskid)

	currentval := redisDb.HIncrBy(tempkey, "record", int64(process)).Val()
	target := redisDb.HMGet("task"+strconv.Itoa(taskid), "target").Val()
	targetval, _ := target[0].(int64)

	if currentval >= targetval {
		redisDb.HIncrBy(tempkey, "check", 1)
		return true
	}
	return false
}

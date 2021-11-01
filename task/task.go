package task

import (
	"strconv"

	state "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//本次任务目前都是一次性的，没有计数要求 2021.11.1

//任务注册和初始化
var (
	ST SelectionTask
	MT MomentTask
	HT HealingTask
)

func init() {
	ST = SelectionTask{
		TID: STID,
	}
	MT = MomentTask{
		TID: MTID,
	}
	HT = HealingTask{
		TID: HTID,
	}
}

//理论上任务有需求都可以在此处扩展,需要实现接口里的主要方法
//一次性任务
type MetaOTask interface {
	AddRecord(int) bool
	CheckMax(int) bool
}

//计数型任务,需要记录任务进度和检查是否完成
type MetaCTask interface {
	AddRecord(int) bool
	Counter(int) bool
	Check(int) bool
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

//设置redis任务缓存,此处还未设置expile time
func CacheTask(userid int, tid int, value interface{}) error {
	redisDb := setting.RedisConn()
	temp := make(map[string]interface{})
	key := strconv.Itoa(userid) + "/task"
	temp[strconv.Itoa(tid)] = value
	err := redisDb.HMSet(key, temp).Err()
	return err
}

//更新任务记录
func UpdateTask(userid int, tid int, value int64) error {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	err := redisDb.HIncrBy(key, strconv.Itoa(tid), value).Err()
	return err
}

//取用用户积分缓存
func GetCachePoints(userid int) int {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	temp := redisDb.HMGet(key, "points").Val()
	if len(temp) < 1 {
		return -1
	}
	return temp[0].(int)
}

//基于mysql更新用户积分缓存
func UpdateCachePoints(userid int, points int) error {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	temp := make(map[string]interface{})
	temp["points"] = points
	err := redisDb.HMSet(key, temp).Err()
	return err
}

//取用redis任务缓存
func GetCacheTask(userid int, tid int) int {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	temp := redisDb.HMGet(key, strconv.Itoa(tid)).Val()
	if len(temp) < 1 {
		return -1
	}
	data := temp[0].(int)
	return data
}

//修改用户点数,任务特供
//同时更新总点数和用户点数
func ChangePoints(point float32, userid int, tid int) bool {
	redisDb := setting.RedisConn()
	mysqlDb := setting.MysqlConn()

	tempkey := strconv.Itoa(userid) + "/point"
	temp := redisDb.HIncrBy(tempkey, "points", int64(point)).Val()
	tempf := redisDb.HIncrBy(tempkey, strconv.Itoa(tid), int64(point)).Val()

	//单独拉出一个协程更新数据库以保证数据一致性
	ch := make(chan int)
	go func() {
		var user state.User
		var tasktable state.TaskTable

		err := mysqlDb.Where("ID = ?", userid).First(&user).Error
		errHandler(err)

		err = mysqlDb.Where("TaskID = ? AND UserID = ?", tid, userid).Find(&tasktable).Error
		errHandler(err)

		user.Points = int(temp)
		err = mysqlDb.Save(&user).Error

		errHandler(err)
		tasktable.Counter = int(tempf)

		err = mysqlDb.Save(&tasktable).Error
		errHandler(err)
		<-ch
	}()
	ch <- 0
	return true
}

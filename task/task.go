package task

import (
	"strconv"

	state "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//具体任务参数
//本次任务目前都是一次性的，没有计数要求

//常量：积分和积分上限要求

//新增定时器
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

//设置redis任务缓存,此处还未设置expile time
func CacheTask(userid int, field string, value interface{}) error {
	redisDb := setting.RedisConn()
	temp := make(map[string]interface{})
	key := strconv.Itoa(userid) + "/task"
	temp[field] = value
	err := redisDb.HMSet(key, temp).Err()
	return err
}

//更新任务记录
func UpdateTask(userid int, field string, value int64) error {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	err := redisDb.HIncrBy(key, field, value).Err()
	return err
}

//取用redis任务缓存
func GetCacheTask(userid int, field string) int {
	redisDb := setting.RedisConn()
	key := strconv.Itoa(userid) + "/task"
	temp := redisDb.HMGet(key, field).Val()
	if len(temp) < 1 {
		return -1
	}
	data := temp[0].(int)
	return data
}

//修改用户点数,任务特供
//同时更新总点数和用户点数
func ChangePoints(point float32, userid int, field string) bool {
	redisDb := setting.RedisConn()
	mysqlDb := setting.MysqlConn()

	tempkey := strconv.Itoa(userid) + "/point"
	temp := redisDb.HIncrBy(tempkey, "points", int64(point)).Val()
	// tempf := redisDb.HIncrBy(tempkey, field, int64(point)).Val()

	//单独拉出一个协程更新数据库以保证数据一致性
	ch := make(chan int)
	go func() {
		var user state.User
		err := mysqlDb.Where("ID = ?", userid).First(&user).Error
		if err != nil {
			panic(err)
		}
		user.Points = int(temp)
		err = mysqlDb.Save(&user).Error
		if err != nil {
			panic(err)
		}
		<-ch
	}()
	ch <- 0
	return true
}

package task

import (
	"strconv"

	state "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//本次任务目前都是一次性的，没有计数要求 2021.11.1
//本次任务目前为止所有用户的任务都是一样的 2021.11.2
//不存在线性任务、个性化任务和条件性任务

//任务注册和初始化，声明变量和init()
var (
	ST SelectionTask
	MT MomentTask
	HT HealingTask
)

func init() {
	//点歌任务
	ST = SelectionTask{
		TID: STID,
	}
	//动态任务
	MT = MomentTask{
		TID: MTID,
	}
	//唱歌任务
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

//错误处理，暂定
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
	key := strconv.Itoa(userid) + "/points"
	temp := redisDb.HMGet(key, "points").Val()
	if len(temp) < 1 {
		// //
		// fmt.Println("redis读取points为空")
		// //
		return -1
	}
	data, ok := temp[0].(string)
	if !ok {
		// //
		// fmt.Println("类型断言有误")
		// //
		return -1
	}
	temp1, check := strconv.Atoi(data)
	if check != nil {
		// //
		// fmt.Println("读取的非数字")
		// //
		return -1
	}
	return temp1
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
	// if len(temp) < 1 {
	// 	return -1
	// }
	if temp[0] == nil {
		return -1
	}
	data := temp[0].(string)
	temp1, check := strconv.Atoi(data)
	if check != nil {
		return -1
	}
	return temp1
}

//同时更新总点数和任务点数
func ChangePoints(point float32, userid int, tid int) bool {
	redisDb := setting.RedisConn()
	mysqlDb := setting.MysqlConn()

	tempkey := strconv.Itoa(userid) + "/point"
	temp := redisDb.HIncrBy(tempkey, "points", int64(point)).Val()
	tempf := redisDb.HIncrBy(tempkey, strconv.Itoa(tid), int64(point)).Val()
	// //redis缓存读取测试
	// fmt.Println(redisDb.HMGet(tempkey, "points").Val())
	// fmt.Println(redisDb.HMGet(tempkey, strconv.Itoa(tid)).Val())
	//单独拉出一个协程更新数据库以保证数据一致性
	//错误处理
	ch := make(chan int)
	go func() {
		var user state.User
		var tasktable state.TaskTable

		err := mysqlDb.Where("id = ?", userid).First(&user).Error
		errHandler(err)

		err = mysqlDb.Where("task_id = ? AND user_id = ?", tid, userid).Find(&tasktable).Error
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

// 已实现于dao/task.go
// //在用户首次登录时创建对应的任务表
// func CreateTaskTable(userid int, taskid int) error {
// 	mysqlDb := setting.MysqlConn()
// 	usertask := state.TaskTable{
// 		UserId: userid,
// 		TaskId: taskid,
// 	}
// 	err := mysqlDb.Create(&usertask).Error
// 	return err
// }

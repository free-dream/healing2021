package sandwich

import (
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

const (
	prefix = "healing2021:"
)

//设置redis任务缓存
func CacheTask(userid int, tid int, value interface{}) error {
	redisDb := setting.RedisConn()
	temp := make(map[string]interface{})
	key := prefix + strconv.Itoa(userid) + "/task"
	temp[strconv.Itoa(tid)] = value
	err := redisDb.HMSet(key, temp).Err()
	return err
}

//更新任务记录,保留一个返回更新后的积分便于扩展
func UpdateTask(userid int, tid int, value int64) (int, error) {
	redisDb := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "/task"
	temp := redisDb.HIncrBy(key, strconv.Itoa(tid), value)
	if err := temp.Err(); err != nil {
		return -1, err
	}
	val := temp.Val()
	return int(val), nil
}

//取用用户单项任务积分缓存
func GetCacheTaskPoints(userid int, tid int) int {
	redisDb := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "/task"
	temp := redisDb.HMGet(key, strconv.Itoa(tid)).Val()
	if len(temp) < 1 {
		return 0
	}
	data, ok := temp[0].(string)
	if !ok {
		return 0
	}
	temp1, check := strconv.Atoi(data)
	if check != nil {
		return -1
	}
	return temp1
}

// //基于mysql更新用户积分缓存
// func UpdateCachePoints(userid int, points int) error {
// 	redisDb := setting.RedisConn()
// 	key := prefix + strconv.Itoa(userid) + "/task"
// 	temp := make(map[string]interface{})
// 	temp["points"] = points
// 	err := redisDb.HMSet(key, temp).Err()
// 	return err
// }

//以下同质化函数太多，可以考虑综合一下，暂时先不做了
//可惜go不支持重载

//缓存当前用户排名,设置了1h的expile时间
func CacheCURanking(userid int, rank string) error {
	db := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "rank"
	err := db.Set(key, rank, time.Hour).Err()
	return err
}

//获取当前用户排名
func GetCURanking(userid int) string {
	db := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "rank"
	data := db.Get(key).Val()
	return data
}

//缓存积分排名，生存周期1h
func CachePointsRanking(school string, data string) error {
	db := setting.RedisConn()
	key := prefix + "ranking"
	err := db.HSet(key, school, data).Err()
	if err != nil {
		return err
	}
	//设置过期时间
	db.Expire(key, time.Hour)
	return nil
}

//取用积分排名
func GetPointsRanking(school string) string {
	db := setting.RedisConn()
	key := prefix + "ranking"
	temp := db.HGet(key, school).Val()
	return temp
}

//缓存每日排名,生命周期5min
func CacheDailyRank(date string, data string) error {
	db := setting.RedisConn()
	key := prefix + "dailyrank"
	err := db.HSet(key, date, data).Err()
	if err != nil {
		return err
	}
	//设置过期时间
	db.Expire(key, time.Minute*5)
	return nil
}

//提取对应的每日排名
func GetDailyRankByDate(date string) string {
	db := setting.RedisConn()
	key := prefix + "dailyrank"
	data := db.HGet(key, date).Val()
	return data
}

package sandwich

import (
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

const (
	prefix = "2021healing:"
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

//更新任务记录
func UpdateTask(userid int, tid int, value int64) error {
	redisDb := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "/task"
	err := redisDb.HIncrBy(key, strconv.Itoa(tid), value).Err()
	return err
}

//取用用户积分缓存
func GetCachePoints(userid int) int {
	redisDb := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "/points"
	temp := redisDb.HMGet(key, "points").Val()
	if len(temp) < 1 {
		return -1
	}
	data, ok := temp[0].(string)
	if !ok {
		return -1
	}
	temp1, check := strconv.Atoi(data)
	if check != nil {
		return -1
	}
	return temp1
}

//基于mysql更新用户积分缓存
func UpdateCachePoints(userid int, points int) error {
	redisDb := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "/task"
	temp := make(map[string]interface{})
	temp["points"] = points
	err := redisDb.HMSet(key, temp).Err()
	return err
}

//以下同质化函数太多，可以考虑综合一下，暂时先不做了
//可惜go不支持重载

//缓存当前用户排名,设置了10min的expile时间
func CacheCURanking(userid int, rank string) error {
	db := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "rank"
	err := db.Set(key, rank, time.Minute*10).Err()
	return err
}

//获取当前用户排名
func GetCURanking(userid int) string {
	db := setting.RedisConn()
	key := prefix + strconv.Itoa(userid) + "rank"
	data := db.Get(key).Val()
	return data
}

//缓存积分排名，每小时更新
func CachePointsRanking(school string, data string) error {
	db := setting.RedisConn()
	key := prefix + "ranking"
	err := db.HSet(key, school, data).Err()
	return err
}

//取用积分排名
func GetPointsRanking(school string) string {
	db := setting.RedisConn()
	key := prefix + "ranking"
	temp := db.HGet(key, school).Val()
	return temp
}

//缓存每日排名
func CacheDailyRank(date string, data string) error {
	db := setting.RedisConn()
	key := prefix + "dailyrank"
	err := db.HSet(key, date, data).Err()
	return err
}

//提取对应的每日排名
func GetDailyRankByDate(date string) string {
	db := setting.RedisConn()
	key := prefix + "dailyrank"
	data := db.HGet(key, date).Val()
	return data
}

//使用redis作为换入区，保留20项记录
func CacheSearch() {}

//取用缓存数据
func GetSearch() {}

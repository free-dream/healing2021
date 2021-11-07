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

//首次调用缓存积分排名
//不设置expile,而是调用cron整点删除一次redis
func CachePointsRanking(school string, data string) error {
	db := setting.RedisConn()
	key := prefix + "ranking"
	err := db.HSet(key, school, data).Err()
	return err
}

//取用缓存
func GetPoingsRanking(school string) string {
	db := setting.RedisConn()
	key := prefix + "ranking"
	temp := db.HGet(key, school).Val()
	return temp
}

//使用redis作为换入区，保留20项记录
func CacheSearch() {}

//取用缓存数据
func GetSearch() {}

// //下述接口暂时废案

// /*翻唱排序系统初始化的时候缓存一次，之后定期更新*/
// //用户信息在登录时缓存一次

// //缓存翻唱排序
// //request="流派or语言"+"排序要求"
// //例："日语综合"="japanese/composite"
// //例："推荐最新"="recommend/latest"
// //索引时按上述说法进行get，例如 HMGet("japanese/composite/cover1/file"),根据设计i=1-15
// func CacheCovers(models []*statements.Cover, request string) {
// 	redisDB := setting.RedisConn()
// 	temp := make(map[string]interface{})
// 	for i, data := range models {
// 		st := reflect.TypeOf(data)
// 		sv := reflect.ValueOf(data)
// 		for i := 0; i < st.NumField(); i++ {
// 			field := st.Field(i)
// 			if tag, ok := field.Tag.Lookup("json"); ok {
// 				if tag == "" {
// 					continue
// 				} else {
// 					temp[tag] = sv.Field(i).Interface()
// 				}
// 			} else {
// 				continue
// 			}
// 		}
// 		key := request + "/cover" + strconv.Itoa(i)
// 		redisDB.HMSet(key, temp)
// 	}
// }

// //缓存点歌排序
// //request="流派"+"排序要求"
// //例："日语综合"="japanese/composite"
// //例："推荐最新"="recommend/latest"
// //索引时按上述说法进行get，例如 HMGet("japanese/composite/selction1","avatar"),根据设计i=1-15
// func CacheSelections(records []*statements.Selection, request string) {
// 	redisDB := setting.RedisConn()
// 	temp := make(map[string]interface{})
// 	for i, data := range records {
// 		st := reflect.TypeOf(data)
// 		sv := reflect.ValueOf(data)
// 		for i := 0; i < st.NumField(); i++ {
// 			field := st.Field(i)
// 			if tag, ok := field.Tag.Lookup("json"); ok {
// 				if tag == "" {
// 					continue
// 				} else {
// 					temp[tag] = sv.Field(i).Interface()
// 				}
// 			} else {
// 				continue
// 			}
// 		}
// 		key := request + "/cover" + strconv.Itoa(i)
// 		redisDB.HMSet(key, temp)
// 	}
// }

// //缓存用户数据，尤其是点歌用的积分
// func CacheUser(record *statements.User) {
// 	redisDB := setting.RedisConn()
// 	temp := make(map[string]interface{})
// 	st := reflect.TypeOf(record)
// 	sv := reflect.ValueOf(record)
// 	for i := 0; i < st.NumField(); i++ {
// 		field := st.Field(i)
// 		if tag, ok := field.Tag.Lookup("json"); ok {
// 			tempkey := strconv.Itoa(int(record.ID)) + "points"
// 			if tag == "" {
// 				continue
// 			} else if tag == "points" {
// 				redisDB.HSet(tempkey, "points", sv.Field(i).Int())
// 			} else if tag == "record" {
// 				redisDB.HSet(tempkey, "record", sv.Field(i).Int())
// 			} else {
// 				temp[tag] = sv.Field(i).Interface()
// 			}
// 		} else {
// 			continue
// 		}
// 		key := "user" + strconv.Itoa(int(record.ID))
// 		redisDB.HMSet(key, temp)
// 	}
// }

// //系统启动时直接初始化，加载任务文本时直接从redis读取
// func CacheTask(records []*statements.Task) {
// 	redisDB := setting.RedisConn()
// 	temp := make(map[string]interface{})
// 	for i, data := range records {
// 		st := reflect.TypeOf(data)
// 		sv := reflect.ValueOf(data)
// 		for i := 0; i < st.NumField(); i++ {
// 			field := st.Field(i)
// 			if tag, ok := field.Tag.Lookup("json"); ok {
// 				if tag == "" {
// 					continue
// 				} else {
// 					temp[tag] = sv.Field(i).Interface()
// 				}
// 			} else {
// 				continue
// 			}
// 		}
// 		key := "task" + strconv.Itoa(i)
// 		redisDB.HMSet(key, temp)
// 	}
// }

// //缓存所有礼品
// func CachePrizes(lotteries []*statements.Lottery) {
// 	redisDB := setting.RedisConn()
// 	temp := make(map[string]interface{})
// 	for i, data := range lotteries {
// 		st := reflect.TypeOf(data)
// 		sv := reflect.ValueOf(data)
// 		for i := 0; i < st.NumField(); i++ {
// 			field := st.Field(i)
// 			if tag, ok := field.Tag.Lookup("json"); ok {
// 				if tag == "" {
// 					continue
// 				} else {
// 					temp[tag] = sv.Field(i).Interface()
// 				}
// 			} else {
// 				continue
// 			}
// 		}
// 		key := "lottery" + strconv.Itoa(i)
// 		redisDB.HMSet(key, temp)
// 	}
// }

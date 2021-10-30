package sandwich

import (
	"reflect"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//缓存当前用户排名,设置了8小时的有效期
func CacheCURanking(userid int, rank string) error {
	db := setting.RedisConn()
	key := strconv.Itoa(userid) + "rank"
	err := db.Set(key, rank, time.Hour*8).Err()
	return err
}

//获取当前用户排名
func GetCURanking(userid int) string {
	db := setting.RedisConn()
	key := strconv.Itoa(userid) + "rank"
	data := db.Get(key).Val()
	return data
}

//下述接口暂时废案

/*翻唱排序系统初始化的时候缓存一次，之后定期更新*/
//用户信息在登录时缓存一次

//缓存翻唱排序
//request="流派or语言"+"排序要求"
//例："日语综合"="japanese/composite"
//例："推荐最新"="recommend/latest"
//索引时按上述说法进行get，例如 HMGet("japanese/composite/cover1/file"),根据设计i=1-15
func CacheCovers(models []*statements.Cover, request string) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	for i, data := range models {
		st := reflect.TypeOf(data)
		sv := reflect.ValueOf(data)
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			if tag, ok := field.Tag.Lookup("json"); ok {
				if tag == "" {
					continue
				} else {
					temp[tag] = sv.Field(i).Interface()
				}
			} else {
				continue
			}
		}
		key := request + "/cover" + strconv.Itoa(i)
		redisDB.HMSet(key, temp)
	}
}

//缓存点歌排序
//request="流派"+"排序要求"
//例："日语综合"="japanese/composite"
//例："推荐最新"="recommend/latest"
//索引时按上述说法进行get，例如 HMGet("japanese/composite/selction1","avatar"),根据设计i=1-15
func CacheSelections(records []*statements.Selection, request string) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	for i, data := range records {
		st := reflect.TypeOf(data)
		sv := reflect.ValueOf(data)
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			if tag, ok := field.Tag.Lookup("json"); ok {
				if tag == "" {
					continue
				} else {
					temp[tag] = sv.Field(i).Interface()
				}
			} else {
				continue
			}
		}
		key := request + "/cover" + strconv.Itoa(i)
		redisDB.HMSet(key, temp)
	}
}

//缓存用户数据，尤其是点歌用的积分
func CacheUser(record *statements.User) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	st := reflect.TypeOf(record)
	sv := reflect.ValueOf(record)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok := field.Tag.Lookup("json"); ok {
			tempkey := strconv.Itoa(int(record.ID)) + "points"
			if tag == "" {
				continue
			} else if tag == "points" {
				redisDB.HSet(tempkey, "points", sv.Field(i).Int())
			} else if tag == "record" {
				redisDB.HSet(tempkey, "record", sv.Field(i).Int())
			} else {
				temp[tag] = sv.Field(i).Interface()
			}
		} else {
			continue
		}
		key := "user" + strconv.Itoa(int(record.ID))
		redisDB.HMSet(key, temp)
	}
}

//系统启动时直接初始化，加载任务文本时直接从redis读取
func CacheTask(records []*statements.Task) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	for i, data := range records {
		st := reflect.TypeOf(data)
		sv := reflect.ValueOf(data)
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			if tag, ok := field.Tag.Lookup("json"); ok {
				if tag == "" {
					continue
				} else {
					temp[tag] = sv.Field(i).Interface()
				}
			} else {
				continue
			}
		}
		key := "task" + strconv.Itoa(i)
		redisDB.HMSet(key, temp)
	}
}

//缓存所有礼品
func CachePrizes(lotteries []*statements.Lottery) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	for i, data := range lotteries {
		st := reflect.TypeOf(data)
		sv := reflect.ValueOf(data)
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			if tag, ok := field.Tag.Lookup("json"); ok {
				if tag == "" {
					continue
				} else {
					temp[tag] = sv.Field(i).Interface()
				}
			} else {
				continue
			}
		}
		key := "lottery" + strconv.Itoa(i)
		redisDB.HMSet(key, temp)
	}
}

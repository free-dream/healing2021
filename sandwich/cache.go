package sandwich

import (
	"fmt"
	"reflect"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

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
			if alias, ok := field.Tag.Lookup("json"); ok {
				if alias == "" {
					continue
				} else {
					temp[alias] = sv.Field(i).Interface()
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
func CacheUser(record *statements.User, openid string) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	st := reflect.TypeOf(record)
	sv := reflect.ValueOf(record)
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
		key := openid
		redisDB.HMSet(key, temp)
	}
}

//sql缓存到redis
func (middle *Middle) Cache(table) {
	redisCli := setting.RedisConn()
	fmt.Println(redisCli)
	// redisDb := setting.RedisClient()
	// mysqlDb := setting.MysqlConn()

	// mysqlDb.
}

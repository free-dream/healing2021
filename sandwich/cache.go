package sandwich

import (
	"fmt"
	"reflect"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//缓存排序
//request="流派"+"排序要求"
//例："日语"+"综合"
//例："推荐"+"最新"
func CacheCovers(models []*statements.Cover, request string) {
	redisDB := setting.RedisConn()
	temp := make(map[string]interface{})
	for _, data := range models {
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
		key := "cover" + strconv.Itoa(i)
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

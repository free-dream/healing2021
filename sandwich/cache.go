package sandwich

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//将对应的表生成键值对
func GenerateKV(model *gorm.Model) {
	st := reflect.TypeOf(model)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("json"); ok {
			if alias == "" {
				continue
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
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

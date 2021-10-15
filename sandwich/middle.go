package sandwich

import (
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//init初始化
func init() {
	mysql := setting.MysqlConn()
	redis := setting.RedisConn()
}

//redis mail
type mail struct {
	key      string
	value    interface{}
	mailtype string
	next     *mail
}

//sqlHandler主体
type middle struct {
	mailbox *mail
	counter int
}

//类型断言，用于对储存数据塑形
func valueCheckInt(value interface{}) (int, bool) {
	if data, ok := value.(int); ok {
		return data, ok
	} else {
		return 0, false
	}
}
func valueCheckString(value interface{}) (string, bool) {
	if data, ok := value.(string); ok {
		return data, ok
	} else {
		return "", false
	}
}

//sql缓存到redis
func (middle *middle) cache() {}

//redis持久化到sql
func (middle *middle) update() {
	//读取msg列表,提取对应的model
	//上传model
}

package middle

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

}

package middle

//sqlHandler主体
type middle struct{}

//sql缓存到redis
func (middle *middle) cache() {}

//redis持久化到sql
func (middle *middle) update() {}

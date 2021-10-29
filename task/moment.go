package task

const (
	MOMENT    float32 = 1.0
	MOMENTMAX float32 = 8.0
)

//动态关联任务
type MomentTask struct {
	field string
}

//有上限的数据
func (s *MomentTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.field)
	if target < 0 {
		if err := CacheTask(userid, s.field, 0); err != nil {
			return false
		}
	} else if target >= int(MOMENTMAX) {
		return false
	}
	return true
}

//更新任务缓存和数据
func (s *MomentTask) AddRecord(userid int) bool {
	if s.CheckMax(userid) {
		if ChangePoints(MOMENT, userid) {
			err := UpdateTask(userid, s.field, 1)
			if err == nil {
				return true
			}
		}
	}
	return false
}

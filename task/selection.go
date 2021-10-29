package task

const (
	SELECTION    float32 = 1.0
	SELECTIONMAX float32 = 10.0
)

//点歌关联任务
type SelectionTask struct {
	field string
}

//有上限的数据
func (s *SelectionTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.field)
	if target < 0 {
		if err := CacheTask(userid, s.field, 0); err != nil {
			return false
		}
	} else if target >= int(SELECTIONMAX) {
		return false
	}
	return true
}

//更新任务缓存和数据
func (s *SelectionTask) AddRecord(userid int) bool {
	if s.CheckMax(userid) {
		if ChangePoints(SELECTION, userid) {
			err := UpdateTask(userid, s.field, 1)
			if err == nil {
				return true
			}
		}
	}
	return false
}

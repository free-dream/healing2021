package task

const (
	MOOD    float32 = 1.0
	MOODMAX float32 = 5.0
	MOFIELD         = "mood"
)

//分享心情任务,产品暂时没有需求，挂起
type MoodTask struct {
	field string
}

//有上限的数据
func (s *MoodTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.field)
	if target < 0 {
		if err := CacheTask(userid, s.field, 0); err != nil {
			return false
		}
	} else if target >= int(MOODMAX) {
		return false
	}
	return true
}

//更新任务缓存和数据
func (s *MoodTask) AddRecord(userid int) bool {
	if s.CheckMax(userid) {
		if ChangePoints(MOOD, userid, s.field) {
			err := UpdateTask(userid, s.field, 1)
			if err == nil {
				return true
			}
		}
	}
	return false
}

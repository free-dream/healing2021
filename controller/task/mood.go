package task

import "git.100steps.top/100steps/healing2021_be/sandwich"

const (
	MOOD    float32 = 1.0
	MOODMAX float32 = 5.0
	MTID            = 3
)

//分享心情任务,产品暂时没有需求
type MoodTask struct {
	TID int
}

//有上限的数据
func (s *MoodTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.TID)
	if target < 0 {
		if err := sandwich.CacheTask(userid, s.TID, 0); err != nil {
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
		if ChangePoints(MOOD, userid, s.TID) {
			err := sandwich.UpdateTask(userid, s.TID, 1)
			if err == nil {
				return true
			}
		}
	}
	return false
}

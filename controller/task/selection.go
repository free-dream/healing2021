package task

import "git.100steps.top/100steps/healing2021_be/sandwich"

const (
	SELECTION    float32 = 1.0
	SELECTIONMAX float32 = 8.0
	STID                 = 1
)

//点歌关联任务
type SelectionTask struct {
	TID int
}

//有上限的数据
func (s *SelectionTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.TID)
	if target < 0 {
		if err := sandwich.CacheTask(userid, s.TID, 0); err != nil {
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
		if ChangePoints(SELECTION, userid, s.TID) {
			err := sandwich.UpdateTask(userid, s.TID, 1)
			if err == nil {
				return true
			}
		}
	}
	return false
}

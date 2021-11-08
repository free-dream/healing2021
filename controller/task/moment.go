package task

import "git.100steps.top/100steps/healing2021_be/sandwich"

const (
	MOMENT    float32 = 1.0
	MOMENTMAX float32 = 8.0
	MFIELD            = "moment"
)

//动态关联任务
type MomentTask struct {
	TID int
}

//有上限的数据
func (s *MomentTask) CheckMax(userid int) bool {
	target := GetCacheTask(userid, s.TID)
	if target < 0 {
		if err := sandwich.CacheTask(userid, s.TID, 0); err != nil {
			return false
		}
	} else if target >= int(MOMENTMAX) {
		return false
	}
	return true
}

//更新任务缓存和数据
func (s *MomentTask) AddRecord(userid int) error {
	if s.CheckMax(userid) {
		return ChangePoints(userid, s.TID, MOMENT)
	}
	return nil
}

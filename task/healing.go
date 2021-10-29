package task

//积分，不给的话就是不设上限
const (
	HEALING    float32 = 3.0
	HEALINGMAX float32 = -1.0
)

//翻唱、治愈关联任务
type HealingTask struct{}

//不设限制，常数标记为-1
func (h *HealingTask) CheckMax(userid int) bool {
	return true
}

//直接调用，不用检查是否超出上限
func (h *HealingTask) AddRecord(userid int) bool {
	return ChangePoints(HEALING, userid)
}

package task

const (
	SIGN    float32 = 1.0
	SIGNMAX float32 = 8.0
	SITID           = -1
)

//签到、转发任务,待定，根据产品需求暂时取消
type SignTask struct {
	TID int
}

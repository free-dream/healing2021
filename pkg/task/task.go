package task

//具体任务参数
//本次任务目前开来都是一次性的，没有计数要求

//积分和积分上限要求
const (
	SELECTION    = 1
	SELECTIONMAX = 10
)
const (
	HEALING = 3
)

const (
	SIGN    = 1
	SIGNMAX = 8
)
const (
	MOOD    = 1
	MOODMAX = 5
)
const (
	MOMENT    = 1
	MOMENTMAX = 8
)

//一次性任务
type MetaOTask interface {
	AddRecord() bool
}

//计数型任务,需要记录任务进度和检查是否完成
type MetaCTask interface {
	AddRecord() bool
	Counter() bool
	Check() bool
}

//点歌
type SelectionTask struct{}

//翻唱、治愈任务
type HealingTask struct{}

//签到、转发任务
type SignTask struct{}

//分享心情任务,待定
type MoodTask struct{}

//分享动态任务
type MomentTask struct{}

func (s *SelectionTask) AddRecord() {}

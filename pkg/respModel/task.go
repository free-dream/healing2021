package respModel

//任务信息
type TaskResp struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Max  int    `json:"max"`
}

//抽奖返回
type TaskTableResp struct {
	Task    TaskResp
	Counter int `json:"counter"`
}

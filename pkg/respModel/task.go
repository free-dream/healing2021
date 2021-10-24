package respModel

//彩票信息
type TaskResp struct {
	ID     int    `json:"id"`
	Text   string `json:"string"`
	Target int    `json:"target"`
}

//抽奖返回
type TaskTableResp struct {
	Check   int `json:"check"`
	Task    TaskResp
	Counter int `json:"counter"`
}

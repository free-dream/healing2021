package respModel

//彩票信息
type LotteryResp struct {
	Picture     string `json:"string"`
	Name        string `json:"name"`
	Possibility int    `json:"possibility"`
}

//抽奖返回
type DrawResp struct {
	Check   int    `json:"check"`
	Name    string `json:"name"`
	Picture string `json:"string"`
}

//用户中奖记录
type PrizesResp struct {
	Name    string `json:"name"`
	Picture string `json:"string"`
}

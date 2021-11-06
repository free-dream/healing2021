package respModel

const (
	Msg0 = "抽奖成功，请耐心等待开奖"
	Msg1 = "不能重复抽奖"
	Msg2 = ""
)

//彩票信息
type LotteryResp struct {
	Picture     string  `json:"picture"`
	Name        string  `json:"name"`
	Possibility float32 `json:"possibility"`
}

//抽奖信息
type DrawResp struct {
	Msg string `json:"msg"`
}

//用户中奖记录
type PrizesResp struct {
	Name    string `json:"name"`
	Picture string `json:"string"`
}

//用户积分
type PointsResp struct {
	Points int `json:"points"`
}

// //在线抽奖返回，废案
// type DrawResp struct {
// 	Check   int    `json:"check"`
// 	Name    string `json:"name"`
// 	Picture string `json:"string"`
// }

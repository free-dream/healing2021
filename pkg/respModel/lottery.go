package respModel

const (
	Msg0 = "积分不足"
	Msg1 = "请填写手机号码"
	Msg2 = "已参与抽奖"
)

//彩票信息
type LotteryResp struct {
	Picture     string `json:"string"`
	Name        string `json:"name"`
	Possibility int    `json:"possibility"`
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

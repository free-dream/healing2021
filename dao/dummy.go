package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

const (
	SCUT  = "华南理工大学"
	SYU   = "中山大学"
	JU    = "暨南大学"
	SCNU  = "华南师范大学"
	OTHER = "其它大学"

	TARGET1 = "bbt2021ad" //靶数据，主要用于查询测试
	TARGET2 = "bbt1021ad"
	TARGET3 = "bbt21bc"

	PRIZESP = "特奖" //目前设计四个奖项，特奖，一等奖，二等奖，三等奖
	PRIZE1  = "一等奖"
	PRIZE2  = "二等奖"
	PRIZE3  = "三等奖"
)

var (
	SchoolPool = []string{SCUT, SYU, JU, SCNU, OTHER}
	TargetPool = []string{TARGET1, TARGET2, TARGET3}
)

//生成dummy用户
func dummyUser() *statements.User {
	check1 := tools.GetRandomNumbers(4)
	check2 := tools.GetRandomNumbers(5)
	var nickname, school string
	//决定nickname
	if check1 != 3 {
		nickname = TargetPool[check1]
	} else {
		nickname = string(tools.GetRandomString(4))
	}
	//决定学校
	school = SchoolPool[check2]
	user := statements.User{
		Openid:    string(tools.GetRandomString(10)),
		Nickname:  nickname,
		RealName:  string(tools.GetRandomString(6)),
		Signature: string(tools.GetRandomString(20)),
		School:    school,
	}

	return &user
}

// //假彩票
// func fakeLotteries() *statements.Lottery {

// }

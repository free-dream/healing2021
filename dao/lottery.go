package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

var (
	RedisDb   = db.RedisConn()
	MysqlDb   = db.MysqlConn()
	usertable = tables.User{}
	lotteries = tables.Lottery{}
)

//不展示奖品归属
func GetAllLotteries() {
	var prize statements.Prize
}

//根据lottery里奖品的归属拉取奖品列表
func GetPrizes() {}

func UpdateUserPoints() bool {
	//获取当前用户
	//-->redis里搜索/更新，成功则尚需要持久化
	//-->mysql里搜索/更新，若上述行不通
	//更新成功/失败把结果抛给调用者
	return true
}

func UpdateLotterybox() bool {
	//没中的情形：抽中已抽中的奖品
	return true
}

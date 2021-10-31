package dao

//废案方法
import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

var (
	RedisDb = db.RedisConn()
	MysqlDb = db.MysqlConn()
)

//获取所有奖品，不展示奖品归属
func GetAllLotteries() ([]tables.Lottery, error) {
	var lotteries []tables.Lottery
	err := MysqlDb.Find(&lotteries).Error
	if err != nil {
		return nil, err
	}
	return lotteries, err
}

//根据lottery里奖品的归属拉取奖品列表
func GetPrizesById(userid int) ([]tables.Lottery, error) {
	var prizes []tables.Lottery
	err := MysqlDb.Where("UserId = ?", userid).Find(&prizes).Error
	if err != nil {
		return nil, err
	}
	return prizes, nil
}

//抽奖确认
func DrawCheck(userid int) {}

//我的回合，抽卡！#废案#
func Draw(id int) (tables.Lottery, error) {
	var target tables.Lottery
	err := MysqlDb.Where("Id = ?", id).Find(target).Error
	if err != nil {
		return target, err
	}
	return target, nil
}

//更新奖品归属
func UpdateLotterybox(lotteryid int, userid int) (bool, error) {
	var target tables.Lottery
	err := MysqlDb.Where("ID = ?", lotteryid).Find(target).Error
	if err != nil {
		return false, err
	}
	target.UserId = userid
	MysqlDb.Save(target)
	return true, nil
}

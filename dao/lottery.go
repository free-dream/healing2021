package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/jinzhu/gorm"
)

const (
	MINPOINTS = 10.0
)

var (
	MysqlDb = db.MysqlConn()
)

//先从缓存里拿，拿不到再读取数据库，后增加缓存
func GetUserPoints(userid int) (int, error) {
	temp := sandwich.GetCachePoints(userid)
	if temp < 0 {
		var user tables.User
		err := MysqlDb.Where("id = ?", userid).First(&user).Error
		if err != nil {
			return -1, err
		}
		err = sandwich.UpdateCachePoints(userid, user.Points)
		if err != nil {
			return -1, err
		}
		return user.Points, err
	}
	return temp, nil
}

//获取用户的nickname
func GetUserNickname(userid int) (string, error) {
	var user tables.User
	err := MysqlDb.Where("id = ?", userid).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Nickname, nil
}

//获取用户的avatar
func GetUserAvatar(userid int) (string, error) {
	var user tables.User
	err := MysqlDb.Where("id = ?", userid).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Avatar, nil
}

//展示所有奖品
func GetAllLotteries() ([]tables.Lottery, error) {
	var lotteries []tables.Lottery
	err := MysqlDb.Find(&lotteries).Error
	if err != nil {
		return nil, err
	}
	return lotteries, err
}

//抽奖确认
func DrawCheck(userid int) (int, error) {
	// points := sandwich.GetCachePoints(userid)
	var err error
	// if points < 0 {
	// 	var user tables.User
	// 	err = MysqlDb.Where("ID = ?", userid).First(&user).Error
	// 	if err != nil {
	// 		return -1, err
	// 	}
	// 	points = user.Points
	// }
	// if points < MINPOINTS {
	// 	return 0, nil
	// }
	var prize statements.Prize
	if err = MysqlDb.Where("user_id = ?", userid).First(&prize).Error; gorm.IsRecordNotFoundError(err) {
		//未抽奖
		return 0, nil
	} else if err != nil {
		return -1, err
	}
	//已抽奖
	return 1, nil
}

//为prize表增加一条记录，用于分配奖品
func CreatePrize(userid int, tel string) error {
	prize := tables.Prize{
		UserId: userid,
		Tel:    tel,
	}
	err := MysqlDb.Create(&prize).Error
	return err
}

// //根据lottery里奖品的归属拉取奖品列表 废案
// func GetPrizesById(userid int) ([]tables.Lottery, error) {
// 	var prizes []tables.Lottery
// 	err := MysqlDb.Where("UserId = ?", userid).Find(&prizes).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return prizes, nil
// }

// //更新奖品归属
// func UpdateLotterybox(lotteryid int, userid int) (bool, error) {
// 	var target tables.Lottery
// 	err := MysqlDb.Where("ID = ?", lotteryid).Find(target).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	target.UserId = userid
// 	MysqlDb.Save(target)
// 	return true, nil
// }

// 线上抽奖 ，临时废案#废案#
// func Draw(id int) (tables.Lottery, error) {
// 	var target tables.Lottery
// 	err := MysqlDb.Where("Id = ?", id).Find(target).Error
// 	if err != nil {
// 		return target, err
// 	}
// 	return target, nil
// }

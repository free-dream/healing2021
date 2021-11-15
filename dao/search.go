package dao

import (
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

// const (
// 	PAGESIZE = 10
// 	PAGE     = 3
// )

//目前只支持单关键字查询
//考虑到Username的
// func SearchUserByKeyWords(keyword ...string) ([]statements.User, int, error) {
// 	mysqlDb := setting.MysqlConn()
// 	db := mysqlDb.Limit(10).Offset((PAGE - 1) * PAGESIZE)
// 	len := len(keyword)
// 	var user []statements.User
// 	var err error
// }

//为保证性能，只获取最多30条记录
//返回数据和数据长度
//仅能匹配歌名和用户名，无法进行风格或者语言的搜索

//根据电话号码获取用户
func SearchUserByTel(tel string) ([]respModel.UserResp, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []respModel.UserResp
	var counter int
	db := mysqlDb.Limit(30).
		Table("user").
		Select("avatar,nickname,signature,id").
		Where("phone_number = ? AND phone_search = ?", tel, 0).
		Find(&data)
	err := db.Error
	if gorm.IsRecordNotFoundError(err) {
		return data, 0, nil
	} else if err != nil {
		return nil, -1, err
	}
	db.Count(&counter)
	return data, counter, nil
}

//其它查询
func SearchUserByKeyword(keyword string) ([]respModel.UserResp, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []respModel.UserResp
	db1 := mysqlDb.Limit(30).
		Table("user").
		Select("avatar,nickname,signature,id").
		Where("(real_name_search = ? AND real_name = ?) OR nickname = ?", 0, keyword, keyword).
		Find(&data)
	err := db1.Error
	if gorm.IsRecordNotFoundError(err) {
		return data, 0, nil
	} else if err != nil {
		return nil, -1, err
	}
	return data, len(data), nil
}

func SearchCoverByKeyword(keyword string) ([]respModel.CoversResp, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []respModel.CoversResp
	db := mysqlDb.Limit(30).
		Table("cover").
		Select("id,avatar,song_name,nickname,classic_id,module,selection_id,created_at").
		Where("song_name = ?", keyword).
		Find(&data)
	err := db.Error
	if gorm.IsRecordNotFoundError(err) {
		return data, 0, nil
	} else if err != nil {
		return nil, -1, err
	}
	return data, len(data), nil
}

func SearchSelectionByKeyword(keyword string) ([]respModel.SelectionResp, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []respModel.SelectionResp
	db := mysqlDb.Limit(30).
		Table("selection").
		Select("avatar,nickname,created_at,song_name,id").
		Where("song_name = ?", keyword).
		Find(&data)
	err := db.Error
	if gorm.IsRecordNotFoundError(err) {
		return data, 0, nil
	} else if err != nil {
		return nil, -1, err
	}
	return data, len(data), nil
}

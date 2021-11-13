package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
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
func SearchUserByTel(tel string) ([]statements.User, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []statements.User
	var counter int
	db := mysqlDb.Limit(30).Where("phone_number = ? AND phone_search = ?", tel, 1).Find(&data)
	err := db.Error
	if err != nil {
		return nil, -1, err
	}
	db.Count(&counter)
	return data, counter, nil
}

//其它查询
func SearchUserByKeyword(keyword string) ([]statements.User, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []statements.User
	var counter int
	db := mysqlDb.Limit(30).Where("nickname LIKE ?", "%"+keyword+"%").Find(&data)
	err := db.Error
	if err != nil {
		return nil, -1, err
	}
	db.Count(&counter)
	return data, counter, nil
}

func SearchCoverByKeyword(keyword string) ([]statements.Cover, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []statements.Cover
	var counter int
	db := mysqlDb.Limit(30).Where("song_name LIKE ?", "%"+keyword+"%").Find(&data)
	err := db.Error
	if err != nil {
		return nil, -1, err
	}
	db.Count(&counter)
	return data, counter, nil
}

func SearchSelectionByKeyword(keyword string) ([]statements.Selection, int, error) {
	mysqlDb := setting.MysqlConn()
	var data []statements.Selection
	var counter int
	db := mysqlDb.Limit(30).Where("song_name LIKE ?", "%"+keyword+"%").Find(&data)
	err := db.Error
	if err != nil {
		return nil, -1, err
	}
	db.Count(&counter)
	return data, counter, nil
}

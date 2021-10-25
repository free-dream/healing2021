package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//基于学校获取排名，仅有sql版
func GetRankingBySchool(school string) ([]tables.User, error) {
	mysqlDb := db.MysqlConn()

	var users []tables.User
	var err error = nil
	if school != "All" {
		//按Record 倒序排列
		err = mysqlDb.Where("School = ?", school).Order("Record desc").Limit(10).Find(&users).Error
	} else if school == "All" {
		err = mysqlDb.Order("Record desc").Limit(10).Find(&users).Error
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 获取排名，有一点小问题
// func GetRankByUserId(userid int) (tables.User, error) {
// 	// mysqlDb := db.MysqlConn()

// 	// var user tables.User
// 	// err := mysqlDb.Where("ID = ?", userid)
// 	return nil, nil
// }

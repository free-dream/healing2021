package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	user tables.User
)

//直接从mysql读取积分和记录
func GetPoints(userid int) (int, error) {
	mysqlDb := db.MysqlConn()

	err := mysqlDb.Where("Id = ?", userid).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Points, err
}

//基于redis数据更新用户数据
func UpdatePoints(userid int, record int, point int) {
	mysqlDb := db.MysqlConn()

	mysqlDb.Model(&user).Where("Id = ?", userid).Update("Points", gorm.Expr("Points + ?", point))
	mysqlDb.Model(&user).Where("Id = ?", userid).Update("Record", gorm.Expr("Record + ?", record))

	return
}

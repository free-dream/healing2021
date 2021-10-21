package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

//基于redis数据更新用户数据
func UpdatePoints(userid int, record int, point int) {
	mysqlDb := db.MysqlConn()

	var user tables.User
	mysqlDb.Model(&user).Where("Id = ?", userid).Update("Points", gorm.Expr("Points + ?", point))
	mysqlDb.Model(&user).Where("Id = ?", userid).Update("Record", gorm.Expr("Record + ?", record))

	return
}

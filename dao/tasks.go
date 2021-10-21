package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

func UpdateTasks(userid int, taskid int, process int, check int) {
	mysqlDb := db.MysqlConn()
	var tasktable tables.TaskTable
	mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Counter", gorm.Expr("IsLiked + ?", process))
	mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Check", check)
}

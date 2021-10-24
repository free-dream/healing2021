package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

//基于给定数据更新task_table
func UpdateTasks(userid int, taskid int, process int, check int) {
	mysqlDb := db.MysqlConn()
	var tasktable tables.TaskTable
	mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Counter", gorm.Expr("IsLiked + ?", process))
	mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Check", check)
}

//提取task_table
func GetTasktables(userid int) ([]tables.TaskTable, error) {
	mysqlDb := db.MysqlConn()
	var tasktables []tables.TaskTable
	err := mysqlDb.Where("UserId = ?", userid).Find(&tasktables).Error
	if err != nil {
		return nil, err
	}
	return tasktables, nil
}

//提取任务信息
func GetTasks(taskid int) (tables.Task, error) {
	mysqlDb := db.MysqlConn()
	var task tables.Task
	err := mysqlDb.Where("ID = ?", taskid).Find(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

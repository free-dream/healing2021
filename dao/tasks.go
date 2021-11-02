package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

//基于给定数据更新task_table
func UpdateTasks(userid int, taskid int, process int) {
	mysqlDb := db.MysqlConn()
	var tasktable tables.TaskTable
	mysqlDb.Model(&tasktable).Where("user_id = ? AND task_id = ?", userid, taskid).UpdateColumn("Counter", gorm.Expr("IsLiked + ?", process))
	// mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Check", check)
}

//基于给定数据生成tasktable,用于用户初始化
func GenerateTasktable(tids []int, userid int) error {
	mysqlDb := db.MysqlConn()
	for tid := range tids {
		task := tables.TaskTable{
			TaskId:  tids[tid],
			UserId:  userid,
			Counter: 0,
		}
		err := mysqlDb.Create(&task).Error
		if err != nil {
			return err
		}
	}

	return nil
}

//任务建立查重,仅用于假登录
func CheckTasks(userid int) (bool, error) {
	mysqlDb := db.MysqlConn()
	var data []tables.TaskTable
	err := mysqlDb.Where("user_id = ?", userid).Find(&data).Error
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

//提取task_table
func GetTasktables(userid int) ([]tables.TaskTable, error) {
	mysqlDb := db.MysqlConn()
	var tasktables []tables.TaskTable
	err := mysqlDb.Where("user_id = ?", userid).Find(&tasktables).Error
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

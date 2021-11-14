package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//基于给定数据生成tasktable,用于用户初始化
func GenerateTasktable(tids []int, userid int) error {
	mysqlDb := db.MysqlConn()
	//补一个任务查重
	var temp []statements.TaskTable
	mysqlDb.Where("user_id = ?", userid).Find(&temp)
	if len(temp) > 0 {
		return nil
	}
	//
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

//更新任务积分
func UpdateTaskPoints(userid int, taskid int, points int, tpoints int) error {
	mysqlDb := db.MysqlConn()
	var user tables.User
	var tasktable tables.TaskTable

	err := mysqlDb.Where("id = ?", userid).First(&user).Error
	if err != nil {
		return err
	}

	err = mysqlDb.Where("task_id = ? AND user_id = ?", taskid, userid).Find(&tasktable).Error
	if err != nil {
		return err
	}

	user.Points = points + user.Points
	err = mysqlDb.Save(&user).Error
	if err != nil {
		return err
	}

	tasktable.Counter = tpoints + tasktable.Counter
	err = mysqlDb.Save(&tasktable).Error
	if err != nil {
		return err
	}
	return nil
}

// //基于给定数据更新task_table
// func UpdateTasks(userid int, taskid int, process int) {
// 	mysqlDb := db.MysqlConn()
// 	var tasktable tables.TaskTable
// 	mysqlDb.Model(&tasktable).Where("user_id = ? AND task_id = ?", userid, taskid).UpdateColumn("Counter", gorm.Expr("IsLiked + ?", process))
// 	// mysqlDb.Model(&tasktable).Where("UserId = ? AND TaskId = ?", userid, taskid).UpdateColumn("Check", check)
// }

package statements

import (
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/sandwich"
)

func TableClean() {
	db := setting.MysqlConn()
	db.Exec("truncate table advertisement")
	//db.Exec("truncate table classic")
	//db.Exec("truncate table devotion")
	db.Exec("truncate table cover")
	db.Exec("truncate table praise")
	db.Exec("truncate table lottery")
	db.Exec("truncate table message")
	db.Exec("truncate table moment")
	db.Exec("truncate table moment_comment")
	db.Exec("truncate table selection")
	db.Exec("truncate table task")
	db.Exec("truncate table task_table")
	db.Exec("truncate table user")
	db.Exec("truncate table sysmsg")
	db.Exec("truncate table usrmsg")
	sandwich.Clean()
	sandwich.CleanSession()
}

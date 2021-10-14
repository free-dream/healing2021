package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type TaskTable struct {
	*gorm.Model
	TaskId  int `gorm:"default:0"`
	UserId  int `gorm:"default:0"`
	Check   int `gorm:"default:0"`
	Counter int `gorm:"default:0"`
}

func TaskTableInit() {
	db := setting.MysqlConn()
	setting.TimeSetting("task_table")
	if !db.HasTable(&TaskTable{}) {
		if err := db.CreateTable(&TaskTable{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table TaskTable has been created")
	} else {
		db.AutoMigrate(&TaskTable{})
		fmt.Println("TaskTable User has existed")
	}
}

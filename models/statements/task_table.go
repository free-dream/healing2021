package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type TaskTable struct {
	gorm.Model
	TaskId int `gorm:"default:0;index"`
	UserId int `gorm:"default:0;index"`
	// Check   int `gorm:"default:0"` //是否达到上限
	Counter int `gorm:"default:0"` //已获得积分
}

func TaskTableInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&TaskTable{}) {
		if err := db.CreateTable(&TaskTable{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table TaskTable has been created")
	} else {
		db.AutoMigrate(&TaskTable{})
		fmt.Println("Table TaskTable has existed")
	}
	setting.TimeSetting("task_table")
}

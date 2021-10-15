package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Text   string `gorm:"default:''"`
	Target int    `gorm:"default:0"`
}

func TaskInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Task{}) {
		if err := db.CreateTable(&Task{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Task has been created")
	} else {
		db.AutoMigrate(&Task{})
		fmt.Println("Table Task has existed")
	}
	setting.TimeSetting("task")
}

package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	SenderId int `gorm:"default:0;index"`
	TakerId  int `gorm:"default:0;index"`
	Content  int `gorm:"default:0"`
}

func MessageInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Message{}) {
		if err := db.CreateTable(&Message{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Message has been created")
	} else {
		db.AutoMigrate(&Message{})
		fmt.Println("Table Message has existed")
	}
	setting.TimeSetting("message")
}

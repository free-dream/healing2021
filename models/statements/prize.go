package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Prize struct {
	gorm.Model
	UserId int    `gorm:"default:0;index" json:"user_id"`
	Tel    string `gorm:"default:0" json:"tel"`
}

func PrizeInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Prize{}) {
		if err := db.CreateTable(&Prize{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Prize has been created")
	} else {
		db.AutoMigrate(&Prize{})
		fmt.Println("Table prize has existed")
	}
	setting.TimeSetting("prize")
}

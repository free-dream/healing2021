package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Lottery struct {
	gorm.Model
	Name        string  `gorm:"default:''" json:"name"`
	Picture     string  `gorm:"default:''" json:"picture"`
	Possibility float64 `gorm:"default:0" json:"possibility"`
}

func LotteryInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Lottery{}) {
		if err := db.CreateTable(&Lottery{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Lottery has been created")
	} else {
		db.AutoMigrate(&Lottery{})
		fmt.Println("Table Lottery has existed")
	}
	setting.TimeSetting("lottery")
}

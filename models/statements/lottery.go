package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Lottery struct {
	gorm.Model
	UserId      int     `gorm:"default:-1;index"`
	Name        string  `gorm:"default:''" json:"Name"`
	Picture     string  `gorm:"default:''" json:"Picture"`
	Possibility float64 `gorm:"default:0" json:"Possiblity"`
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

package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Lottery struct {
	gorm.Model
	UserId      int     `gorm:"default:-1 index"`
	Name        string  `gorm:"default:''"`
	Picture     string  `gorm:"default:''"`
	Possibility float64 `gorm:"default:0"`
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

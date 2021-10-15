package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Advertisement struct {
	gorm.Model
	Url     string `gorm:"default:''"`
	Address string `gorm:"default:''"`
	Weight  int    `gorm:"default:0"`
}

func AdvertisementInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Advertisement{}) {
		if err := db.CreateTable(&Advertisement{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Advertisement has been created")
	} else {
		db.AutoMigrate(&Advertisement{})
		fmt.Println("Table Advertisement has existed")
	}
	setting.TimeSetting("advertisement")
}

package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Moment struct {
	*gorm.Model
	Content string `gorm:"default:''"`
	SongId  int    `gorm:"default:0"`
	States  string `gorm:"default:''"`
	Picture string `gorm:"default:''"`
}

func MomentInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Moment{}) {
		if err := db.CreateTable(&Moment{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Moment has been created")
	} else {
		db.AutoMigrate(&Moment{})
		fmt.Println("Table Moment has existed")
	}
	setting.TimeSetting("moment")
}

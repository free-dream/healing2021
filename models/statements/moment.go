package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Moment struct {
	gorm.Model
	UserId      int    `gorm:"default:0;index"`
	Content     string `gorm:"default:''"`
	SongName    string `gorm:"default:''"`
	SelectionId int    `gorm:"default:0;index"`
	ClassicId   int    `gorm:"default:0;index"`
	Module      int    `gorm:"default:2;index"` // 0经典点歌，1童年分享，2无歌曲
	State       string `gorm:"default:''"`
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

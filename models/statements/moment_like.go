package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type MomentLike struct {
	*gorm.Model
	UserId   int `gorm:"default:0"`
	MomentId int `gorm:"default:0"`
	IsLiked  int `gorm:"default:0"`
}

func MomentLikeInit() {
	db := setting.MysqlConn()
	setting.TimeSetting("moment_like")
	if !db.HasTable(&MomentLike{}) {
		if err := db.CreateTable(&MomentLike{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table MomentLike has been created")
	} else {
		db.AutoMigrate(&MomentLike{})
		fmt.Println("Table MomentLike has existed")
	}
}

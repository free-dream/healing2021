package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Praise struct {
	gorm.Model
	CoverId         int `gorm:"default:0;index"`
	UserId          int `gorm:"default:0;index"`
	IsLiked         int `gorm:"default:0;index"`
	MomentId        int `gorm:"default:0;index"`
	MomentCommentId int `gorm:"default:0;index"`
}

func PraiseInit() {
	db := setting.MysqlConn()
	if !db.HasTable(&Praise{}) {
		if err := db.CreateTable(&Praise{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table praise has been created")
	} else {
		db.AutoMigrate(&Praise{})
		fmt.Println("Table praise has existed")
	}
	setting.TimeSetting("praise")
}

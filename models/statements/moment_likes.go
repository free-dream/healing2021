package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type MomentLike struct {
	*gorm.Model
	UserId  int    `gorm:"default:0"`
	MomentId  int    `gorm:"default:0"`
	IsLike  int    `gorm:"default:0"`
}

func MomentLikeInit() {
	db := setting.MysqlConn()
	if !db.HasTable(&Moment{}) {
		if err := db.CreateTable(&Moment{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table MomentLike has been created")
	} else {
		db.AutoMigrate(&Moment{})
		fmt.Println("Table MomentLike has existed")
	}
}

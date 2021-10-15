package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type MomentComment struct {
	*gorm.Model
	UserId   int `gorm:"default:0"`
	MomentId int `gorm:"default:0"`
	Comment  int `gorm:"default:0"`
}

func MomentCommentInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&MomentComment{}) {
		if err := db.CreateTable(&MomentComment{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table MomentComment has been created")
	} else {
		db.AutoMigrate(&MomentComment{})
		fmt.Println("Table MomentComment has existed")
	}
	setting.TimeSetting("moment_comment")
}

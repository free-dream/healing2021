package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type MomentCommentLike struct {
	*gorm.Model
	UserId   int `gorm:"default:0"`
	MomentId int `gorm:"default:0"`
	Comment  int `gorm:"default:0"`
}

func MomentCommentLikeInit() {
	db := setting.MysqlConn()
	setting.TimeSetting("moment_comment_like")
	if !db.HasTable(&MomentCommentLike{}) {
		if err := db.CreateTable(&MomentCommentLike{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table MomentCommentLike has been created")
	} else {
		db.AutoMigrate(&MomentCommentLike{})
		fmt.Println("Table MomentCommentLike has existed")
	}
}

package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type SongLike struct {
	*gorm.Model
	UserId   int	`gorm:"default:0"`
	SongId   int	`gorm:"default:0"`
	IsLiked  int  	`gorm:"default:0"`
}
func SongLikeInit() {
	db := setting.MysqlConn()
	if !db.HasTable(&SongLike{}) {
		if err := db.CreateTable(&SongLike{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table SongLike has been created")
	} else {
		db.AutoMigrate(&SongLike{})
		fmt.Println("Table SongLike has existed")
	}
}
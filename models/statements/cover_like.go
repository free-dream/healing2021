package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type CoverLike struct{
	*gorm.Model
	CoverId int    `gorm:"default:0"`
	UserId  int   `gorm:"default:0"`
	IsLiked int   `gorm:"default:0"`
}
func CoverLikeInit() {
	db := setting.MysqlConn()
	if !db.HasTable(&CoverLike{}) {
		if err := db.CreateTable(&CoverLike{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table CoverLike has been created")
	} else {
		db.AutoMigrate(&CoverLike{})
		fmt.Println("Table CoverLike existed")
	}
}
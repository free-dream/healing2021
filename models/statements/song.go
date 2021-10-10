package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Song struct {
	*gorm.Model
	SelectionId int    `gorm:"default:0"`
	UserId      int    `gorm:default:0`
	Song        string `gorm:default:''`
	Name        string `gorm:default:''`
}

func SongInit() {
	db := setting.MysqlConn()
	if !db.HasTable(&Song{}) {
		if err := db.CreateTable(&Song{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Song has been created")
	} else {
		db.AutoMigrate(&Song{})
		fmt.Println("Table Song has existed")
	}
}

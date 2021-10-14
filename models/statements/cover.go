package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Cover struct {
	*gorm.Model
	UserId      int    `gorm:"default:0"`
	Avatar      string `gorm:"default:''"`
	SelectionId string `gorm:"default:''"`
	NickName    string `gorm:"default:''"`
	ClassicId   int    `gorm:"default:0"`
	Likes       int    `gorm:"default:0"`
	File        string `gorm:"default:''"`
}

func CoverInit() {
	db := setting.MysqlConn()
	setting.TimeSetting("cover")
	if !db.HasTable(&Cover{}) {
		if err := db.CreateTable(&Cover{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Cover has been created")
	} else {
		db.AutoMigrate(&Cover{})
		fmt.Println("Table Cover existed")
	}
}

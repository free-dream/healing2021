package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Selection struct {
	*gorm.Model
	SongName string `gorm:"default:''"`
	Remark   string
	Language string `gorm:"default:''"`
	Style    string `gorm:"default:''"`
	UserId   int    `gorm:"default:0"`
	Avatar   string `gorm:"default:''"`
	Module   string `gorm:"default:''"`
}

func SelectionInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Selection{}) {
		if err := db.CreateTable(&Selection{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Selection has been created")
	} else {
		db.AutoMigrate(&Selection{})
		fmt.Println("Table Selection has existed")
	}
	setting.TimeSetting("selection")
}

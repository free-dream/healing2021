package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Selection struct {
	gorm.Model
	SongName string `gorm:"default:''" json:"songname"`
	Remark   string `gorm:"default:''" json:"remark"`
	Language string `gorm:"default:''" json:"language"`
	Style    string `gorm:"default:''" json:"style"`
	UserId   int    `gorm:"default:0;index" json:"userid"`
	Avatar   string `gorm:"default:''" json:"avatar"`
	Module   string `gorm:"default:''" json:"module"`
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

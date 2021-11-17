package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Cover struct {
	gorm.Model
	UserId      int    `gorm:"default:0;index" json:"userid"`
	Nickname    string `gorm:"default:''" json:"nickname"`
	Avatar      string `gorm:"default:''" json:"avatar"`
	SelectionId int    `gorm:"default:0;index" json:"selection_id"`
	SongName    string `gorm:"default:''" json:"song_name"`
	ClassicId   int    `gorm:"default:0;index" json:"classic_id"`
	File        string `gorm:"default:''" json:"file"`
	Style       string `gorm:"default:''" json:"style"`
	Language    string `gorm:"default:''" json:"language"`
	Module      int    `gorm:"default:0" json:"module"`
	IsAnon      bool   `gorm:"default:0" json:"is_anon"`
}

func CoverInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Cover{}) {
		if err := db.CreateTable(&Cover{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Cover has been created")
	} else {
		db.AutoMigrate(&Cover{})
		fmt.Println("Table Cover has existed")
	}
	setting.TimeSetting("cover")
}

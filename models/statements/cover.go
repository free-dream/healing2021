package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Cover struct {
	gorm.Model
	UserId      int    `gorm:"default:0;index" json:"userid"`
	Avatar      string `gorm:"default:''" json:"avatar"`
<<<<<<< HEAD
	SelectionId string `gorm:"default:'';index" json:"selectionid"`
	SongName    string `gorm:"default:''" json:"songname"`
	Nickname	string 	`gorm:"default:''" json:"nickname"`
	ClassicId   int    `gorm:"default:0;index" json:"classicid"`
=======
	SelectionId string `gorm:"default:'';index" json:"selection_id"`
	SongName    string `gorm:"default:''" json:"song_name"`
	ClassicId   int    `gorm:"default:0;index" json:"classic_id"`
>>>>>>> 0a9c658cad10bccf1f7c09c239fe72665d5c8795
	Likes       int    `gorm:"default:0" json:"likes"`
	File        string `gorm:"default:''" json:"file"`
	Module      int    `gorm:"default:0" json:"module"`
	Style       string `gorm:"default:''" json:"style"`
	Language    string `gorm:"default:''" json:"language"`
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

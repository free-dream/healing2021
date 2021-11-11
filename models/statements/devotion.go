package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Devotion struct {
	gorm.Model
	SongName string `gorm:"default:''" json:"song_name"`
	File     string `gorm:"default:''" json:"file"`
	Singer   string `gorm:"default:''" json:"singer"`
}

func DevotionInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Devotion{}) {
		if err := db.CreateTable(&Devotion{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Devotion has been created")
	} else {
		db.AutoMigrate(&Devotion{})
		fmt.Println("Table Devotion has existed")
	}
	setting.TimeSetting("Devotion")
}

package statements

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Classic struct {
	gorm.Model
	Remark   string `gorm:"default:''"`
	SongName string `gorm:"default:''"`
	Icon     string `gorm:"default:''"`
	Singer   string `gorm:"default:''"`
	WorkName string `gorm:"default:''"`
	Click    int	`gorm:"default:0"`
	Lyrics   string `gorm:"default:''"`
}

func ClassicInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Classic{}) {
		if err := db.CreateTable(&Classic{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("TableClassic has been created")
	} else {
		db.AutoMigrate(&Classic{})
		fmt.Println("Table Classic existed")
	}
	setting.TimeSetting("classic")
}

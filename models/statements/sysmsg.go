package statements

import (
	"fmt"
	"time"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Sysmsg struct {
	Uid       uint   `gorm:not null`
	Type      int    `gorm:not null`
	ContentId uint   `gorm:not null`
	Song      string `gorm:not null`
	Time      time.Time
	IsSend    int
	gorm.Model
}

func SysmsgInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Sysmsg{}) {
		if err := db.CreateTable(&Sysmsg{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table sysmsg has been created")
	} else {
		db.AutoMigrate(&Sysmsg{})
		fmt.Println("Table sysmsg has existed")
	}
}

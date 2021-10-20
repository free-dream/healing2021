package statements

import (
	"fmt"
	"time"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Openid         string `gorm:"default: '';index"`
	Nickname       string `gorm:"default: ''"`
	RealName       string `gorm:"default: ''"`
	Signature      string `gorm:"default: ''"`
	Avatar         string `gorm:"default: ''"`
	PhoneNumber    string `gorm:"default: ''"`
	Sex            int    `gorm:"default: 0"`
	School         string `gorm:"default: ''"`
	Points         int    `gorm:"default: 0"`
	Record         int    `gorm:"default: 0"`
	Background     string `gorm:"default: ''"`
	AvatarVisible  int    `gorm:"default:0"`
	PhoneSearch    int    `gorm:"default:0"`
	RealNameSearch int    `gorm:"default:0"`
	LoginTime      time.Time
}

func UserInit() {
	db := setting.DB

	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table User has been created")
	} else {
		db.AutoMigrate(&User{})
		fmt.Println("Table User has existed")
	}
	setting.TimeSetting("user")
}

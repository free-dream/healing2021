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
	Nickname       string `gorm:"default: ''" json:"nickname"`
	RealName       string `gorm:"default: ''" json:"real_name"`
	Signature      string `gorm:"default: ''" json:"signature"`
	Avatar         string `gorm:"default: ''" json:"avatar"`
	PhoneNumber    string `gorm:"default: ''" json:"phone_number"`
	Sex            int    `gorm:"default: 0" json:"sex"`
	School         string `gorm:"default: ''" json:"school"`
	Points         int    `gorm:"default: 0" json:"points"`
	Record         int    `gorm:"default: 0" json:"record"`
	Background     string `gorm:"default: ''" json:"background"`
	AvatarVisible  int    `gorm:"default:0" json:"avatar_visible"`
	PhoneSearch    int    `gorm:"default:0" json:"phone_search"`
	RealNameSearch int    `gorm:"default:0" json:"real_name_search"`
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

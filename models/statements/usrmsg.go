package statements

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Usrmsg struct {
<<<<<<< HEAD
	FromUser uint `gorm:"not null"`
	ToUser   uint `gorm:"not null"`
=======
	FromUser uint `gorm:not null`
	ToUser   uint `gorm:not null`
>>>>>>> 922030c154cc558ac41520ba99ca5b16da30c9b1
	Url      string
	Song     string
	Message  string
	gorm.Model
	IsSend int
}

func UsrmsgInit() {
	db := setting.MysqlConn()

	if !db.HasTable(&Usrmsg{}) {
		if err := db.CreateTable(&Usrmsg{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Usrmsg has been created")
	} else {
		db.AutoMigrate(&Usrmsg{})
		fmt.Println("Table Usrmsg has existed")
	}
}

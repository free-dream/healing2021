package setting

import (
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	dbName := tools.GetConfig("mysql", "dbName")
	user := tools.GetConfig("mysql", "user")
	password := tools.GetConfig("mysql", "password")
	port := tools.GetConfig("mysql", "port")
	dbInfo := user + ":" + password + "@tcp(" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"
	//connect
	var err error
	DB, err = gorm.Open("mysql", dbInfo)
	DB.SingularTable(true)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().SetMaxIdleConns(20)
	if err != nil {
		panic(err)
	}
}

func MysqlConnTest() {
	dbName := tools.GetConfig("mysql", "dbName")
	user := tools.GetConfig("mysql", "user")
	password := tools.GetConfig("mysql", "password")
	port := tools.GetConfig("mysql", "port")
	dbInfo := user + ":" + password + "@tcp(" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	//connect
	db, err := gorm.Open("mysql", dbInfo)
	db.SingularTable(true)
	db.Close()
	if err != nil {
		panic(err)
	}
}

func MysqlConn() *gorm.DB {
	return DB
}

func TimeSetting(string) {
	db := MysqlConn()
	//插入时自动更新created_at字段
	db.Exec("ALTER TABLE " + " MODIFY created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL")
	//更新时自动更新updated_at字段
	db.Exec("ALTER TABLE " + "MODIFY updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;")

}

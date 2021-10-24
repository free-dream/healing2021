package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//根据日期查询对应热榜
func GetCoversByDate(date string) ([]statements.Cover, error) {
	mysqlDb := setting.MysqlConn()

	data := make([]statements.Cover, 10)
	err := mysqlDb.Where("CreatedAt LIKE ?", (date + "%")).Order("Likes desc").Limit(10).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

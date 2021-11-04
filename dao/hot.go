package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//根据日期查询对应热榜
func GetCoversByDate(date string) ([]statements.Cover, error) {
	mysqlDb := setting.MysqlConn()

	var data []statements.Cover

	err := mysqlDb.Where("CreatedAt LIKE ?", (date + "%")).Order("IsLiked desc").Limit(10).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

//获取全时间获赞最高项
func GetCoversByLikes() ([]statements.Cover, error) {
	mysqlDb := setting.MysqlConn()

	var datas []statements.Cover

	err := mysqlDb.Order("IsLiked desc").Limit(10).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}

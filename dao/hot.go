package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//根据日期获取对应热榜
func GetCoversByDate(date string) ([]tables.Cover, []resp.CoverRank, error) {
	mysqlDb := setting.MysqlConn()

	//组合获得日期模糊匹配
	temp := "%" + date + "%"
	var data []tables.Cover
	var likes []resp.CoverRank

	//子查询准备
	subquery := mysqlDb.Table("cover").Select("id").Where("created_at like ?", temp)
	if err1 := subquery.Error; err1 != nil {
		return nil, nil, err1
	}

	//主查询
	err := mysqlDb.Order("likes desc").
		Table("praise").
		Select("cover_id, count(cover_id) as likes").
		Group("cover_id").
		Where("cover_id in (?)", subquery).
		Limit(10).Error

	if err != nil {
		return nil, nil, err
	}
	return data, likes, nil
}

//获取全时间获赞最高项
func GetCoversByLikes() ([]tables.Cover, []resp.CoverRank, error) {
	mysqlDb := setting.MysqlConn()
	var likes []resp.CoverRank
	covers := make([]tables.Cover, 0)

	err := mysqlDb.Order("likes desc").
		Table("praise").
		Select("cover_id, count(cover_id) as likes").
		Group("cover_id").
		Limit(10).
		Scan(&likes).Error

	if err != nil {
		return nil, nil, err
	}
	for _, item := range likes {
		var temp tables.Cover
		err = mysqlDb.Where("id = ?", item.CoverId).First(&temp).Error
		if err != nil {
			return nil, nil, err
		}
		covers = append(covers, temp)
	}
	return covers, likes, nil
}

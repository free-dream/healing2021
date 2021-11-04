package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//根据日期获取对应热榜
func GetCoversByDate(date string) ([]tables.Cover, error) {
	mysqlDb := setting.MysqlConn()

	var data []tables.Cover
	var likes []resp.CoverRank

	err := mysqlDb.Order("likes desc").
		Table("praise").
		Select("cover_id, count(cover_id) as likes").
		Group("cover_id").
		Scan(&likes).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

//获取全时间获赞最高项
func GetCoversByLikes() ([]tables.Cover, []resp.CoverRank, error) {
	mysqlDb := setting.MysqlConn()
	// 用cover表做了一个测试
	// var test []Test
	// err1 := mysqlDb.Order("test desc").Table("cover").Select("nickname, count(nickname) as test").Group("nickname").Scan(&test).Error
	// if err1 != nil {
	// 	panic(err1)
	// }
	// fmt.Println(test)
	// return nil, nil
	// //
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

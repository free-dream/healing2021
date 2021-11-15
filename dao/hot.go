package dao

import (
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

//根据日期获取对应热榜
func GetCoversByDate(date string) ([]resp.HotResp, error) {
	mysqlDb := setting.MysqlConn()

	//组合获得日期模糊匹配
	temp := "%" + date + "%"

	var likes []resp.HotResp

	//子查询准备
	subquery := mysqlDb.Select("id").Where("created_at like ?", temp).Table("cover").SubQuery()
	//主查询
	err := mysqlDb.Order("likes desc").
		Table("praise").
		Select("cover_id, count(cover_id) as likes, cover.Avatar as avatar, cover.Nickname as nickname, cover.song_name as songname,cover.module as module,cover.selection_id as selection_id,cover.created_at as created_at,cover.classic_id as classic_id").
		Where("cover_id <> ? AND is_liked = ? AND cover.MODULE <> ?", 0, 1, 2).
		Joins("left join cover on cover.id = cover_id").
		Group("cover_id").
		Where("cover_id in (?)", subquery).
		Limit(10).
		Find(&likes).Error

	if err != nil {
		return nil, err
	}

	return likes, nil
}

//获取全时间获赞最高项
func GetCoversByLikes() ([]resp.HotResp, error) {
	mysqlDb := setting.MysqlConn()
	var likes []resp.HotResp

	err := mysqlDb.Order("likes desc").
		Table("praise").
		Select("cover_id, count(cover_id) as likes, cover.Avatar as avatar, cover.Nickname as nickname, cover.song_name as songname,cover.module as module,cover.selection_id as selection_id,cover.created_at as created_at,cover.classic_id as classic_id").
		Where("cover_id <> ? AND is_liked = ? AND cover.module <> ?", 0, 1, 2).
		Joins("left join cover on cover.id = cover_id").
		Group("cover_id").
		Limit(10).
		Scan(&likes).Error

	if err != nil {
		return nil, err
	}

	return likes, nil
}

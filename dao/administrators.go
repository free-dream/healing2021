package dao

import "git.100steps.top/100steps/healing2021_be/pkg/setting"

func Authentication(openid string) bool {
	redisCli := setting.RedisConn()
	return redisCli.SIsMember("healing2021:administrator", openid).Val()
}

func DeleteContent(id int) error {
	db := setting.MysqlConn()
	err := db.Table("moment_comment").Select("is_deleted").Where("id=?", id).Update("is_deleted", 1).Error
	return err
}

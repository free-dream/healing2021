package dao

import "git.100steps.top/100steps/healing2021_be/pkg/setting"

func Authentication(nickname string) bool {
	return setting.RedisClient.SIsMember("healing2021:administrator", nickname).Val()
}

func DeleteContent(id int) error {
	err := setting.DB.Table("moment_comment").Select("is_deleted").Where("id", id).Update(1).Error
	return err
}

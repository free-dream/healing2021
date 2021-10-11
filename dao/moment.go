package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

/**
 * @Description 获取所有的动态【没有加行锁的必要】
 * @Param 获取方式method string, 关键字keyword string
 * @return 含有所有动态信息的切片AllMoment(按时间排序)，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetAllMoment(method string, keyword string) ([]statements.Moment, bool) {
	var AllMoment []statements.Moment
	MysqlDB := setting.MysqlConn()
	if err := MysqlDB.Find(&AllMoment).Error; err != nil {
		return AllMoment, false
	}
	return AllMoment, true
}

/**
* @description: $description$
* @param: $params$
* @return: $returns$
*/


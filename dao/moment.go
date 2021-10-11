package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

/**【未完工】
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

/**【未完工】
* @description: 用歌曲 Id 找歌曲名字
* @param: 歌曲Id
* @return: 歌曲名字
*/
func GetSongNameById(SongId int)  string{
	var SongName string
	return SongName
}

/**【未完工】
* @description: 通过动态的 Id 来统计动态被点赞数
* @param: 动态 Id
* @return: 点赞次数
 */
func CountLaudsById(MomentId int) int {
	var Tot int
	return Tot
}

/**【未完工】
* @description: 通过动态的 Id 来判断当前用户是否点过赞
* @param: 动态 Id
* @return: 1 表示已经点过赞
 */
func HaveLauded(MomentId int) int {
	var ok int
	return ok
}

/**【未完工】
* @description: 通过动态的 Id 来统计评论总数
* @param: 动态 Id
* @return: 评论总数
 */
func CountCommentsById(MomentId int) int {
	var Tot int
	return Tot
}


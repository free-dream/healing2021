package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

/**【未完工】
 * @Description 获取指定的一页(十条)动态【没有加行锁的必要】
 * @Param 获取方式method string, 关键字keyword string，页码page int
 * @return 含有所有动态信息的切片AllMoment(按时间排序)，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetMomentPage(Method string, Keyword string, Page int) ([]statements.Moment, bool) {
	MysqlDB := setting.MysqlConn()
	var AllMoment []statements.Moment
	if Method == "new" {
		if err := MysqlDB.Order("created_at DESC").Offset(Page * 10).Limit(10).Find(&AllMoment).Error; err != nil {
			return AllMoment, false
		}
	} else if Method == "recommend" {
		// 得写聚合分组查询
		//if err := MysqlDB.Order("created_at DESC").Offset(Page*10).Limit(10).Find(&AllMoment).Error; err != nil {
		//	return AllMoment, false
		//}
	} else {
		// 得写模糊查找
		//if err := MysqlDB.Order("created_at DESC").Offset(Page*10).Limit(10).Find(&AllMoment).Error; err != nil {
		//	return AllMoment, false
		//}
	}

	return AllMoment, true
}

/**【未测试】
* @description: 创建新动态
* @param: Moment 结构体
* @return: 操作是否成功 ok
 */
func CreateMoment(Moment statements.Moment) bool {
	MysqlDB := setting.MysqlConn()
	if err := MysqlDB.Create(&Moment); err != nil {
		return false
	}
	return true
}

/**【未测试】
* @description: 用动态 Id 找动态的记录
* @param: 动态Id
* @return: Moment 结构体
 */
func GetMomentById(MomentId int) (statements.Moment, bool) {
	MysqlDB := setting.MysqlConn()
	var Moment statements.Moment

	if err := MysqlDB.Where("id=?", MomentId).First(&Moment); err != nil {
		return Moment, false
	}

	return Moment, true
}

/**【未完工】
* @description: 通过动态的 Id 来统计动态被点赞数
* @param: 动态 Id
* @return: 点赞次数
 */
func CountMLaudsById(MomentId int) int {
	var Tot int
	return Tot
}

/**【未完工】
* @description: 通过动态的 Id 来判断当前用户是否点过赞
* @param: 动态 Id
* @return: 1 表示已经点过赞
 */
func HaveMLauded(UserId int, MomentId int) int {
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

/**【未测试】
* @description: 创建新评论
* @param: Comment 结构体
* @return: 操作是否成功 ok
 */
func CreateComment(Comment statements.MomentComment) bool {
	MysqlDB := setting.MysqlConn()
	if err := MysqlDB.Create(&Comment); err != nil {
		return false
	}
	return true
}

/**【未实现】
* @description: 拉取一个动态下的评论列表
* @param: 动态的Id int
* @return: 操作是否成功 ok
 */
func GetCommentsByMomentId(MomentId int) ([]statements.MomentComment, bool) {
	// MysqlDB := setting.MysqlConn()
	var CommentList []statements.MomentComment
	// 聚类函数
	return CommentList, true
}

/**【未完工】
* @description: 通过评论的 Id 来统计动态被点赞数
* @param: 评论 Id
* @return: 点赞次数
 */
func CountCLaudsById(CommentId int) int {
	var Tot int
	return Tot
}

/**【未完工】
* @description: 通过评论的 Id 来判断当前用户是否点过赞
* @param: 评论 Id
* @return: 1 表示已经点过赞
 */
func HaveCLauded(UserId int, CommentId int) int {
	var ok int
	return ok
}

/**【未完工】
* @description: 通过评论的 Id 点赞
* @param: 评论 Id
* @return: ok 表示点过赞操作是否成功
 */
func CLaudedById(UserId int, CommentId int) bool {
	return true
}

/**【未完工】
* @description: 通过动态的 Id 点赞
* @param: 动态 Id
* @return: ok 表示点过赞操作是否成功
 */
func MLaudedById(UserId int, CommentId int) bool {
	return true
}

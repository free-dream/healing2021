package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

/**【未测试】
 * @Description 获取指定的一页(十条)动态【查询没有加行锁的必要】
 * @Param 获取方式method string, 关键字keyword string，页码page int
 * @return 含有所有动态信息的切片AllMoment(按时间排序)，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetMomentPage(Method string, Keyword string, Page int) ([]statements.Moment, bool) {
	MysqlDB := setting.MysqlConn()
	var AllMoment []statements.Moment
	if Method == "new" {
		// 按时间排序
		if err := MysqlDB.Order("created_at DESC").Offset(Page * 10).Limit(10).Find(&AllMoment).Error; err != nil {
			return AllMoment, false
		}
	} else if Method == "recommend" {
		// 按点赞排序
		if err := MysqlDB.Order("like_num DESC").Offset(Page*10).Limit(10).Find(&AllMoment).Error; err != nil {
			return AllMoment, false
		}
	} else {
		// 模糊查找
		if err := MysqlDB.Where("name LIKE ?", Keyword).Order("created_at DESC").Offset(Page*10).Limit(10).Find(&AllMoment).Error; err != nil {
			return AllMoment, false
		}
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

/**【未测试】
* @description: 通过动态的 Id 来统计动态被点赞数
* @param: 动态 Id
* @return: 点赞次数
 */
func CountMLaudsById(MomentId int) int {
	Moment, ok := GetMomentById(MomentId)
	if !ok{
		return -1
	}

	return Moment.LikeNum
}

/**【未测试】
* @description: 通过动态的 Id 来判断当前用户是否点过赞
* @param: 动态 Id
* @return: 1 表示已经点过赞, -1 表示发生异常情况
 */
func HaveMLauded(UserId int, MomentId int) int {
	MysqlDB := setting.MysqlConn()

	err := MysqlDB.Where("user_id=? and moment_id=?", UserId, MomentId).First(&statements.Praise{}).Error
	if gorm.IsRecordNotFoundError(err){
		return 0
	} else if err != nil {
		return -1
	}
	return 1
}

/**【未测试】
* @description: 通过动态的 Id 来统计评论总数
* @param: 动态 Id
* @return: 评论总数, 异常时返回-1
 */
func CountCommentsById(MomentId int) int {
	MysqlDB := setting.MysqlConn()
	// 用聚类函数
	var Tot int
	err := MysqlDB.Where("moment_id", MomentId).Count(&Tot).Error
	if err != nil {
		return  -1
	}
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

/**【未测试】
* @description: 拉取一个动态下的评论列表
* @param: 动态的Id int
* @return: 操作是否成功 ok
 */
func GetCommentsByMomentId(MomentId int) ([]statements.MomentComment, bool) {
	MysqlDB := setting.MysqlConn()
	var CommentList []statements.MomentComment
	err := MysqlDB.Where("moment_id", MomentId).Find(&CommentList).Error
	if err != nil {
		return CommentList, false
	}

	return CommentList, true
}

/**【未测试】
* @description: 用评论 Id 找评论的记录
* @param: 评论 Id
* @return: Comment 结构体
 */
func GetCommentIdById(CommentId int) (statements.MomentComment, bool) {
	MysqlDB := setting.MysqlConn()
	var Comment statements.MomentComment

	if err := MysqlDB.Where("id=?", CommentId).First(&Comment); err != nil {
		return Comment, false
	}

	return Comment, true
}

/**【未测试】
* @description: 通过评论的 Id 来统计动态被点赞数
* @param: 评论 Id
* @return: 点赞次数
 */
func CountCLaudsById(CommentId int) int {
	Comment, ok := GetCommentIdById(CommentId)
	if !ok{
		return -1
	}

	return Comment.LikeNum
}

/**【未测试】
* @description: 通过评论的 Id 来判断当前用户是否点过赞
* @param: 评论 Id
* @return: 1 表示已经点过赞
 */
func HaveCLauded(UserId int, CommentId int) int {
	MysqlDB := setting.MysqlConn()

	err := MysqlDB.Where("user_id=? and moment_id=?", UserId, CommentId).First(&statements.Praise{}).Error
	if gorm.IsRecordNotFoundError(err){
		return 0
	} else if err != nil {
		return -1
	}
	return 1
}

/**【未测试】
* @description: 通过评论的 Id 点赞
* @param: 评论 Id
* @return: ok 表示点过赞操作是否成功
 */
func CLaudedById(UserId int, CommentId int) bool {
	MysqlDB := setting.MysqlConn()
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false
	}

	Comment := statements.MomentComment{}

	// 锁住指定 id 的 Comment 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&Comment, CommentId).Error; err != nil {
		tx.Rollback()
		return false
	}

	// 更新
	if err := MysqlDB.Model(statements.MomentComment{}).Where("id = ? ", CommentId).Update("like_num", gorm.Expr("like_num+ ?", 1)).Error; err != nil {
		return false
	}

	// 插入一条点赞记录
	Praise := statements.Praise{
		UserId : UserId,
		MomentCommentId: CommentId,
	}
	if err := MysqlDB.Create(&Praise); err != nil {
		return false
	}

	// commit事务，释放锁
	if err := tx.Commit().Error; err != nil {
		return false
	}

	return true
}

/**【未测试】
* @description: 通过动态的 Id 点赞
* @param: 动态 Id
* @return: ok 表示点过赞操作是否成功
 */
func MLaudedById(UserId int, MomentId int) bool {
	MysqlDB := setting.MysqlConn()
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false
	}

	Moment := statements.Moment{}

	// 锁住指定 id 的 Comment 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&Moment, MomentId).Error; err != nil {
		tx.Rollback()
		return false
	}

	// 更新
	if err := MysqlDB.Model(statements.Moment{}).Where("id = ? ", MomentId).Update("like_num", gorm.Expr("like_num+ ?", 1)).Error; err != nil {
		return false
	}

	// 插入一条点赞记录
	Praise := statements.Praise{
		UserId : UserId,
		MomentId: MomentId,
	}
	if err := MysqlDB.Create(&Praise); err != nil {
		return false
	}

	// commit事务，释放锁
	if err := tx.Commit().Error; err != nil {
		return false
	}

	return true
}

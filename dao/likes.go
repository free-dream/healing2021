package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

//基于直接点赞更新mysql，加锁
func UpdateLikesByID(user int, target int, likes int, kind string) error {
	mysqlDb := db.MysqlConn()
	var (
		err  error
		like tables.Praise
	)
	//加行锁
	lock := mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			lock.Rollback()
		}
	}()
	if err = lock.Error; err != nil {
		return err
	}

	// TODO:检查原本是否已经点赞、取消点赞

	//更新数据库
	if kind == "cover" {
		err = lock.Model(&like).Where("cover_id = ? AND user_id = ?", target, user).UpdateColumn("is_liked", gorm.Expr("is_liked + ?", likes)).Error
	} else if kind == "moment" {
		err = lock.Model(&like).Where("moment_id = ? AND user_id = ?", target, user).UpdateColumn("is_liked", gorm.Expr("is_liked + ?", likes)).Error
	} else if kind == "momentcomment" {
		err = lock.Model(&like).Where("moment_comment_id = ? AND user_id = ?", target, user).UpdateColumn("is_liked", gorm.Expr("is_liked + ?", likes)).Error
	} else {
		panic("wrong type")
	}
	//更新失败就回滚
	if err != nil {
		lock.Rollback()
		return err
	}
	//提交事务
	if err = lock.Commit().Error; err != nil {
		return err
	}
	return nil
}

/*【planB】
点赞方案重构
数据一致性不好处理，干脆放弃 redis
取消 moment表、commont表、cover表 中 likeNum 这一字段

进行点赞  直接在点赞表中额外插入一条 is_liked=1 的记录
取消点赞	直接在点赞表中额外插入一条 is_liked=1 的记录 （或者找到之前的记录进行删除）
获取点赞	直接使用聚类查找 praise 表
*/

//备选方案，基于redis的更新，直接在goroutine加锁
func RUpdateLikesByID(user int, target int, likes int, kind string) bool {
	mysqlDb := db.MysqlConn()
	var like tables.Praise
	var err error
	if kind == "cover" {
		err = mysqlDb.Model(&like).Where("CoverId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes)).Error
	} else if kind == "moment" {
		err = mysqlDb.Model(&like).Where("MomentId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes)).Error
	} else if kind == "momentcomment" {
		err = mysqlDb.Model(&like).Where("MomentCommentId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes)).Error
	} else {
		err = nil
	}
	if err == nil {
		return true
	} else {
		panic("something wrong when get" + kind + " like record")
	}
}

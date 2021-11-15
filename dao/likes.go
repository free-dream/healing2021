package dao

import (
	"log"

	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/jinzhu/gorm"
)

//包装一下给controller用
func PackageCheck(user int, kind string, target int) (bool, error) {
	db := setting.MysqlConn()
	boolean, err := CheckMysql(db, user, target, kind, 1, true)
	return boolean, err
}

//跳转mysql检查
func CheckMysql(lock *gorm.DB, user int, target int, kind string, likes int, choose bool) (bool, error) {

	//redis检查
	redisCheck, err1 := sandwich.Check(target, kind, user)
	// db.RedisClient.Pipeline()
	if err1 != nil {
		log.Printf(err1.Error())
	} else {
		return redisCheck, nil
	}

	//redis爆炸的保险，下面基本上不会触发，只是以防万一
	var (
		err  error
		like tables.Praise
	)
	switch kind {
	case "cover":
		err = lock.Where("cover_id = ? AND user_id = ?", target, user).Find(&like).Error
	case "moment":
		err = lock.Where("moment_id = ? AND user_id = ?", target, user).Find(&like).Error
	case "momentcomment":
		err = lock.Where("moment_comment_id = ? AND user_id = ?", target, user).Find(&like).Error
	default:
		panic("wrong type") //基本上不可能抵达,除非有意
	}
	//仅判断是否存在
	if choose {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else if err == nil {
			return true, nil
		} else {
			return false, err
		}
	}
	//点赞判断
	if gorm.IsRecordNotFoundError(err) {
		if likes > 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else if err != nil {
		if likes < 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, err
	}
}

//基于直接点赞更新mysql，加锁
func UpdateLikesByID(user int, target int, likes int, kind string) error {

	//直接走redis
	var (
		err  error
		like tables.Praise
	)

	mysqlDb := db.MysqlConn()
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
	// 检查原本是否已经点赞|取消点赞,要求redis缓存匹配
	// 用户第一次进行点赞时没有记录能进行 update
	// 若为取消点赞，应检查是否存有当前用户id,若为点赞，则反之

	if likes == 1 { //创建点赞表,更新coverLikes字段
		switch kind {
		case "cover":
			like = tables.Praise{
				UserId:  user,
				CoverId: target,
				IsLiked: likes,
			}
			err = lock.Create(&like).Error
		case "moment":
			like = tables.Praise{
				UserId:   user,
				MomentId: target,
				IsLiked:  likes,
			}
			err = lock.Create(&like).Error
		case "momentcomment":
			like = tables.Praise{
				UserId:          user,
				MomentCommentId: target,
				IsLiked:         likes,
			}
			err = lock.Create(&like).Error
		default:
			panic("wrong type") //基本上不可能抵达,除非有意设计
		}
		//错误处理
		if err != nil {
			lock.Rollback()
			return err
		}
	} else if likes == -1 { //更新点赞表
		switch kind {
		case "cover":
			err = lock.Where("cover_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		case "moment":
			err = lock.Where("moment_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		case "momentcomment":
			err = lock.Where("moment_comment_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		default:
			panic("wrong type") //基本上不可能抵达,除非有意
		}
	}
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

func markMomentInPraise(momentId int) error {
	mysqlDb := db.MysqlConn()
	like := tables.Praise{
		UserId:   0,
		MomentId: momentId,
		IsLiked:  0,
	}
	err := mysqlDb.Create(&like).Error
	return err
}

func ViolenceGetLikeheck(id int, resp CoverDetails, ch chan CoverDetails) {
	boolean, err := PackageCheckMysql(id, "cover", resp.ID)
	if err != nil {
		log.Printf(err.Error())
		resp.Check = 0
	} else if boolean {
		resp.Check = 1
	} else {
		resp.Check = 0
	}
	ch <- resp
}

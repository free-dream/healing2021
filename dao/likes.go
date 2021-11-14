package dao

import (
	"log"

	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/jinzhu/gorm"
)

//自定义错误
type LikesExistError struct{}
type ZeroUserIdError struct{}

func (l *LikesExistError) Error() string {
	return "点赞/取消点赞失败,redis记录不匹配"
}

func (l *ZeroUserIdError) Error() string {
	return "当前用户id为0,无法点赞"
}

//重复点赞判断函数
func IsLikesExistError(err error) bool {
	if err.Error() == "点赞/取消点赞失败,redis记录不匹配" {
		return true
	} else {
		return false
	}
}

//0用户id判断参数
func IsZeroUserIdError(err error) bool {
	if err.Error() == "当前用户id为0,无法点赞" {
		return true
	} else {
		return false
	}
}

//包装一下给controller用
func PackageCheckMysql(user int, kind string, target int) (bool, error) {
	db := setting.MysqlConn()
	boolean, err := CheckMysql(db, user, target, kind, 1, true)
	return boolean, err
}

//跳转mysql检查
func CheckMysql(lock *gorm.DB, user int, target int, kind string, likes int, choose bool) (bool, error) {
	//redis检查
	redisCheck, err1 := sandwich.Check(target, kind, user)
	if err1 != nil {
		log.Printf(err1.Error())
	} else {
		return redisCheck, nil
	}
	//
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
	if user == 0 {
		var err error = &ZeroUserIdError{}
		return err
	}
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
	// 检查原本是否已经点赞|取消点赞,要求redis缓存匹配
	// 用户第一次进行点赞时没有记录能进行 update
	// 若为取消点赞，应检查是否存有当前用户id,若为点赞，则反之

	check, _ := sandwich.Check(target, kind, user)
	if likes == -1 && check {
		err = sandwich.CancelLike(target, kind, user)
		if err != nil {
			return err
		}
	} else if likes == 1 && !check {
		err = sandwich.AddLike(target, kind, user)
		if err != nil {
			return err
		}
	} else {
		var err1 error = &LikesExistError{}
		return err1
	}
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

func ViolenceGetLikeNum(id int, ch chan int) int {
	db := setting.MysqlConn()
	num := 0
	db.Select("praise").Where("cover_id=? and is_liked=?", id, 1).Count(&num)
	ch <- id
	return num
}

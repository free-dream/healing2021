package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/sandwich"
)

//自定义错误
type LikesExistError struct{}

func (l *LikesExistError) Error() string {
	return "点赞/取消点赞失败,redis记录不匹配"
}

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

	// 检查原本是否已经点赞|取消点赞,要求redis缓存匹配
	// 用户第一次进行点赞时没有记录能进行 update
	// 若为取消点赞，应检查是否存有当前用户id,若为点赞，则反之
	check := sandwich.Check(target, kind, user)
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

	//分支结构估计跑起来不快,不过先这样
	if target == 1 { //创建点赞表
		if kind == "cover" {
			like = tables.Praise{
				CoverId: target,
				IsLiked: likes,
			}
			err = lock.Create(&like).Error
		} else if kind == "moment" {
			like = tables.Praise{
				MomentId: target,
				IsLiked:  likes,
			}
			err = lock.Create(&like).Error
		} else if kind == "momentcomment" {
			like = tables.Praise{
				MomentCommentId: target,
				IsLiked:         likes,
			}
			err = lock.Create(&like).Error
		} else {
			panic("wrong type") //基本上不可能抵达,除非有意设计
		}
		//错误处理
		if err != nil {
			lock.Rollback()
			return err
		}
	} else if target == -1 { //删除点赞表
		if kind == "cover" {
			err = lock.Model(&like).Where("cover_id = ? AND user_id = ?", target, user).Delete(&like).Error
		} else if kind == "moment" {
			err = lock.Model(&like).Where("moment_id = ? AND user_id = ?", target, user).Delete(&like).Error
		} else if kind == "momentcomment" {
			err = lock.Model(&like).Where("moment_comment_id = ? AND user_id = ?", target, user).Delete(&like).Error
		} else {
			panic("wrong type") //基本上不可能抵达,除非有意
		}
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

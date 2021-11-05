package dao

import (
	"fmt"

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
	//
	fmt.Println("加行锁")
	//
	// 检查原本是否已经点赞|取消点赞,要求redis缓存匹配
	// 用户第一次进行点赞时没有记录能进行 update
	// 若为取消点赞，应检查是否存有当前用户id,若为点赞，则反之

	check := sandwich.Check(target, kind, user)
	//
	fmt.Println(check)
	//
	if likes == -1 && check {
		err = sandwich.CancelLike(target, kind, user)
		if err != nil {
			//
			fmt.Println("一重身")
			//
			return err
		}
	} else if likes == 1 && !check {
		err = sandwich.AddLike(target, kind, user)
		if err != nil {
			//
			fmt.Println("二重身")
			//
			return err
		}
	} else {
		//
		fmt.Println("三重身")
		//
		var err1 error = &LikesExistError{}
		return err1
	}
	//
	fmt.Println("过检查")
	//
	if likes == 1 { //创建点赞表,更新coverLikes字段
		switch kind {
		case "cover":
			like = tables.Praise{
				UserId:  user,
				CoverId: target,
				IsLiked: likes,
			}
			//
			fmt.Println("cover建表")
			//
			err = lock.Create(&like).Error
		case "moment":
			like = tables.Praise{
				UserId:   user,
				MomentId: target,
				IsLiked:  likes,
			}
			//
			fmt.Println("moment建表")
			//
			err = lock.Create(&like).Error
		case "momentcomment":
			like = tables.Praise{
				UserId:          user,
				MomentCommentId: target,
				IsLiked:         likes,
			}
			//
			fmt.Println("comment建表")
			//
			err = lock.Create(&like).Error
		default:
			//
			fmt.Println("panic wrong type")
			//
			panic("wrong type") //基本上不可能抵达,除非有意设计
		}
		//错误处理
		if err != nil {
			lock.Rollback()
			return err
		}
	} else if likes == -1 { //删除点赞表
		switch kind {
		case "cover":
			//
			fmt.Println("删除cover")
			//
			err = lock.Where("cover_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		case "moment":
			//
			fmt.Println("删除moment")
			//
			err = lock.Where("moment_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		case "momentcomment":
			//
			fmt.Println("删除momentcomment")
			//
			err = lock.Where("moment_comment_id = ? AND user_id = ?", target, user).Unscoped().Delete(&like).Error
		default:
			panic("wrong type") //基本上不可能抵达,除非有意
		}
	}
	if err != nil {
		return err
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

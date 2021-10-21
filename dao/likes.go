package dao

import (
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

func UpdateLikesByID(user int, target int, likes int, kind string) bool {
	mysqlDb := db.MysqlConn()
	var like tables.Praise
	var err interface{}
	if kind == "cover" {
		err = mysqlDb.Model(&like).Where("CoverId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes))
	} else if kind == "moment" {
		err = mysqlDb.Model(&like).Where("MomentId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes))
	} else if kind == "momentcomment" {
		err = mysqlDb.Model(&like).Where("MomentCommentId = ? AND UserId = ?", target, user).UpdateColumn("IsLiked", gorm.Expr("IsLiked + ?", likes))
	} else {
		err = nil
	}
	if err == nil {
		return true
	} else {
		panic("something wrong when get" + kind + " like record")
	}
}

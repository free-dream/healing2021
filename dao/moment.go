package dao

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
	"time"
)

var momentPageCache []int

// 定时更新推荐动态
func UpdateMomentPage()  {
	db := setting.MysqlConn()
	for {
		//执行代码
		rows, err := db.Model(&statements.Praise{}).Select("moment_id").Where("moment_id<>?", 0).Group("moment_id").Order("count(is_liked) DESC").Rows()
		if err != nil {
			fmt.Println("动态推荐更新失败")
		}
		defer rows.Close()

		var momentRecord ForPraiseMRecord
		var tmpCache []int
		for rows.Next() {
			// 全扫描进结构体
			err := db.ScanRows(rows, &momentRecord)
			if err != nil {
				break
			}
			tmpCache = append(tmpCache, momentRecord.MomentId)
		}

		t := time.NewTimer(time.Second * 300)
		<-t.C
	}
}

// 获取指定的一页(十条)动态
// 拆分动态获取 one -> three
type ForPraiseMRecord struct {
	MomentId int `gorm:"moment_id"`
}
func GetMomentPageNew(page int) ([]statements.Moment, error) {
	db := setting.MysqlConn()
	var momentPage []statements.Moment
	err := db.Order("created_at DESC").Offset(page * 10).Limit(10).Find(&momentPage).Error
	return momentPage, err
}
func GetMomentPageRecommend(page int) ([]statements.Moment, error) {
	db := setting.MysqlConn()
	var momentPage []statements.Moment

	left := page*10
	right := left+10

	if left>len(momentPageCache) {
		return momentPage, nil
	} else if right>len(momentPageCache) {
		right = len(momentPageCache)
	}

	for _, record := range momentPageCache[left:right] {
		moment := statements.Moment{}
		err := db.Where("id=?", record).First(&moment).Error
		if err != nil {
			return momentPage, err
		}
		momentPage = append(momentPage, moment)
	}
	return momentPage, nil
}
func GetMomentPageSearch(page int, keyWords string) ([]statements.Moment, error) {
	db := setting.MysqlConn()
	var momentPage []statements.Moment

	Fuzzy := "%" + keyWords + "%"
	err := db.Where("content LIKE ? or song_name LIKE ? or state LIKE ?", Fuzzy, Fuzzy, Fuzzy).Order("created_at DESC").Offset(page * 10).Limit(10).Find(&momentPage).Error
	return momentPage, err
}

// 创建新动态
func CreateMoment(Moment statements.Moment) error {
	MysqlDB := setting.MysqlConn()
	err := MysqlDB.Create(&Moment).Error
	if  err != nil {
		return err
	}
	err = markMomentInPraise(int(Moment.ID))
	if err != nil {
		return err
	}
	return nil
}

// 用动态 Id 找动态的记录
func GetMomentById(MomentId int) (statements.Moment, bool) {
	MysqlDB := setting.MysqlConn()
	Moment := statements.Moment{}
	if err := MysqlDB.Where("id=?", MomentId).First(&Moment).Error; err != nil {
		return Moment, false
	}
	return Moment, true
}

// 通过动态的 Id 来统计动态被点赞数
func CountMLaudsById(MomentId int) int {
	MysqlDB := setting.MysqlConn()
	var Lauds int
	err := MysqlDB.Model(&statements.Praise{}).Where("is_liked=? and moment_id=?", 1, MomentId).Count(&Lauds).Error
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return Lauds
}

// 通过动态的 Id 来判断当前用户是否点过赞
func HaveMLauded(UserId int, MomentId int) int {
	MysqlDB := setting.MysqlConn()
	err := MysqlDB.Where("user_id=? and moment_id=? and is_liked=?", UserId, MomentId, 1).First(&statements.Praise{}).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0
	} else if err != nil {
		return -1
	}
	return 1
}

// 通过动态的 Id 来统计评论总数
func CountCommentsById(MomentId int) int {
	MysqlDB := setting.MysqlConn()
	var Tot = 0

	// 用聚类函数来操作
	err := MysqlDB.Model(&statements.MomentComment{}).Where("moment_id=? and is_deleted=?", MomentId, 0).Count(&Tot).Error
	if err != nil {
		return -1
	}
	return Tot
}

// 创建新评论,返回创建好的评论的 id
func CreateComment(Comment statements.MomentComment) (int, error) {
	db := setting.MysqlConn()
	err := db.Create(&Comment).Error
	return int(Comment.ID), err
}

// 拉取一个动态下的评论列表
func GetCommentsByMomentId(MomentId int) ([]statements.MomentComment, bool) {
	MysqlDB := setting.MysqlConn()
	var CommentList []statements.MomentComment

	err := MysqlDB.Where("moment_id=? and is_deleted=?", MomentId, 0).Find(&CommentList).Error
	if err != nil {
		return CommentList, false
	}
	return CommentList, true
}

// 用评论 Id 找评论
func GetCommentIdById(CommentId int) (statements.MomentComment, bool) {
	MysqlDB := setting.MysqlConn()
	var Comment statements.MomentComment

	if err := MysqlDB.Where("id=?", CommentId).First(&Comment).Error; err != nil {
		return Comment, false
	}
	return Comment, true
}

// 通过评论的 Id 来统计评论被点赞数
func CountCLaudsById(CommentId int) int {
	MysqlDB := setting.MysqlConn()
	var Lauds int
	err := MysqlDB.Model(&statements.Praise{}).Where("is_liked=? and moment_comment_id=?", 1, CommentId).Count(&Lauds).Error
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return Lauds
}

// 通过评论的 Id 来判断当前用户是否点过赞
func HaveCLauded(UserId int, CommentId int) int {
	MysqlDB := setting.MysqlConn()

	err := MysqlDB.Where("user_id=? and moment_comment_id=? and is_liked=?", UserId, CommentId, 1).First(&statements.Praise{}).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0
	} else if err != nil {
		return -1
	}
	return 1
}

// 通过动态id获得动态发送者userId
type MomentSenderId struct {
	UserId int `gorm:"user_id"`
}

func GetMomentSenderId(MomentId int) (int, error) {
	momentSenderId := MomentSenderId{}
	db := setting.MysqlConn()
	err := db.Model(&statements.Moment{}).Select("user_id").Where("id=?", MomentId).Scan(&momentSenderId).Error
	return momentSenderId.UserId, err
}

// 通过评论id获得动态发送者userId
type CommentSenderId struct {
	UserId int `gorm:"user_id"`
}

func GetCommentSenderId(CommentId int) (int, error) {
	commentSenderId := MomentSenderId{}
	db := setting.MysqlConn()
	err := db.Model(&statements.MomentComment{}).Select("user_id").Where("id=?", CommentId).Scan(&commentSenderId).Error
	return commentSenderId.UserId, err
}

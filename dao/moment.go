package dao

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
)

// 获取指定的一页(十条)动态
type ForPraiseMRecord struct {
	MomentId int `gorm:"moment_id"`
	Tot      int `gorm:"count(is_liked)"`
	//Tot int `gorm:"tot"`
}

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
		// 先查点赞表找到对应的动态
		var MomentRecords []ForPraiseMRecord
		var MomentRecord ForPraiseMRecord
		rows, err := MysqlDB.Model(&statements.Praise{}).Select("moment_id, count(is_liked)").Where("moment_id<>?", 0).Group("moment_id").Order("count(is_liked) DESC").Offset(Page * 10).Limit(10).Rows()
		//SQLstr := "select aa.moment_id, tot1+tot2 as tot from (select moment_id, count(moment_id) as tot1 from moment_comment where is_deleted<>1 group by moment_id) as aa right outer join (select moment_id, count(is_liked) as tot2 from praise group by moment_id) as bb on aa.moment_id=bb.moment_id order by tot desc " + " limit " + strconv.Itoa(Page * 10) + ",10;"
		//rows, err := MysqlDB.Raw(SQLstr).Rows()
		if err == gorm.ErrRecordNotFound {
			return AllMoment, true // 说明这时候就是空的
		} else if err != nil {
			fmt.Println("====")
			fmt.Println(err)
			fmt.Println("====")
			return AllMoment, false
		}
		defer rows.Close()
		for rows.Next() {
			// 全扫描进结构体
			err := MysqlDB.ScanRows(rows, &MomentRecord)
			if err != nil {
				break
			}
			if len(MomentRecords) > 0 && MomentRecord.MomentId == MomentRecords[len(MomentRecords)-1].MomentId {
				break
			}
			MomentRecords = append(MomentRecords, MomentRecord)
		}

		//fmt.Println(len(MomentRecords))
		for _, record := range MomentRecords {
			moment := statements.Moment{}
			if err := MysqlDB.Where("id=?", record.MomentId).First(&moment).Error; err != nil {
				return AllMoment, false
			}
			AllMoment = append(AllMoment, moment)
		}
	} else {
		// 模糊查找
		Fuzzy := "%" + Keyword + "%"
		if err := MysqlDB.Where("content LIKE ? or song_name LIKE ? or state LIKE ?", Fuzzy, Fuzzy, Fuzzy).Order("created_at DESC").Offset(Page * 10).Limit(10).Find(&AllMoment).Error; err != nil {
			return AllMoment, false
		}
	}
	return AllMoment, true
}

// 创建新动态
func CreateMoment(Moment statements.Moment) bool {
	MysqlDB := setting.MysqlConn()
	if err := MysqlDB.Create(&Moment).Error; err != nil {
		return false
	}

	// 塞一条 is_like=0 的点赞记录，用作标记
	err := markMomentInPraise(int(Moment.ID))
	if err != nil {
		return false
	}

	// 塞一条 标记用的评论
	// 还没干（也许已经不用了

	return true
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
	fmt.Println(err)
	if err != nil {
		return -1
	}
	return Tot
}

// 创建新评论,返回创建好的评论的 id
//type CommentId struct {
//	Id int `gorm:"id"`
//}
func CreateComment(Comment statements.MomentComment) (int, bool) {
	MysqlDB := setting.MysqlConn()
	if err := MysqlDB.Create(&Comment).Error; err != nil {
		return 0, false
	}

	//var commentId CommentId
	//if err := MysqlDB.Model(&statements.MomentComment{}).Where("user_id=? and moment_id=? and comment=?", Comment.UserId, Comment.MomentId, Comment.Comment).Scan(&commentId).Error; err != nil {
	//	fmt.Println(err)
	//	return 0, false
	//}
	return int(Comment.ID), true
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

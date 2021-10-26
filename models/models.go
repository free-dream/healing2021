package models

import (
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

func TableInit() {
	go statements.AdvertisementInit()
	go statements.ClassicInit()
	go statements.CoverInit()
	go statements.PraiseInit()
	go statements.LotteryInit()
	go statements.MessageInit()
	go statements.MomentInit()
	go statements.MomentCommentInit()
	go statements.PrizeInit()
	go statements.SelectionInit()
	go statements.TaskInit()
	go statements.TaskTableInit()
	go statements.UserInit()
	time.Sleep(time.Second * 2)
}

// 假用户
func CreateFakeUser(nickname string, openid string, time time.Time) { /*hobby map[string]string) */
	User := statements.User{
		Openid:    openid,
		Nickname:  nickname,
		LoginTime: time,
	}

	db := setting.DB
	db.Create(&User)
}
func AddFakeUsers() {
	CreateFakeUser("heng1", "123456", time.Now())
	CreateFakeUser("heng2", "123456321", time.Date(2002, 12, 11, 10, 16, 55, 05, time.Local))
	CreateFakeUser("heng3", "1231", time.Date(2021, 10, 20, 10, 16, 55, 05, time.Local))
	CreateFakeUser("heng4", "99999", time.Now())
	CreateFakeUser("juryo", "juryo", time.Now())
}

// 假动态
func CreateFakeMoment(id int, likes int, content string, songName string, selectId int, states string, picture string) {
	Moment := statements.Moment{
		UserId:      id,
		Content:     content,
		SongName:    songName,
		SelectionId: selectId,
		State:       states,
		LikeNum:     likes,
	}

	db := setting.MysqlConn()
	db.Create(&Moment)
}
func AddFakeMoments() {
	CreateFakeMoment(1, 2, "第一条假动态", "第一首假点歌", 1, "状态1", "图片1地址")
	CreateFakeMoment(1, 2, "第二条假动态", "第二首假点歌", 1, "状态2", "图片2地址")
	CreateFakeMoment(1, 2, "第三条假动态", "第三首假点歌", 1, "状态3", "图片3地址")
	CreateFakeMoment(1, 2, "第四条假动态", "第四首假点歌", 1, "状态4", "图片4地址")
	CreateFakeMoment(1, 2, "第五条假动态", "第五首假点歌", 1, "状态5", "图片5地址")
	CreateFakeMoment(1, 2, "第六条假动态", "第六首假点歌", 1, "状态6", "图片6地址")
	CreateFakeMoment(1, 2, "第七条假动态", "第七首假点歌", 1, "状态7", "图片7地址")
	CreateFakeMoment(1, 2, "第八条假动态", "第八首假点歌", 1, "状态8", "图片8地址")
}

// 假评论
func CreateFakeComment(uid int, mid int, comment string, likes int) {
	Comment := statements.MomentComment{
		UserId:   uid,
		MomentId: mid,
		Comment:  comment,
		LikeNum:  likes,
	}

	db := setting.MysqlConn()
	db.Create(&Comment)
}
func AddFakeComments() {
	CreateFakeComment(1, 1, "第yi条假评论", 3)
	CreateFakeComment(1, 1, "第er条假评论", 3)
	CreateFakeComment(1, 1, "第san条假评论", 3)
	CreateFakeComment(1, 1, "第si条假评论", 3)
	CreateFakeComment(1, 1, "第wu条假评论", 3)
}

func AddFakeSelections() {
	for index := 1; index < 6; index++ {
		CreateFakeSelection(index, "测试歌曲")
	}
}

// 假评论
func CreateFakeSelection(uid int, name string) {
	selection := statements.Selection{
		UserId:   uid,
		SongName: name,
	}

	db := setting.MysqlConn()
	db.Create(&selection)
}
func AddFakeCovers() {
	for index := 1; index < 6; index++ {
		CreateFakeCovers(index+2, "测试歌曲", index+1)
	}
}

func CreateFakeCovers(uid int, name string, cid int) {
	cover := statements.Cover{
		UserId:      uid,
		SongName:    name,
		SelectionId: strconv.Itoa(cid),
	}

	db := setting.MysqlConn()
	db.Create(&cover)
}

// 造点测试用的假数据
func FakeData() {
	TableInit()
	AddFakeCovers()
	AddFakeUsers()
	AddFakeMoments()
	AddFakeComments()
	AddFakeSelections()
}

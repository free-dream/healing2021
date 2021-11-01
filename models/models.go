package models

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

func TableInit() {
	statements.AdvertisementInit()
	statements.ClassicInit()
	statements.CoverInit()
	statements.PraiseInit()
	statements.LotteryInit()
	statements.MessageInit()
	statements.MomentInit()
	statements.MomentCommentInit()
	statements.PrizeInit()
	statements.SelectionInit()
	statements.TaskInit()
	statements.TaskTableInit()
	statements.UserInit()
	statements.SysmsgInit()
	statements.UsrmsgInit()
}

// 假用户
func CreateFakeUser(nickname string, openid string, avatar string) { /*hobby map[string]string) */
	User := statements.User{
		Openid:   openid,
		Nickname: nickname,
		Avatar:   avatar,
	}

	db := setting.DB
	db.Create(&User)
}

//可以生成给定了学校和特定用户名的用户
func CreateDummyUser(model *statements.User) {
	db := setting.DB
	db.Create(&model)
}

//生成假用户
func AddFakeUsers() {
	avatar := "http://cdn.healing2020.100steps.top/static/personal/avatarFemale.png"
	CreateFakeUser("heng1", "123456", avatar)
	CreateFakeUser("heng2", "123456321", avatar)
	CreateFakeUser("heng3", "1231", avatar)
	CreateFakeUser("heng4", "99999", avatar)
	for i := 0; i < 10; i++ {
		CreateDummyUser(dummyUser())
	}
}

//假点歌
func CreateDummySelections(userid int) {
	style := StylePool[tools.GetRandomNumbers(len(StylePool))]
	language := LanguagePool[tools.GetRandomNumbers(len(LanguagePool))]
	song := SongPool[tools.GetRandomNumbers(len(SongPool))]
	selection, err := dummySelections(userid, song, language, style)
	if err != nil {
		panic(err)
	}

	db := setting.DB
	db.Create(&selection)
}

func CreateFakeSelection(uid int, name string) {
	selection := statements.Selection{
		UserId:   uid,
		SongName: name,
		Module:   tools.GetRandomNumbers(2) + 1,
	}

	db := setting.MysqlConn()
	db.Create(&selection)
}

//目前总计181个点歌请求
func AddFakeSelections() {
	for index := 1; index < 6; index++ {
		CreateFakeSelection(index, "测试歌曲")
	}

	//目前目录下的假用户有15个,每个用户生成15条点歌需求
	for i := 1; i < 15; i++ {
		for j := 0; j < 15; j++ {
			CreateDummySelections(i)
		}
	}
}

//假翻唱,随机取
func CreateDummyCovers() {
	//随机选20首歌(可复选)生成10个翻唱，用户随机取
	for i := 1; i < 20; i++ {
		rand1 := tools.GetRandomNumbers(215)
		if rand1 == 0 {
			rand1 = 1
		}
		rand2 := tools.GetRandomNumbers(14)
		if rand2 == 0 {
			rand2 = 1
		}

		temp, err := dummyCovers(rand2, rand1, -1)
		if err != nil {
			panic(err)
		}
		db := setting.MysqlConn()
		db.Create(temp)
	}
}
func CreateFakeCovers(uid int, name string, cid int, classicId int, module int) {
	cover := statements.Cover{
		UserId:      uid,
		SongName:    name,
		SelectionId: strconv.Itoa(cid),
		Module:      module,
		ClassicId:   classicId,
	}

	db := setting.MysqlConn()
	db.Create(&cover)
}
func AddFakeCovers() {
	// 经典翻唱 5
	for index := 1; index < 6; index++ {
		CreateFakeCovers(index+2, "songName"+strconv.Itoa(index), index+1, 0, 1)
	}

	// 童年翻唱 3x13=39
	for index := 1; index < 14; index++ {
		CreateFakeCovers(index+2, "songName"+strconv.Itoa(index), index+1, index, 2)
		CreateFakeCovers(index+2, "songName"+strconv.Itoa(index), index+1, index, 2)
		CreateFakeCovers(index+2, "songName"+strconv.Itoa(index), index+1, index, 2)
	}
	CreateDummyCovers()
}

//向记录中增加奖品记录,奖品记录为真
func CreateLottery() {

}

//假动态
func CreateFakeMoment(id int, likes int, content string, songName string, selectId int, states string, picture string) {
	Moment := statements.Moment{
		UserId:      id,
		Content:     content,
		SongName:    songName,
		SelectionId: selectId,
		State:       states,
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

// 假童年原唱(名字对不上假翻唱
func CreateFakeClassic(remark string, songName string, icon string, singer string, workName string, click int, file string) {
	Comment := statements.Classic{
		Remark:   remark,
		SongName: songName,
		Icon:     icon,
		Singer:   singer,
		WorkName: workName,
		Click:    click,
		File:     file,
	}

	db := setting.MysqlConn()
	db.Create(&Comment)
}
func AddFakeClassic() {
	for i := 0; i < 15; i++ {
		remark := "fake remark " + strconv.Itoa(i)
		songName := "songName" + strconv.Itoa(i)
		icon := "icon" + strconv.Itoa(i)
		singer := "singer" + strconv.Itoa(i)
		workName := "workName" + strconv.Itoa(i)
		click := 10 + i
		file := "file" + strconv.Itoa(i)
		CreateFakeClassic(remark, songName, icon, singer, workName, click, file)
	}
}

// 造点测试用的假数据
func FakeData() {
	AddFakeUsers()
	AddFakeMoments()
	AddFakeComments()
	AddFakeSelections()
	AddFakeCovers()
	AddFakeClassic()
}

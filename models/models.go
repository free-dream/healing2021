package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"

	"github.com/jinzhu/gorm"
)

//奖品是真实概率数据
//目前设计三个奖项，一等奖 2%，二等奖 8%，三等奖 20%
const (
	PRIZE1 = "蓝牙耳机/八音盒"
	PRIZE2 = "有线耳机"
	PRIZE3 = "小玩偶/台灯"
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
	statements.DevotionInit()
	//任务和奖品初始化,不用删
	AddTask()
	AddLotteries()
}

/*-------------------------系统启动时初始化奖品和任务表---------------------------*/
//真奖品
func Lottery(name string, possilbity float64) *statements.Lottery {
	lottery := statements.Lottery{
		Name:        name,
		Possibility: possilbity,
	}
	return &lottery
}

//真任务
func Task(text string, max int) *statements.Task {
	task := statements.Task{
		Text: text,
		Max:  max,
	}
	return &task
}

//任务记录
func CreateTask() {
	db := setting.DB
	db.Exec("truncate table task")
	db.Create(Task("点歌一次", 8))
	db.Create(Task("治愈一次", -1))
	db.Create(Task("发动态一次", 8))
}

//向记录中增加奖品记录,奖品记录为真
func CreateLottery() {
	db := setting.DB
	db.Exec("truncate table lottery")
	db.Create(Lottery(PRIZE1, 0.02))
	db.Create(Lottery(PRIZE2, 0.08))
	db.Create(Lottery(PRIZE3, 0.2))
}

func AddTask() {
	CreateTask()
}

func AddLotteries() {
	CreateLottery()
}

/*-------------------------------其下为假数据----------------------------------*/
func CreateFakeUser(nickname string, openid string, avatar string, realname string, real int, tel string, telcheck int) { /*hobby map[string]string) */
	User := statements.User{
		Openid:         openid,
		Nickname:       nickname,
		Avatar:         avatar,
		Points:         0,
		RealName:       realname,
		RealNameSearch: real,
		PhoneNumber:    tel,
		PhoneSearch:    telcheck,
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
	CreateFakeUser("heng1", "123456", avatar, "代知言", 1, "15838896053", 0)
	CreateFakeUser("heng2", "123456321", avatar, "代知言", 0, "15838896053", 1)
	CreateFakeUser("heng3", "1231", avatar, "言知代", 1, "15838896054", 1)
	CreateFakeUser("heng4", "99999", avatar, "言知代", 1, "15838896055", 1)
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

func AddFakeHomeC() {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	rows, _ := db.Table("cover").Rows()
	defer rows.Close()
	coverDetails := dao.CoverDetails{}
	cover := statements.Cover{}
	for rows.Next() {
		db.ScanRows(rows, &cover)
		db.Table("user").Select("cover.selection_id,cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)

		coverDetails.CreatedAt = tools.DecodeTime(cover.CreatedAt)
		value, _ := json.Marshal(coverDetails)

		if cover.Style != "" {
			redisCli.RPush("healing2021:cover."+"1"+"."+cover.Style, string(value))
		}
		if cover.Language != "" {
			redisCli.RPush("healing2021:cover."+"1"+"."+cover.Language, string(value))
		}
		redisCli.RPush("healing2021:cover."+"1"+"."+"all", string(value))

	}
}

func AddFakeHomeS() {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	rows, _ := db.Table("selection").Rows()
	defer rows.Close()
	selectionDetails := dao.SelectionDetails{}
	selection := statements.Selection{}
	for rows.Next() {
		db.ScanRows(rows, &selection)
		db.Table("user").Select("user.sex,selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at,selection.remark").Where("selection.id=?", selection.ID).Joins("left join selection on user.id=selection.user_id").Scan(&selectionDetails)
		selectionDetails.CreatedAt = tools.DecodeTime(selection.CreatedAt)
		value, _ := json.Marshal(selectionDetails)
		if selection.Style != "" {
			redisCli.RPush("healing2021:selection"+"."+selection.Style, string(value))
		}
		if selection.Language != "" {
			redisCli.RPush("healing2021:selection"+"."+selection.Language, string(value))
		}
		redisCli.RPush("healing2021:selection"+"."+"all", string(value))

	}
}

func CreateFakeSelection(uid int, name string) {
	selection := statements.Selection{
		UserId:   uid,
		SongName: name,
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
func CreateFakeCovers(uid int, name string, sid int, classicId int, module int) {
	cover := statements.Cover{
		UserId:      uid,
		SongName:    name,
		SelectionId: sid,
		Module:      module,
		ClassicId:   classicId,
		Nickname:    "测试小子",
		Language:    "中文",
		File:        "address",
		Style:       "cool",
		Avatar:      "头像",
	}

	db := setting.MysqlConn()
	db.Create(&cover)
}

//func AddFakeCovers() {
//	// 经典翻唱 5
//	for index := 1; index < 19; index++ {
//		CreateFakeCovers(index+2, "songName"+strconv.Itoa(index), index+1, 0, 1)
//	}
//
//	CreateDummyCovers()
//}

//假翻唱点赞表
func CreatePraise() bool {
	db := setting.DB
	temp, c1, c2 := dummyLikes()
	var praise statements.Praise
	if err := MysqlDb.Where("user_id = ? AND cover_id = ?", c1, c2).First(&praise).Error; gorm.IsRecordNotFoundError(err) {
		db.Create(temp)
		return true
	}
	return false
}

func AddFakePraises() {
	i := 0
	for {
		if CreatePraise() {
			i++
		}
		if i > 30 {
			break
		}
	}
}

//假动态
func CreateFakeMoment(id int, content string, songName string, selectId int, states string) {
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
	for i := 1; i < 60; i++ {
		CreateFakeMoment(i, "第"+strconv.Itoa(i)+"条假动态", "第"+strconv.Itoa(i)+"首假点歌", 1, "状态"+strconv.Itoa(i))
	}
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
	//AddFakeUsers()
	//AddFakeMoments()
	//AddFakeComments()
	//AddFakeSelections()
	CreateFakeCovers(1, "刀剑如梦", 0, 45, 2)
	CreateFakeCovers(1, "刀剑如梦", 0, 45, 2)
	CreateFakeCovers(1, "刀剑如梦", 0, 45, 2)
	//AddFakePraises()
	//AddFakeClassic()
}
func AddClassic() {
	csc, err := os.Open("classic.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csc.Close()
	readCsv := csv.NewReader(csc)
	readAll, err := readCsv.ReadAll()
	db := setting.MysqlConn()
	for _, list := range readAll {
		classic := statements.Classic{
			SongName: list[1],
			WorkName: list[2],
			File:     list[3],
			Icon:     list[4],
			Singer:   list[5],
		}
		err = db.Table("classic").Create(&classic).Error
		if err != nil {
			fmt.Println(classic)
		}
	}

}
func AddDevotion() {
	csc, err := os.Open("devotion.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csc.Close()
	readCsv := csv.NewReader(csc)
	readAll, err := readCsv.ReadAll()
	fmt.Println(readAll)
	db := setting.MysqlConn()
	for _, list := range readAll {
		dev := statements.Devotion{
			SongName: list[1],

			File: list[2],

			Singer: list[3],
		}
		err = db.Table("devotion").Create(&dev).Error
		if err != nil {
			fmt.Println(dev)
		}
	}

}

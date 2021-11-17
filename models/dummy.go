package models

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"math/rand"
)

//测试学校名
const (
	SCUT  = "华南理工大学"
	SYU   = "中山大学"
	JU    = "暨南大学"
	SCNU  = "华南师范大学"
	OTHER = "其它大学"
)

//靶用户名数据，主要用于查询测试
const (
	TARGET1 = "bbt2021ad"
	TARGET2 = "bbt1021ad"
	TARGET3 = "bbt21bc"
)

//测试用歌曲名
const (
	G1 = "稻香"
	G2 = "忐忑"
	G3 = "海阔天空" //可粤语可中文
	G4 = "光辉岁月"
	G5 = "富士山下"
	G6 = "砂之惑星" //日文歌
	G7 = "向夜晚奔去"
	G8 = "蓝二乘"
	G9 = "初音未来的消失"
	GA = "viva la vida" //英文歌
	GB = "Numb"
	GC = "Never Gonna Give You Up"
	GD = "Monster"
)

//童年曲目
const (
	CH1 = "葫芦娃"
	CH2 = "黑猫警长"
	CH3 = "邋遢大王奇遇记"
	CH4 = "小英雄哪吒"
)

//语言
const (
	L1 = "Chinese"
	L2 = "Cantonese"
	L3 = "Japanese"
	L4 = "English"
	L5 = "Other"
)

//风格
const (
	S1 = "pop"
	S2 = "classical"
	S3 = "ACG"
	S4 = "Rock"
	S5 = "Tiktok"
	S6 = "Other"
)

//欢迎根据需要扩充测试用例,常量更新完后添加到全局变量的列表里
//由于翻唱是个性要求，所以录音的歌唱语言、歌唱风格与歌曲本身不做绑定
//也就是说，用户完全可以要求一首摇滚版的大悲咒
var (
	SchoolPool = []string{SCUT, SYU, JU, SCNU, OTHER}
	TargetPool = []string{TARGET1, TARGET2, TARGET3}
	// PrizePool    = []string{PRIZE1, PRIZE2, PRIZE3}
	StylePool    = []string{S1, S2, S3, S4, S5, S6}
	SongPool     = []string{G1, G2, G3, G4, G5, G6, G7, G8, G9, GA, GB, GC, GD}
	LanguagePool = []string{L1, L2, L3, L4, L5}
	ChildPool    = []string{CH1, CH2, CH3, CH4}
	MysqlDb      = setting.MysqlConn()
)

//获取用户id
func GetUserid(openid string) (int, error) {
	var user tables.User
	err := MysqlDb.Where("OpenId = ?", openid).First(&user).Error
	if err != nil {
		return -1, err
	}
	return int(user.ID), nil
}

//获取用户的nickname
func GetUserNickname(userid int) (string, error) {
	var user tables.User
	err := MysqlDb.Where("Id = ?", userid).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Nickname, nil
}

//获取用户的avatar
func GetUserAvatar(userid int) (string, error) {
	var user tables.User
	err := MysqlDb.Where("Id = ?", userid).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Avatar, nil
}

//获取selection的通用信息
func GetAttr(selectionid int) (string, string, string, error) {
	var selection tables.Selection
	err := MysqlDb.Where("Id = ?", selectionid).First(&selection).Error
	if err != nil {
		return "", "", "", err
	}
	return selection.SongName, selection.Language, selection.Style, nil
}

//假点赞表
func dummyLikes() (*statements.Praise, int, int) {
	check1 := tools.GetRandomNumbers(10) + 1
	check2 := tools.GetRandomNumbers(29) + 1
	return &statements.Praise{
		CoverId: check2,
		UserId:  check1,
	}, check1, check2
}

//生成dummy用户
func dummyUser() *statements.User {
	check1 := tools.GetRandomNumbers(4)
	check2 := tools.GetRandomNumbers(5)
	var nickname, school string
	//决定nickname
	if check1 != 3 {
		nickname = TargetPool[check1]
	} else {
		nickname = string(tools.GetRandomString(4))
	}
	//决定学校

	avatar, _ := GetUserAvatar(rand.Intn(2))
	school = SchoolPool[check2]
	user := statements.User{
		Openid:    string(tools.GetRandomString(10)),
		Avatar:    avatar,
		Nickname:  nickname,
		RealName:  string(tools.GetRandomString(6)),
		Signature: string(tools.GetRandomString(20)),
		School:    school,
	}

	return &user
}

//基于用户创建点歌
func dummySelections(userid int, song string, language string, style string) (*statements.Selection, error) {
	avatar, err := GetUserAvatar(userid)
	if err != nil {
		return nil, err
	}
	selection := statements.Selection{
		SongName: song,
		Remark:   string(tools.GetRandomString(20)),
		Language: language,
		UserId:   userid,
		Avatar:   avatar,
		Style:    style,
	}
	return &selection, nil
}

//基于用户创建翻唱,若为非童年,classicid = -1
//翻唱基于点歌要求而存在，所以一定和点歌要求相同
func dummyCovers(userid int, selectionid int, classicid int) (*statements.Cover, error) {
	nickname, err := GetUserNickname(userid)
	if err != nil {
		return nil, err
	}
	avatar, err := GetUserAvatar(userid)
	if err != nil {
		return nil, err
	}
	song, language, style, err := GetAttr(selectionid)
	if err != nil {
		return nil, err
	}
	cover := statements.Cover{
		UserId:      userid,
		Nickname:    nickname,
		Avatar:      avatar,
		SelectionId: selectionid,
		ClassicId:   classicid,
		File:        string(tools.GetRandomString(30)),
		SongName:    song,
		Language:    language,
		Style:       style,
	}
	return &cover, nil
}

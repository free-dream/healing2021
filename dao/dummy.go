package dao

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

const (
	//可扩充
	SCUT  = "华南理工大学"
	SYU   = "中山大学"
	JU    = "暨南大学"
	SCNU  = "华南师范大学"
	OTHER = "其它大学"

	TARGET1 = "bbt2021ad" //靶数据，主要用于查询测试
	TARGET2 = "bbt1021ad"
	TARGET3 = "bbt21bc"

	PRIZESP = "特奖" //目前设计四个奖项，特奖 1%，一等奖 5%，二等奖 15%，三等奖 29%
	PRIZE1  = "一等奖"
	PRIZE2  = "二等奖"
	PRIZE3  = "三等奖"

	//欢迎扩充测试用例
	//流行歌
	CSONG1 = "一剪梅" //中文歌
	CSONG2 = "稻香"
	CSONG3 = "忐忑"
	CSONG4 = "海阔天空" //可粤语可中文
	CSONG5 = "光辉岁月"
	CSONG6 = "富士山下"
	JSONG1 = "砂之惑星" //日文歌
	JSONG2 = "向夜晚奔去"
	JSONG3 = "蓝二乘"
	JSONG4 = "初音未来的消失"
	ESONG1 = "viva la vida" //英文歌
	ESONG2 = "Numb"
	ESONG3 = "Never Gonna Give You Up"
	ESONG4 = "Monster"
	//童年曲目
	CHILDHOOD1 = "葫芦娃"
	CHILDHOOD2 = "黑猫警长"
	CHILDHOOD3 = "邋遢大王奇遇记"
	CHILDHOOD4 = "小英雄哪吒"
)

var (
	SchoolPool = []string{SCUT, SYU, JU, SCNU, OTHER}
	TargetPool = []string{TARGET1, TARGET2, TARGET3}
	Prizetool  = []string{PRIZESP, PRIZE1, PRIZE2, PRIZE3}
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
	school = SchoolPool[check2]
	user := statements.User{
		Openid:    string(tools.GetRandomString(10)),
		Nickname:  nickname,
		RealName:  string(tools.GetRandomString(6)),
		Signature: string(tools.GetRandomString(20)),
		School:    school,
	}

	return &user
}

//假彩票
func fakeLotteries(name string, possilbity float64) *statements.Lottery {
	lottery := statements.Lottery{
		Name:        name,
		Possibility: possilbity,
	}
	return &lottery
}

//基于用户创建点歌
func dummySelections(userid int, song string, language string) (*statements.Selection, error) {
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
	}
	return &selection, nil
}

//基于用户创建翻唱,若为非童年,classicid = -1
func dummyCovers(userid int, selectionid int, songname string, classicid int) (*statements.Cover, error) {
	nickname, err := GetUserNickname(userid)
	if err != nil {
		return nil, err
	}
	avatar, err := GetUserAvatar(userid)
	if err != nil {
		return nil, err
	}
	cover := statements.Cover{
		UserId:      userid,
		Nickname:    nickname,
		Avatar:      avatar,
		SelectionId: strconv.Itoa(selectionid),
		ClassicId:   classicid,
		File:        string(tools.GetRandomString(30)),
		SongName:    songname,
	}
	return &cover, nil
}

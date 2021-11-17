package dao

import (
	"strconv"

	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

const (
	a = "200-500"
	b = "500--1000"
	c = ">1000"
)

//基于学校获取排名
func GetRankingBySchool(school string) ([]respModel.RankingResp, error) {
	mysqlDb := db.MysqlConn()

	var users []respModel.RankingResp
	var err error = nil
	if school != "All" {
		err = mysqlDb.
			Limit(10).
			Table("user").
			Select("id as userid,avatar,nickname").
			Where("School = ?", school).
			Order("points desc").
			Find(&users).
			Error
	} else if school == "All" {
		err = mysqlDb.
			Limit(10).
			Table("user").
			Select("id as userid,avatar,nickname").
			Order("points desc").
			Find(&users).
			Error
	}
	if err != nil {
		panic(err)
		return nil, err
	}
	return users, nil
}

//获取当前用户排名,之后就直接读缓存
func GetRankByCUserId(userid int) (string, error) {
	mysqlDb := db.MysqlConn()
	var users []tables.User
	//同分以字母序为准
	err := mysqlDb.Limit(1000).
		Order("Points desc,Nickname").
		Find(&users).Error
	if err != nil {
		return "", err
	}
	i := 0
	for ; i < len(users); i++ {
		if int(users[i].ID) == userid {
			break
		}
	}
	var resp string
	if i <= 199 {
		resp = strconv.Itoa(i + 1)
	} else if i <= 499 {
		resp = a
	} else if i <= 999 {
		resp = b
	} else if i == 1000 {
		resp = c
	} else {
		resp = ""
	}
	return resp, nil
}

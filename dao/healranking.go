package dao

import (
	"strconv"

	tables "git.100steps.top/100steps/healing2021_be/models/statements"
	db "git.100steps.top/100steps/healing2021_be/pkg/setting"
)

const (
	a = "200-500"
	b = "500--1000"
	c = ">1000"
)

//基于学校获取排名，仅有sql版
func GetRankingBySchool(school string) ([]tables.User, error) {
	mysqlDb := db.MysqlConn()

	var users []tables.User
	var err error = nil
	if school != "All" {
		//按Points倒序排列
		err = mysqlDb.Where("School = ?", school).Order("Points desc").Limit(10).Find(&users).Error
	} else if school == "All" {
		err = mysqlDb.Order("Points desc").Limit(10).Find(&users).Error
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

//获取当前用户排名,只在登陆时获取一次，之后就直接读缓存
func GetRankByCUserId(userid int) (string, error) {
	mysqlDb := db.MysqlConn()
	var users []tables.User
	//同分以字母序为准
	err := mysqlDb.Limit(1000).Order("Points desc,Nickname").Find(&users).Error
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

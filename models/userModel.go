package models

import (
	"errors"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type User struct {
	ID             int
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	RealName       string `json:"real_name"`
	PhoneNumber    string `json:"phone_number"`
	Sex            int    `json:"sex"`
	School         string `json:"school"`
	Avatar         string `json:"avatar"`
	AvatarVisible  int    `json:"avatar_visible"`
	PhoneSearch    int    `json:"phone_search"`
	RealNameSearch int    `json:"real_name_search"`
	Signature      string `json:"signature"`
}

func FakeCreateUser(user *User) (string, error) {

	err := setting.DB.Table("user").Create(&user).Error
	if err != nil {
		return "", err
	}
	openid := user.Nickname
	return openid, nil

}

func CreateUser(user *User) (error, error) {
	count := 0
	setting.DB.Table("user").Where("nickname=?", user.Nickname).Count(&count)
	if count != 0 {
		return errors.New("error"), nil
	}
	err := setting.DB.Table("user").Create(&user).Error
	return nil, err

}
func UpdateUser(user *User, openid string) error {
	message := make(map[string]interface{})
	message["avatar"] = user.Avatar
	message["nickname"] = user.Nickname
	message["avatar_visible"] = user.AvatarVisible
	message["phone_search"] = user.PhoneSearch
	message["real_name_search"] = user.RealNameSearch
	message["signature"] = user.Signature
	err := setting.DB.Table("user").Where("openid=?", openid).Update(message).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUser(openid string) interface{} {
	user := statements.User{}
	resp := make(map[string]interface{})
	setting.DB.Table("user").Where("openid=?", openid).First(&user)
	selection := statements.Selection{}
	cover := statements.Cover{}
	moment := statements.Moment{}

	praise := statements.Praise{}

	resp["message"] = user
	resp["mySelections"] = get(user.ID, "selection", "user_id=?", selection)
	resp["mySongs"] = get(user.ID, "cover", "user_id=?", cover)
	resp["moments"] = get(user.ID, "moment", "user_id=?", moment)
	resp["myLikes"] = getPraise(praise, user.ID)

	return resp
}

func UpdateBackground(openid string, background string) error {
	err := setting.DB.Table("user").Where("openid=?", openid).Update("background", background).Error
	return err
}

func GetCallee(id int) interface{} {
	user := statements.User{}
	resp := make(map[string]interface{})
	setting.DB.Table("user").Where("id=?", id).Scan(&user)
	selection := statements.Selection{}
	cover := statements.Cover{}
	moment := statements.Moment{}

	praise := statements.Praise{}

	resp["message"] = user
	resp["mySelections"] = get(user.ID, "selection", "user_id=?", selection)
	resp["mySongs"] = get(user.ID, "cover", "user_id=?", cover)
	resp["moments"] = get(user.ID, "moment", "user_id=?", moment)
	resp["myLikes"] = getPraise(praise, user.ID)

	return resp

}

//返回所有信息
func get(value interface{}, tableName string, condition string, obj interface{}) interface{} {
	index := 0
	content := make(map[int]interface{})
	rows, err := setting.DB.Table(tableName).Where(condition, value).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		content[index] = obj
		index++
	}
	return content
}
func getPraise(praise statements.Praise, id uint) interface{} {
	index := 0
	content := make(map[int]interface{})
	rows, err := setting.DB.Table("praise").Where("user_id=?", id).Not("cover_id", "").Select("cover_id").Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	cover := statements.Cover{}
	for rows.Next() {
		err := setting.DB.ScanRows(rows, &praise)
		if err != nil {
			panic(err)
		}
		setting.DB.Table("cover").Where("id=?", praise.CoverId).Scan(cover)
		content[index] = cover
		index++
	}
	return content
}

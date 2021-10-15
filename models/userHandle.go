package models

import (
	"errors"
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type User struct {
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
	fmt.Println(message)
	err := setting.DB.Table("user").Where("openid=?", openid).Update(message).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUser(openid string) {
	/*user:=statements.User{}
	selection:=statements.Selection{}
	song:=statements.Song{}
	momentLike:=statements.MomentLike{}
	coverLike:=statements.CoverLike{}
	songLike:=statements.SongLike{}
	setting.DB.Table("user").Where("openid=?",openid).First(&user)

	setting.DB.Table("selection").Where("user_id=?",user.ID).
	setting.DB*/

}

func UpdateBackground() {

}

func GetCallee() {

}

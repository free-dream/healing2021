package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type UsrMsg struct {
	ID        int    `json:"selectionId"`
	Style     string `json:"style"`
	CreatedAt string `json:"created_at"`
	SongName  string `json:"song_name"`
	Remark    string `json:"remark"`
	Nickname  string `json:"nickname"`
}
type CovMsg struct {
	UserId   int    `json:"user_id"`
	Likes    int    `json:"likes"`
	Nickname string `json:"nickname"`
}

var rows *sql.Rows
var err error

//处理治愈详情页
//点赞数debug，尚未测试
//结构体疑似有bug
func GetHealingPage(selectionId int) (interface{}, error) {
	userMsg := UsrMsg{}
	resp := make(map[string]interface{})
	setting.DB.Table("selection").Select("selection.id,selection.song_name,selection.style,selection.created_at,selection.remark,user.nickname").Joins("left join user on user.id=selection.user_id").Where("selection.id=?", selectionId).Scan(&userMsg)
	/*if err!=nil{
		panic(err)
		return nil, err
	}*/
	rows, err = setting.DB.Table("cover").Select("cover.user_id,cover.likes,user.nickname").Joins("left join user on user.id=cover.user_id").Where("cover.selection_id=?", selectionId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	obj := CovMsg{}
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		if err != nil {
			return nil, err
		}
		content[index] = obj
		index++
	}
	resp["user"] = userMsg
	resp["singer"] = content

	return resp, err
}

//获取广告url
func GetAds() (interface{}, error) {
	rows, err = setting.DB.Table("advertisement").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	var obj interface{}
	content := make(map[int]interface{})
	for rows.Next() {
		setting.DB.ScanRows(rows, &obj)
		content[index] = obj
		index++
	}

	return content, err

}

type Tags struct {
	Sort     int    `json:"sort" binding:"required"`
	Style    string `json:"style"`
	Language string `json:"language"`
	RankWay  int    `json:"rankWay" binding:"required"`
}
type SelectionDetails struct {
	Nickname  string `json:"nickname"`
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
}

//获取点歌页以推荐方式
//由于推荐方法尚未知道，暂未完成

func GetSelections(tag Tags) (interface{}, error) {
	switch tag.Sort {
	case 1:
		switch tag.RankWay {
		case 1:
		case 2:
		}
	case 2:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880").Joins("left join selection on user.id=selection.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Joins("left join selection on user.id=selection.user_id").Order("selection.created_at DESC").Rows()
		}
	case 3:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,selection.song_name,selection.created_at").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and selection.style=?", tag.Style).Joins("left join selection on user.id=selection.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Where("selection.style=?", tag.Style).Joins("left join selection on user.id=selection.user_id").Order("created_at DESC").Rows()
		}
	case 4:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,selection.song_name,selection.created_at").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and selection.language=?", tag.Language).Joins("left join selection on user.id=selection.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Where("selection.language=?", tag.Language).Joins("left join selection on user.id=selection.user_id").Order("created_at DESC").Rows()
		}
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	index := 0
	selectionDetails := SelectionDetails{}
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &selectionDetails)
		if err != nil {
			break
		}
		content[index] = selectionDetails
		index++
	}
	return content, nil

}

type CoverDetails struct {
	Nickname  string `json:"nickname"`
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
	File      string `json:"file"`
}

func GetCovers(tag Tags) (interface{}, error) {
	switch tag.Sort {
	case 1:
		switch tag.RankWay {
		case 1:
		case 2:
		}
	case 2:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880").Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()

		case 2:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Joins("left join cover on user.id=cover.user_id").Order("created_at DESC").Rows()

		}
	case 3:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and cover.style=?", tag.Style).Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()

			fmt.Println(err)
		case 2:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.style=?", tag.Style).Joins("left join cover on user.id=cover.user_id").Order("created_at DESC").Rows()

			fmt.Println(err)
		}
	case 4:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and cover.language=?", tag.Language).Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()
			fmt.Println(err)

		case 2:
			rows, err = setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.language=?", tag.Language).Joins("left join cover on user.id=cover.user_id").Order("created_at DESC").Rows()

		}
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	coverDetails := CoverDetails{}
	index := 0
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &coverDetails)
		if err != nil {
			break
		}
		content[index] = coverDetails
		index++
	}
	return content, nil

}

//治愈系对应的录音

//翻唱点赞
func Like(coverId int, openid string) error {
	praise := statements.Praise{}
	user := User{}
	setting.DB.Table("user").Where("openid=?", openid).Scan(&user)
	setting.DB.Table("praise").Where("cover_id=? and user_id=?", coverId, user.ID).Scan(&praise)
	/*redis读写点赞数操作


	 */

	switch praise.IsLiked {
	case 1:
		setting.DB.Table("praise").Where("cover_id=? and user_id=?", coverId, user.ID).Update("is_liked", 0)

	}
	return err

}

func CreateRecord(id string, file string, uid int) (CoverDetails, error) {
	intId, _ := strconv.Atoi(id)
	selectionId := intId
	db := setting.MysqlConn()
	userId := uid

	var selection statements.Selection
	result1 := db.Model(&statements.Selection{}).Where("id=?", selectionId).First(&selection)
	if result1.Error != nil {
		return CoverDetails{}, errors.New("selection_id is unvalid")
	}
	var cover statements.Cover
	cover.SelectionId = strconv.Itoa(selectionId)
	cover.UserId = userId
	cover.SongName = selection.SongName
	cover.Likes = 0
	cover.File = file
	cover.Style = selection.Style
	cover.Language = selection.Language
	coverDetails := CoverDetails{}
	err := db.Model(&statements.Cover{}).Create(&cover).Error
	setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	value, err := json.Marshal(coverDetails)
	if err != nil {
		return coverDetails, err
	}
	if cover.Style != "" {
		setting.RedisClient.RPush("cover"+cover.Style, string(value))
	}
	if cover.Language != "" {
		setting.RedisClient.RPush("cover"+cover.Language, string(value))
	}
	setting.RedisClient.RPush("coverall", string(value))
	return coverDetails, err
}

func Select(selection statements.Selection) (SelectionDetails, error) {
	err = setting.DB.Table("selection").Create(&selection).Error
	selectionDetails := SelectionDetails{}
	fmt.Println(selection.ID)
	err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Where("selection.id=?", selection.ID).Joins("left join selection on user.id=selection.user_id").Scan(&selectionDetails).Error
	value, err := json.Marshal(selectionDetails)
	if err != nil {
		return selectionDetails, err
	}
	if selection.Style != "" {
		setting.RedisClient.RPush("selection"+selection.Style, string(value))
	}
	if selection.Language != "" {
		setting.RedisClient.RPush("selection"+selection.Language, string(value))
	}
	setting.RedisClient.RPush("selectionall", string(value))
	return selectionDetails, err
}

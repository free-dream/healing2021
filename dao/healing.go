package dao

import (
	"database/sql"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type UsrMsg struct {
	ID        int    `json:"selectionId"`
	Style     string `json:"style"`
	CreatedAt string `json:"created_at"`
	SongName  string `json:"song_name":""`
	Remark    string `json:"remark":""`
	Nickname  string `json:"nickname":""`
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
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &content)
		if err != nil {
			return nil, err
		}
		content[index] = content
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
			rows, err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Joins("left join selection on user.selection_id=selection.id").Order("created_at DESC").Rows()
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
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &content)
		if err != nil {
			break
			return nil, err
		}
		content[index] = content
		index++
	}
	return content, nil

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
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880").Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Joins("left join  on user.cover_id=cover.id").Order("created_at DESC").Rows()
		}
	case 3:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and cover.style=?", tag.Style).Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.style=?", tag.Style).Joins("left join cover on user.id=cover.user_id").Order("created_at DESC").Rows()
		}
	case 4:
		switch tag.RankWay {
		case 1:
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("TIMESTAMPDIFF(second,user.login_time,now())<2880 and cover.language=?", tag.Language).Joins("left join cover on user.id=cover.user_id").Order("rand()").Rows()
		case 2:
			rows, err = setting.DB.Table("user").Select("cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.language=?", tag.Language).Joins("left join cover on user.id=cover.user_id").Order("created_at DESC").Rows()
		}
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	index := 0
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &content)
		if err != nil {
			break
			return nil, err
		}
		content[index] = content
		index++
	}
	return content, nil

}

//治愈系对应的录音

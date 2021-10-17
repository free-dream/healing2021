package dao

import (
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
	rows, err := setting.DB.Table("cover").Select("cover.user_id,cover.likes,user.nickname").Joins("left join user on user.id=cover.user_id").Where("cover.selection_id=?", selectionId).Rows()
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	index := 0
	content := make(map[int]interface{})
	for rows.Next() {
		err := setting.DB.ScanRows(rows, &content)
		if err != nil {
			panic(err)
			return nil, err
		}
		content[index] = content
		index++
	}
	resp["user"] = userMsg
	resp["singer"] = content

	return resp, err
}

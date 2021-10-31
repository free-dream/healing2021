package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"math/rand"
	"sort"
	"strconv"
	"time"
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
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Likes    int    `json:"likes"`
	Nickname string `json:"nickname"`
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
	rows, err := setting.DB.Table("cover").Select("cover.user_id,user.nickname,cover.id").Joins("left join user on user.id=cover.user_id").Where("cover.selection_id=?", selectionId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	obj := CovMsg{}
	content := make(map[int]interface{})
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		setting.DB.Table("praise").Where("cover_id=? and is_liked", obj.ID, 0).Count(&obj.Likes)
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
	rows, err := setting.DB.Table("advertisement").Rows()
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
	Label   string `json:"label"`
	RankWay int    `json:"rankWay" binding:"required"`
	Page    int    `json:"page" binding:"required"`
}

type SelectionDetails struct {
	Nickname  string `json:"nickname"`
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
}

//为分页器做缓存
//缓存首次请求的结果,非实时,不可更新
//且由于每次分页资源不同所以key可相同
//例首页:home
func Cache(key string, resp interface{}) {
	value, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	err = setting.RedisClient.HSet(key, "cache", value).Err()
	fmt.Println(err)
}

//分页器
//从分页器里面取出
func Pager(key string, page int) (interface{}, error) {
	var resp []interface{}
	by, err := setting.RedisClient.HGet(key, "cache").Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(by, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp)/10 > page {
		return resp[(page-1)*10 : (page-1)*10+10], nil
	} else {
		return resp[(page-1)*10:], nil
	}

}

//获取点歌页
//module1表示治愈系，2表示童年，但返回参数可能不同
//可参考
func GetSelections(module string, id int, tag Tags) (interface{}, error) {
	index := 0
	var resp []SelectionDetails
	var err error
	if tag.Label == "recommend" {
		var hobby []string
		by, err := setting.RedisClient.HGet("hobby", strconv.Itoa(id)).Bytes()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(by, &hobby) //解析json
		if err != nil {
			return nil, err
		}
		if hobby == nil {
			hobby = []string{"all"}
		}
		var size int
		for _, value := range hobby {
			lenth := setting.RedisClient.LLen("selection" + module + value).Val()
			size += int(lenth)
		}
		resp = make([]SelectionDetails, size)
		for _, value := range hobby {
			if setting.RedisClient.Exists("selection"+module+value).Val() == 0 {
				continue
			}
			lenth := setting.RedisClient.LLen("selection" + module + value).Val()
			for _, content := range setting.RedisClient.LRange("selection"+module+value, 0, lenth).Val() {
				by = []byte(content)
				err = json.Unmarshal(by, &resp[index])
				index++
			}
		}
		//第一次查询做缓存,与分页
		Cache("home"+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			return resp, nil
		}
	} else {
		lenth := setting.RedisClient.LLen("selection" + module + tag.Label).Val()
		resp = make([]SelectionDetails, lenth)
		for index < int(lenth) {
			for _, content := range setting.RedisClient.LRange("selection"+module+tag.Label, 0, lenth).Val() {
				by := []byte(content)
				err = json.Unmarshal(by, &resp[index])
				if err != nil {
					return nil, err
				}
				index++
			}
		}
		Cache("home"+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			return resp, err
		}

	}

}

type CoverDetails struct {
	Nickname  string `json:"nickname"`
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
	File      string `json:"file"`
	Likes     int    `json:"likes"`
}

func GetCovers(module string, id int, tag Tags) (interface{}, error) {
	index := 0
	var resp []CoverDetails
	var err error
	if tag.Label == "recommend" {
		var hobby []string
		by, err := setting.RedisClient.HGet("hobby", strconv.Itoa(id)).Bytes()
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(by, &hobby) //解析json
		if err != nil {
			panic(err)
		}
		if hobby == nil {
			hobby = []string{"all"}
		}
		var size int
		for _, value := range hobby {
			lenth := setting.RedisClient.LLen("cover" + module + value).Val()
			size += int(lenth)
		}
		resp = make([]CoverDetails, size)
		for _, value := range hobby {
			if setting.RedisClient.Exists("cover"+module+value).Val() == 0 {
				continue
			}
			lenth := setting.RedisClient.LLen("cover" + module + value).Val()
			for _, content := range setting.RedisClient.LRange("cover"+module+value, 0, lenth).Val() {
				by = []byte(content)
				err = json.Unmarshal(by, &resp[index])
				index++
			}
		}
		//第一次查询做缓存,与分页
		Cache("home"+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			for i, _ := range resp {
				setting.DB.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			for i, _ := range resp {
				setting.DB.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, nil
		}
	} else {
		lenth := setting.RedisClient.LLen("cover" + module + tag.Label).Val()
		resp = make([]CoverDetails, lenth)
		for index < int(lenth) {
			for _, content := range setting.RedisClient.LRange("cover"+module+tag.Label, 0, lenth).Val() {
				by := []byte(content)
				err = json.Unmarshal(by, &resp[index])
				if err != nil {
					return nil, err
				}
				index++
			}
		}
		Cache("home"+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			for i, _ := range resp {
				setting.DB.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			for i, _ := range resp {
				setting.DB.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, err
		}

	}

}

//治愈系对应的录音
func CreateRecord(module int, id string, file string, uid int) (CoverDetails, error) {
	intId, _ := strconv.Atoi(id)
	selectionId := intId
	db := setting.MysqlConn()
	userId := uid
	var selection statements.Selection
	result1 := db.Model(&statements.Selection{}).Where("id=?", selectionId).First(&selection)
	if result1.Error != nil {
		return CoverDetails{}, errors.New("selection_id is invalid")
	}
	var cover statements.Cover
	cover.SelectionId = strconv.Itoa(selectionId)
	cover.UserId = userId
	cover.SongName = selection.SongName
	cover.File = file
	cover.Style = selection.Style
	cover.Language = selection.Language
	cover.Module = module
	coverDetails := CoverDetails{}
	err := db.Model(&statements.Cover{}).Create(&cover).Error
	setting.DB.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	coverDetails.CreatedAt = tools.DecodeTime(cover.CreatedAt)
	value, err := json.Marshal(coverDetails)
	if err != nil {
		return coverDetails, err
	}
	if cover.Style != "" {
		setting.RedisClient.RPush("cover"+strconv.Itoa(module)+cover.Style, string(value))
	}
	if cover.Language != "" {
		setting.RedisClient.RPush("cover"+strconv.Itoa(module)+cover.Language, string(value))
	}
	setting.RedisClient.RPush("cover"+strconv.Itoa(module)+"all", string(value))

	return coverDetails, err
}

func Select(selection statements.Selection) (SelectionDetails, error) {
	err := setting.DB.Table("selection").Create(&selection).Error
	selectionDetails := SelectionDetails{}
	err = setting.DB.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at").Where("selection.id=?", selection.ID).Joins("left join selection on user.id=selection.user_id").Scan(&selectionDetails).Error
	selectionDetails.CreatedAt = tools.DecodeTime(selection.CreatedAt)
	value, err := json.Marshal(selectionDetails)
	if err != nil {
		return selectionDetails, err
	}
	if selection.Style != "" {
		setting.RedisClient.RPush("selection"+strconv.Itoa(selection.Module)+selection.Style, string(value))
	}
	if selection.Language != "" {
		setting.RedisClient.RPush("selection"+strconv.Itoa(selection.Module)+selection.Language, string(value))
	}
	setting.RedisClient.RPush("selection"+strconv.Itoa(selection.Module)+"all", string(value))

	return selectionDetails, err
}

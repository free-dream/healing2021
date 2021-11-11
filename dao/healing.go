package dao

import (
	"encoding/json"
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
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
	db := setting.MysqlConn()
	userMsg := UsrMsg{}
	resp := make(map[string]interface{})
	db.Table("selection").Select("selection.id,selection.song_name,selection.style,selection.created_at,selection.remark,user.nickname").Joins("left join user on user.id=selection.user_id").Where("selection.id=?", selectionId).Scan(&userMsg)
	/*if err!=nil{
		panic(err)
		return nil, err
	}*/
	rows, err := db.Table("cover").Select("cover.user_id,user.nickname,cover.id").Joins("left join user on user.id=cover.user_id").Where("cover.selection_id=?", selectionId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	obj := CovMsg{}
	content := make(map[int]interface{})
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		db.Table("praise").Where("cover_id=? and is_liked", obj.ID, 0).Count(&obj.Likes)
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
	db := setting.MysqlConn()

	rows, err := db.Table("advertisement").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	var obj interface{}
	content := make(map[int]interface{})
	for rows.Next() {
		db.ScanRows(rows, &obj)
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
	Remark    string `json:"remark"`
}

//为分页器做缓存
//缓存首次请求的结果,非实时,不可更新
//且由于每次分页资源不同所以key可相同
//例首页:home
func Cache(key string, resp interface{}) {
	redisCli := setting.RedisConn()
	value, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	err = redisCli.HSet(key, "cache", value).Err()

}

//分页器
//从分页器里面取出
func Pager(key string, page int) (interface{}, error, int) {
	redisCli := setting.RedisConn()
	var resp []interface{}
	by, err := redisCli.HGet(key, "cache").Bytes()
	if err != nil {
		return nil, err, 0
	}
	err = json.Unmarshal(by, &resp)
	if err != nil {
		return nil, err, 0
	}
	var pageNum int
	if len(resp)%10 == 0 {
		pageNum = len(resp) / 10
	} else {
		pageNum = len(resp)/10 + 1
	}
	if len(resp)/10 > page {

		return resp[(page-1)*10 : (page-1)*10+10], nil, pageNum
	} else {

		return resp[(page-1)*10:], nil, pageNum
	}

}

//获取点歌页
//module1表示治愈系，2表示童年，但返回参数可能不同
//可参考
func GetSelections(id int, tag Tags) (interface{}, error) {
	redisCli := setting.RedisConn()
	index := 0
	var resp []SelectionDetails
	var err error
	if tag.Label == "recommend" {
		var hobby []string
		by, err1 := redisCli.HGet("healing2021:hobby", strconv.Itoa(id)).Bytes()
		if err1 != nil {
			return nil, err1
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
			lenth := redisCli.LLen("healing2021:selection." + value).Val()
			size += int(lenth)
		}
		resp = make([]SelectionDetails, size)
		for _, value := range hobby {
			if redisCli.Exists("healing2021:selection."+value).Val() == 0 {
				continue
			}
			lenth := redisCli.LLen("healing2021:selection." + value).Val()
			for _, content := range redisCli.LRange("healing2021:selection."+value, 0, lenth).Val() {
				by = []byte(content)
				err = json.Unmarshal(by, &resp[index])
				index++
			}
		}
		//第一次查询做缓存,与分页
		Cache("healing2021:home."+strconv.Itoa(id), resp)
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
		lenth := redisCli.LLen("healing2021:selection." + tag.Label).Val()
		resp = make([]SelectionDetails, lenth)
		for index < int(lenth) {
			for _, content := range redisCli.LRange("healing2021:selection."+tag.Label, 0, lenth).Val() {
				by := []byte(content)
				err = json.Unmarshal(by, &resp[index])
				if err != nil {
					return nil, err
				}
				index++
			}
		}
		Cache("healing2021:home."+strconv.Itoa(id), resp)
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
	db := setting.MysqlConn()

	redisCli := setting.RedisConn()
	index := 0
	var resp []CoverDetails

	if tag.Label == "recommend" {
		var hobby []string
		by, err := redisCli.HGet("healing2021:hobby", strconv.Itoa(id)).Bytes()
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
			lenth := redisCli.LLen("healing2021:cover." + module + "." + value).Val()
			size += int(lenth)
		}
		resp = make([]CoverDetails, size)
		for _, value := range hobby {
			if redisCli.Exists("healing2021:cover."+module+"."+value).Val() == 0 {
				continue
			}
			lenth := redisCli.LLen("healing2021:cover." + module + "." + value).Val()
			for _, content := range redisCli.LRange("healing2021:cover."+module+"."+value, 0, lenth).Val() {
				by = []byte(content)
				err = json.Unmarshal(by, &resp[index])
				index++
			}
		}
		//第一次查询做缓存,与分页
		Cache("healing2021:home."+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			for i, _ := range resp {
				db.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			for i, _ := range resp {
				db.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, nil
		}
	} else {
		lenth := redisCli.LLen("healing2021:cover." + module + "." + tag.Label).Val()
		resp = make([]CoverDetails, lenth)
		for index < int(lenth) {
			for _, content := range redisCli.LRange("healing2021:cover."+module+"."+tag.Label, 0, lenth).Val() {
				by := []byte(content)
				err := json.Unmarshal(by, &resp[index])
				if err != nil {
					return nil, err
				}
				index++
			}
		}
		Cache("healing2021:home."+strconv.Itoa(id), resp)
		if len(resp) > 10 {
			resp = resp[0:10]
		}
		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			for i, _ := range resp {
				db.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, nil
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			for i, _ := range resp {
				db.Table("praise").Where("cover_id=? and is_liked", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			return resp, nil
		}

	}

}

//治愈系对应的录音
func CreateRecord(module int, id string, file string, uid int, isAnon bool) (int, CoverDetails, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	intId, _ := strconv.Atoi(id)
	selectionId := intId
	userId := uid
	var selection statements.Selection
	result1 := db.Model(&statements.Selection{}).Where("id=?", selectionId).First(&selection)
	if result1.Error != nil {
		return 0, CoverDetails{}, errors.New("selection_id is invalid")
	}
	var cover statements.Cover
	cover.SelectionId = strconv.Itoa(selectionId)
	cover.UserId = userId
	cover.SongName = selection.SongName
	cover.File = file
	cover.Style = selection.Style
	cover.Language = selection.Language
	cover.Module = module
	cover.IsAnon = isAnon
	coverDetails := CoverDetails{}

	err := db.Model(&statements.Cover{}).Create(&cover).Error
	db.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at,cover.likes").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	user_id := 0
	db.Table("selection").Select("user_id").Where("cover_id=?", coverDetails.ID).Scan(&user_id)
	if !isAnon {
		coverDetails.CreatedAt = tools.DecodeTime(cover.CreatedAt)
		value, err1 := json.Marshal(coverDetails)
		if err1 != nil {
			return 0, coverDetails, err1
		}
		if cover.Style != "" {
			redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+cover.Style, string(value))
		}
		if cover.Language != "" {
			redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+cover.Language, string(value))
		}
		redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+"all", string(value))
	}

	return user_id, coverDetails, err
}

func Select(selection statements.Selection) (SelectionDetails, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	err := db.Table("selection").Create(&selection).Error
	selectionDetails := SelectionDetails{}
	err = db.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at,remark").Where("selection.id=?", selection.ID).Joins("left join selection on user.id=selection.user_id").Scan(&selectionDetails).Error
	selectionDetails.CreatedAt = tools.DecodeTime(selection.CreatedAt)
	value, err := json.Marshal(selectionDetails)
	if err != nil {
		return selectionDetails, err
	}
	if selection.Style != "" {
		redisCli.RPush("healing2021:selection"+"."+selection.Style, string(value))
	}
	if selection.Language != "" {
		redisCli.RPush("healing2021:selection"+"."+selection.Language, string(value))
	}
	redisCli.RPush("healing2021:selection"+"."+"all", string(value))

	return selectionDetails, err
}

type DevMsg struct {
	ID       int    `json:"id"`
	SongName string `json:"song_name"`
	Singer   string `json:"singer"`
	File     string `json:"file"`
	Likes    int    `json:"likes"`
}

func PlayDevotion() (map[string]interface{}, error) {
	resp := make(map[string]interface{})
	content := make(map[int]interface{})
	content2 := make(map[int]interface{})
	likes := 0
	index := 0
	db := setting.MysqlConn()
	devotion := DevMsg{}
	rows, err := db.Table("devotion").Where("singer=?", "阿细").Rows()
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &devotion)
		db.Table("praise").Where("devotion_id=?", devotion.ID).Count(&likes)
		devotion.Likes = likes
		content[index] = devotion
		index++
	}
	resp["阿细"] = content
	index = 0
	rows, err = db.Table("devotion").Where("singer=?", "梁山山").Rows()
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		db.ScanRows(rows, &devotion)
		if err != nil {
			return resp, err
		}
		db.Table("praise").Where("devotion_id=?", devotion.ID).Count(&likes)
		devotion.Likes = likes
		content2[index] = devotion
		index++
	}
	resp["梁山山"] = content2
	return resp, err
}

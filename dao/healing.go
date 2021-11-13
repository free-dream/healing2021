package dao

import (
	"encoding/json"
	"errors"
	"log"
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
	Avatar    string `json:"avatar"`
}
type CovMsg struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Likes    int    `json:"likes"`
	Nickname string `json:"nickname"`
	File     string `json:"file"`
	Check    int    `json:"check"`
	Avatar   string `json:"avatar"`
}

//处理治愈详情页
//点赞数debug，尚未测试
//结构体疑似有bug
func GetHealingPage(selectionId int, userId int) (interface{}, error) {
	db := setting.MysqlConn()
	userMsg := UsrMsg{}
	resp := make(map[string]interface{})
	db.Table("selection").Select("user.avatar,selection.id,selection.song_name,selection.style,selection.created_at,selection.remark,user.nickname").Joins("left join user on user.id=selection.user_id").Where("selection.id=?", selectionId).Scan(&userMsg)
	/*if err!=nil{
		panic(err)
		return nil, err
	}*/
	rows, err := db.Table("cover").Select("user.avatar,cover.file,cover.user_id,user.nickname,cover.id").Joins("left join user on user.id=cover.user_id").Where("cover.selection_id=?", selectionId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	index := 0
	obj := CovMsg{}
	content := make(map[int]interface{})
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		db.Table("praise").Where("cover_id=? and is_liked=?", obj.ID, 1).Count(&obj.Likes)
		//插入点赞确认
		check, err1 := PackageCheckMysql(userId, "cover", obj.ID)
		if err1 != nil {
			log.Printf(err1.Error())
			obj.Check = 0
		} else if check {
			obj.Check = 1
		} else {
			obj.Check = 0
		}
		//
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
func Pager(key string, page int) (interface{}, error) {
	redisCli := setting.RedisConn()
	var resp []interface{}
	by, err := redisCli.HGet(key, "cache").Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(by, &resp)
	if err != nil {
		return nil, err
	}
	var pageNum int
	if len(resp)%10 == 0 {
		pageNum = len(resp) / 10
		if pageNum >= page {

			return resp[(page-1)*10 : (page-1)*10+10], nil
		} else {

			return nil, errors.New("out of range")
		}

	} else {
		pageNum = len(resp)/10 + 1
		if pageNum > page {

			return resp[(page-1)*10 : (page-1)*10+10], nil
		} else if pageNum == page {
			return resp[(page-1)*10:], nil
		} else {

			return nil, errors.New("out of range")

		}
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

		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
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

		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
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
	Check     int    `json:"check"`
}

//传入userid以确认
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
			//确认是否点赞
			for i, _ := range resp {
				boolean, err := PackageCheckMysql(id, "cover", resp[i].ID)
				if err != nil {
					log.Printf(err.Error())
					resp[i].Check = 0
				} else if boolean {
					resp[i].Check = 1
				} else {
					resp[i].Check = 0
				}
			}
			//
		}
		//第一次查询做缓存,与分页

		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			for i, _ := range resp {
				//确认是否点赞
				boolean, err := PackageCheckMysql(id, "cover", resp[i].ID)
				if err != nil {
					log.Printf(err.Error())
					resp[i].Check = 0
				} else if boolean {
					resp[i].Check = 1
				} else {
					resp[i].Check = 0
				}
				//
				db.Table("praise").Where("cover_id=? and is_liked=?", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, err
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt < resp[j].CreatedAt
			})
			for i, _ := range resp {
				//确认是否点赞
				boolean, err := PackageCheckMysql(id, "cover", resp[i].ID)
				if err != nil {
					log.Printf(err.Error())
					resp[i].Check = 0
				} else if boolean {
					resp[i].Check = 1
				} else {
					resp[i].Check = 0
				}
				//
				db.Table("praise").Where("cover_id=? and is_liked=?", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
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
				//确认是否点赞
				boolean, err := PackageCheckMysql(id, "cover", resp[i].ID)
				if err != nil {
					log.Printf(err.Error())
					resp[i].Check = 0
				} else if boolean {
					resp[i].Check = 1
				} else {
					resp[i].Check = 0
				}
				//
				db.Table("praise").Where("cover_id=? and is_liked=?", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, nil
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt < resp[j].CreatedAt
			})
			for i, _ := range resp {
				//确认是否点赞
				boolean, err := PackageCheckMysql(id, "cover", resp[i].ID)
				if err != nil {
					log.Printf(err.Error())
					resp[i].Check = 0
				} else if boolean {
					resp[i].Check = 1
				} else {
					resp[i].Check = 0
				}
				//
				db.Table("praise").Where("cover_id=? and is_liked=?", resp[i].ID, 1).Count(&resp[i].Likes)
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, nil
		}

	}

}

//治愈系对应的录音
func CreateRecord(module int, selectionId int, file string, uid int, isAnon bool) (int, CoverDetails, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()

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
	db.Table("user").Select("cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	if !isAnon {
		coverDetails.CreatedAt = tools.DecodeTime(cover.CreatedAt)
		value, err1 := json.Marshal(coverDetails)
		if err1 != nil {
			return 0, coverDetails, err1
		}

		if cover.Style == cover.Language && cover.Style != "" {
			redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+cover.Style, string(value))
		} else {
			if cover.Style != "" {
				redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+cover.Style, string(value))
			}
			if cover.Language != "" {
				redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+cover.Language, string(value))
			}
		}

		redisCli.RPush("healing2021:cover."+strconv.Itoa(module)+"."+"all", string(value))
	}

	return selection.UserId, coverDetails, err
}

func Select(selection statements.Selection) (int, SelectionDetails, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	user := statements.User{}
	db.Table("user").Select("selection_num").Where("id=?", selection.UserId).Scan(&user)
	if 0 < user.SelectionNum {
		err := db.Table("selection").Create(&selection).Error
		db.Table("user").Where("id=?", selection.UserId).Update("selection_num", user.SelectionNum-1)
		selectionDetails := SelectionDetails{}
		err = db.Table("user").Select("selection.user_id,selection.id,user.nickname,user.avatar,selection.song_name,selection.created_at,remark").Where("selection.id=?", selection.ID).Joins("left join selection on user.id=selection.user_id").Scan(&selectionDetails).Error
		selectionDetails.CreatedAt = tools.DecodeTime(selection.CreatedAt)
		value, err := json.Marshal(selectionDetails)
		if err != nil {
			return user.SelectionNum, selectionDetails, err
		}
		if selection.Style == selection.Language && selection.Style != "" {
			redisCli.RPush("healing2021:selection"+"."+selection.Style, string(value))
		} else {
			if selection.Style != "" {
				redisCli.RPush("healing2021:selection"+"."+selection.Style, string(value))
			}
			if selection.Language != "" {
				redisCli.RPush("healing2021:selection"+"."+selection.Language, string(value))
			}
		}
		redisCli.RPush("healing2021:selection"+"."+"all", string(value))
		return user.SelectionNum, selectionDetails, err
	} else {
		return user.SelectionNum, SelectionDetails{}, errors.New("今日次数已用尽")
	}

}

type DevMsg struct {
	ID       int    `json:"id"`
	SongName string `json:"song_name"`
	Singer   string `json:"singer"`
	File     string `json:"file"`
	Likes    int    `json:"likes"`
	Check    int    `json:"check"`
}

func PlayDevotion(userid int) (map[string]interface{}, error) {
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
		//插入点赞确认
		check, err1 := PackageCheckMysql(userid, "cover", devotion.ID)
		if err1 != nil {
			log.Printf(err1.Error())
			devotion.Check = 0
		} else if check {
			devotion.Check = 1
		} else {
			devotion.Check = 0
		}
		//
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
		//插入点赞确认
		check, err1 := PackageCheckMysql(userid, "cover", devotion.ID)
		if err1 != nil {
			log.Printf(err1.Error())
			devotion.Check = 0
		} else if check {
			devotion.Check = 1
		} else {
			devotion.Check = 0
		}
		//
		db.Table("praise").Where("devotion_id=?", devotion.ID).Count(&likes)
		devotion.Likes = likes
		content2[index] = devotion
		index++
	}
	resp["梁山山"] = content2
	return resp, err
}

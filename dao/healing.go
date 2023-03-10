package dao

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

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
	Avatar    string `json:"avatar"`
	UserId    int    `json:"user_id"`
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
func GetHealingPage(selectionId int, userId int) (interface{}, interface{}) {
	db := setting.MysqlConn()
	userMsg := UsrMsg{}
	obj := []CoverDetails{}
	db.Table("selection").Select("selection.user_id,user.avatar,selection.id,selection.song_name,selection.style,selection.created_at,selection.remark,user.nickname").Joins("left join user on user.id=selection.user_id").Where("selection.id=?", selectionId).Scan(&userMsg)
	db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id=" + strconv.Itoa(selectionId) + ")" + " as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
		Scan(&obj)
	ch := make(chan int, 15)
	for i, _ := range obj {
		//确认是否点赞
		ViolenceGetLikeheckC(userId, obj[i], ch)
		obj[i].Check = <-ch

	}
	return userMsg, obj
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
	//Module  string `json:"module"`
}

const LanguageConst = "国语 粤语 日语 英语"

const StyleConst = "流行 古风 民谣 摇滚 RAP ACG 其他"

type SelectionDetails struct {
	Nickname    string `json:"nickname"`
	ID          int    `json:"id"`
	SongName    string `json:"song_name"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	Avatar      string `json:"avatar"`
	Remark      string `json:"remark"`
	Sex         int    `json:"sex"`
	PhoneNumber string `json:"phone_number"`
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
//可参考
func GetSelections(id int, tag Tags) (interface{}, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	var resp []SelectionDetails
	VTable := db.Table("selection").Select("user.phone_number,user.sex,user.avatar,user.nickname,selection.id,selection.song_name,selection.user_id,selection.id,selection.created_at,selection.remark ").
		Joins("inner join user on user.id=selection.user_id").
		Order("selection.created_at desc")
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
			VTable.Scan(&resp)
		} else {
			for i := range hobby {
				hobby[i] = "'" + hobby[i] + "'"
			}
			hobbyArr := strings.Join(hobby, ",")
			VTable.
				Where("selection.language in " + "(" + hobbyArr + ")" + " or selection.style in " + "(" + hobbyArr + ")").
				Scan(&resp)
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
		if ok := strings.Contains(LanguageConst, tag.Label); ok {
			VTable.
				Where("selection.language=?", tag.Label).
				Scan(&resp)
		} else if ok = strings.Contains(StyleConst, tag.Label); ok {
			VTable.
				Where("selection.style=?", tag.Label).
				Scan(&resp)
		} else {
			VTable.Scan(&resp)
		}

		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, nil
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

	}

}

type CoverDetails struct {
	Nickname    string `json:"nickname"`
	ID          int    `json:"id"`
	SongName    string `json:"song_name"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	Avatar      string `json:"avatar"`
	File        string `json:"file"`
	Likes       int    `json:"likes"`
	Check       int    `json:"check"`
	SelectionId int    `json:"selection_id"`
	ClassicId   int    `json:"classic_id"`
	Module      int    `json:"module"`
}

//简单的翻唱对象
type LikeObj struct {
	Likes   int `json:"likes"`
	CoverId int `json:"cover_id"`
}

//传入userid以确认
func GetCovers(id int, tag Tags) (interface{}, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
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
			db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id<>0) as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
				Scan(&resp).Scan(&resp)
		} else {
			for i := range hobby {
				hobby[i] = "'" + hobby[i] + "'"
			}
			hobbyArr := strings.Join(hobby, ",")
			db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id<>0 AND cover.language in " + "(" + hobbyArr + ")" + " or cover.style in " + "(" + hobbyArr + ")" + ") as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
				Scan(&resp)
		}

		//第一次查询做缓存,与分页

		if tag.RankWay == 1 {
			rand.Seed(time.Now().Unix())
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			ch := make(chan int, 15)
			for i, _ := range resp {
				//确认是否点赞
				ViolenceGetLikeheckC(id, resp[i], ch)
				resp[i].Check = <-ch
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, err
		} else {
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, nil
		}
	} else {
		if ok := strings.Contains(LanguageConst, tag.Label); ok {
			db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id<>0 AND cover.language=" + "'" + tag.Label + "'" + ")" + " as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
				Scan(&resp)
		} else if ok = strings.Contains(StyleConst, tag.Label); ok {
			db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id<>0 AND cover.style=" + "'" + tag.Label + "'" + ")" + " as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
				Scan(&resp)
		} else {
			db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.selection_id<>0) as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
				Scan(&resp)
		}
		if tag.RankWay == 1 {
			//采用rand.Shuffle，将切片随机化处理后返回
			rand.Seed(time.Now().Unix())
			rand.Shuffle(len(resp), func(i, j int) { resp[i], resp[j] = resp[j], resp[i] })
			ch := make(chan int, 15)
			for i, _ := range resp {
				//确认是否点赞
				ViolenceGetLikeheckC(id, resp[i], ch)
				resp[i].Check = <-ch
			}
			Cache("healing2021:home."+strconv.Itoa(id), resp)
			if len(resp) > 10 {
				resp = resp[0:10]
			}
			return resp, nil
		} else {
			sort.Slice(resp, func(i, j int) bool {
				return resp[i].CreatedAt > resp[j].CreatedAt
			})
			ch := make(chan int, 15)
			for i, _ := range resp {
				//确认是否点赞
				ViolenceGetLikeheckC(id, resp[i], ch)
				resp[i].Check = <-ch
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
func CreateRecord(module int, selectionId int, file string, uid int, isAnon bool) (*statements.Selection, *CoverDetails, error) {
	db := setting.MysqlConn()
	userId := uid
	var cover statements.Cover
	var selection statements.Selection
	var classic statements.Classic
	switch module {
	case 1:
		cover.SelectionId = selectionId
		db.Model(&statements.Selection{}).Where("id=?", selectionId).Scan(&selection)
		cover.Style = selection.Style
		cover.Language = selection.Language
		cover.SongName = selection.SongName
	case 2:
		cover.ClassicId = selectionId
		db.Model(&statements.Classic{}).Where("id=?", selectionId).Scan(&classic)
		cover.SongName = classic.SongName
	}

	//

	//补一个拿人头和名字的
	data, err1 := GetUserInfo(uid)
	if err1 != nil {
		log.Printf(err1.Error())
	}
	//
	cover.Avatar = data[0]
	cover.UserId = userId
	cover.File = file
	cover.Module = module
	cover.IsAnon = isAnon
	cover.Nickname = data[1]
	coverDetails := CoverDetails{}
	err := db.Model(&statements.Cover{}).Create(&cover).Error
	switch module {
	case 1:
		db.Table("user").Select("cover.selection_id,cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	case 2:
		db.Table("user").Select("cover.classic_id,cover.file,cover.user_id,cover.id,user.nickname,user.avatar,cover.song_name,cover.created_at").Where("cover.id=?", cover.ID).Joins("left join cover on user.id=cover.user_id").Scan(&coverDetails)
	}
	return &selection, &coverDetails, err
}

func Select(selection statements.Selection, avatar string, nickname string) (int, int, error) {
	db := setting.MysqlConn()
	user := statements.User{}

	db.Table("user").Select("selection_num").Where("id=?", selection.UserId).Scan(&user)
	if 0 < user.SelectionNum {
		selection.Nickname = nickname
		selection.Avatar = avatar
		err := db.Table("selection").Create(&selection).Error
		db.Table("user").Where("id=?", selection.UserId).Update("selection_num", user.SelectionNum-1)
		return int(selection.ID), user.SelectionNum, err
	} else {
		return int(selection.ID), user.SelectionNum, errors.New("今日次数已用尽")
	}

}

type DevMsg struct {
	ID       int    `json:"id"`
	SongName string `json:"song_name"`
	Singer   string `json:"singer"`
	File     string `json:"file"`
}

func PlayDevotion() interface{} {
	db := setting.MysqlConn()
	devotion := []DevMsg{}
	devotion2 := []DevMsg{}
	resp := map[string][]DevMsg{}
	db.Table("devotion").Select("id,song_name,file").
		Where("singer='阿细'").
		Scan(&devotion)
	resp["阿细"] = devotion
	db.Table("devotion").Select("id,song_name,file").
		Where("singer='梁山山'").
		Scan(&devotion2)

	resp["梁山山"] = devotion2
	return resp
}

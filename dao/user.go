package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"math/rand"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

/**
 * @Description 通过用户 id 返回该用户的所有信息
 * @Param 用户 id
 * @return 含有该用户的所有信息的结构体，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetUserById(Id int) (statements.User, bool) {
	MysqlDB := setting.MysqlConn()
	OneUser := statements.User{}
	fmt.Println("reach")
	if err := MysqlDB.Where("id=?", Id).First(&OneUser).Error; err != nil {
		return OneUser, false
	}
	fmt.Println("reach")
	return OneUser, true
}

type User struct {
	ID             uint
	Openid         string   `json:"openid"`
	Nickname       string   `json:"nickname"`
	RealName       string   `json:"real_name"`
	PhoneNumber    string   `json:"phone_number"`
	Sex            int      `json:"sex"`
	School         string   `json:"school"`
	Avatar         string   `json:"avatar"`
	AvatarVisible  int      `json:"avatar_visible"`
	PhoneSearch    int      `json:"phone_search"`
	RealNameSearch int      `json:"real_name_search"`
	Signature      string   `json:"signature"`
	Hobby          []string `json:"hobby"`
}

func FakeCreateUser(user *statements.User) (int, error) {
	err := setting.DB.Table("user").Where("nickname=?", user.Nickname).Scan(&user).Error
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil

}

func CreateUser(param *User) (int, error) {
	count := 0
	user := statements.User{
		Nickname:    param.Nickname,
		RealName:    param.RealName,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		School:      param.School,
	}
	setting.DB.Table("user").Where("nickname=?", user.Nickname).Count(&count)
	if count != 0 {
		return 0, errors.New("error")
	}
	setting.DB.Table("user").Create(&user)
	value, err := json.Marshal(param.Hobby)
	if err != nil {
		panic(err)
		// return 0, err
	}
	setting.RedisClient.HSet("hobby", strconv.Itoa(int(user.ID)), value)

	return int(user.ID), nil

}

func UpdateUser(user *User, id int, avatar string) (string, error) {
	other := statements.User{}
	setting.DB.Table("user").Where("nickname=?", user.Nickname).Scan(&other)
	if int(other.ID) != id && other.ID != 0 {

		return "", errors.New("error")
	}
	message := make(map[string]interface{})
	message["nickname"] = user.Nickname
	message["avatar_visible"] = user.AvatarVisible
	if user.AvatarVisible == 1 {
		avatar = tools.GetAvatarUrl(rand.Intn(2))
	}
	message["avatar"] = avatar
	message["phone_search"] = user.PhoneSearch
	message["real_name_search"] = user.RealNameSearch
	message["signature"] = user.Signature
	err := setting.DB.Table("user").Select("nickname,signature,avatar_visible,phone_search,real_name_search ").Where("id=?", id).Update(message).Error
	if err != nil {
		return "", err
	}

	return avatar, nil
}

type UserMsg struct {
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	ID        int    `json:"id"`
	School    string `json:"school"`
	Signature string `json:"signature"`
}
type SelectionMsg struct {
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	CreatedAt string `json:"created_at"`
}
type SelectionMsgV2 struct {
	ID        int       `json:"id"`
	SongName  string    `json:"song_name"`
	CreatedAt time.Time `json:"created_at"`
}
type CoverMsg struct {
	ID        int    `json:"id"`
	SongName  string `json:"song_name"`
	CreatedAt string `json:"created_at"`
	Likes     int    `json:"likes"`
}
type CoverMsgV2 struct {
	ID        int       `json:"id"`
	SongName  string    `json:"song_name"`
	CreatedAt time.Time `json:"created_at"`
}
type PraiseMsg struct {
	SongName  string `json:"song_name"`
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	CoverId   int    `json:"cover_id"`
	Likes     int    `json:"likes"`
}
type PraiseMsgV2 struct {
	CoverId   int       `json:"cover_id"`
	SongName  string    `json:"song_name"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
}
type MomentMsg struct {
	SongName  string   `json:"song_name"`
	CreatedAt string   `json:"created_at"`
	ID        int      `json:"id"`
	State     []string `json:"state"`
	Content   string   `json:"content"`
	Likes     int      `json:"likes"`
}
type MomentMsgV2 struct {
	SongName  string    `json:"song_name"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	State     string    `json:"state"`
	Content   string    `json:"content"`
}

func GetUser(id int) interface{} {
	user := UserMsg{}
	resp := make(map[string]interface{})
	setting.DB.Table("user").Select("id,avatar,nickname,school,signature").Where("id=?", id).Scan(&user)
	resp["message"] = user
	resp["mySelections"] = getSelections(user.ID, "selection", "user_id=?")
	resp["mySongs"] = getCovers(user.ID, "cover", "user_id=?")
	resp["moments"] = getMoments(user.ID, "moment", "user_id=?")
	resp["myLikes"] = getPraises(user.ID, "praise", "praise.user_id=?")

	return resp
}

func UpdateBackground(openid string, background string) error {
	err := setting.DB.Table("user").Where("openid=?", openid).Update("background", background).Error
	return err
}

func GetCallee(id int) interface{} {
	user := UserMsg{}
	resp := make(map[string]interface{})
	setting.DB.Table("user").Where("id=?", id).Scan(&user)
	resp["message"] = user
	resp["mySelections"] = getSelections(user.ID, "selection", "user_id=?")
	resp["mySongs"] = getCovers(user.ID, "cover", "user_id=?")
	resp["moments"] = getMoments(user.ID, "moment", "user_id=?")
	resp["myLikes"] = getPraises(user.ID, "praise", "praise.user_id=?")

	return resp

}

//返回所有信息采用扫描行的方式

func getPraises(value interface{}, tableName string, condition string) interface{} {
	obj := PraiseMsgV2{}
	resp := PraiseMsg{}
	rows, err := setting.DB.Table(tableName).Joins("left join cover on cover.id=praise.cover_id").Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.ID = obj.ID
		resp.CoverId = obj.CoverId
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		setting.DB.Table("praise").Where("cover_id=? and is_liked=?", resp.CoverId, 1).Count(&resp.Likes)
		content[index] = resp
		index++
	}
	return content
}
func getCovers(value interface{}, tableName string, condition string) interface{} {
	obj := CoverMsgV2{}
	resp := CoverMsg{}
	rows, err := setting.DB.Table(tableName).Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.ID = obj.ID
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		setting.DB.Table("praise").Where("cover_id=? and is_liked=?", resp.ID, 1).Count(&resp.Likes)
		content[index] = resp
		index++
	}
	return content
}
func getSelections(value interface{}, tableName string, condition string) interface{} {
	obj := SelectionMsgV2{}
	resp := SelectionMsg{}
	rows, err := setting.DB.Table(tableName).Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		content[index] = resp
		index++
	}
	return content
}

func getMoments(value interface{}, tableName string, condition string) interface{} {
	obj := MomentMsgV2{}
	resp := MomentMsg{}
	rows, err := setting.DB.Table(tableName).Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = setting.DB.ScanRows(rows, &obj)
		resp.State = tools.DecodeStrArr(obj.State)
		resp.ID = int(obj.ID)
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		resp.Content = obj.Content
		setting.DB.Table("praise").Where("cover_id=? and is_liked=?", resp.ID, 1).Count(&resp.Likes)
		if err != nil {
			panic(err)
		}
		content[index] = resp
		index++
	}
	return content
}

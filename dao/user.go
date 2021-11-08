package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

/**
 * @Description 通过用户 id 返回该用户的所有信息
 * @Param 用户 id
 * @return 含有该用户的所有信息的结构体，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetUserById(Id int) (statements.User, bool) {
	MysqlDB := setting.MysqlConn()
	OneUser := statements.User{}
	if err := MysqlDB.Where("id=?", Id).First(&OneUser).Error; err != nil {
		return OneUser, false
	}
	return OneUser, true
}

/*type User struct {
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
}*/

func FakeCreateUser(user *statements.User) (int, error) {
	index := 0
	exUser := statements.User{}
	db := setting.MysqlConn()

	db.Table("user").Where("nickname=?", user.Nickname).Count(&index).Scan(&exUser)

	if index == 0 {
		db.Table("user").Create(&user)
		return int(user.ID), nil
	} else {
		if user.Openid == exUser.Openid {
			return int(exUser.ID), nil
		} else {
			//逻辑更改，返回已有用户
			return int(exUser.ID), errors.New("昵称已存在")
		}

	}

}
func CreateUser(user statements.User) (int, int) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	if !redisCli.SIsMember("healing2021:openid", user.Openid).Val() {

		db.Table("user").Create(&user)
		return 0, int(user.ID)
	} else {
		db.Table("user").Where("openid=?", user.Openid).Scan(&user)
		return 1, int(user.ID)
	}

}
func GetPhoneNumber(id int) (error, int) {
	db := setting.MysqlConn()
	phone_number := 0
	err := db.Table("user").Where("id=? and phone_search=?", id, 0).Select("phone_number").Scan(&phone_number).Error
	if err != nil {
		return err, 0
	}
	return nil, phone_number
}

func RefineUser(param statements.User, id int) error {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	user := statements.User{
		Nickname:    param.Nickname,
		RealName:    param.RealName,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		School:      param.School,
	}

	db.Table("user").Where("id=?", id).Select("nickname,real_name,phone_number,sex,school").Update(&user)
	//value, _ := json.Marshal(param.Openid)
	redisCli.SAdd("healing2021:openid", param.Openid)
	return nil

}

type Hobby struct {
	Hobby []string `json:"hobby"`
}

func HobbyStore(hobby []string, id int) error {

	redisCli := setting.RedisConn()
	value, err := json.Marshal(hobby)
	if err != nil {
		return err
	}
	err = redisCli.HSet("healing2021:hobby.", strconv.Itoa(id), value).Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *statements.User, id int, avatar string) (string, error) {
	db := setting.MysqlConn()
	other := statements.User{}
	db.Table("user").Where("nickname=?", user.Nickname).Scan(&other)
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
	err := db.Table("user").Select("nickname,signature,avatar_visible,phone_search,real_name_search ").Where("id=?", id).Update(message).Error
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
	Sex       int    `json:"sex"`
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
	ID          int    `json:"id"`
	SelectionId int    `json:"selection_id"`
	SongName    string `json:"song_name"`
	CreatedAt   string `json:"created_at"`
	Likes       int    `json:"likes"`
}
type CoverMsgV2 struct {
	ID          int       `json:"id"`
	SelectionId int       `json:"selection_id"`
	SongName    string    `json:"song_name"`
	CreatedAt   time.Time `json:"created_at"`
}
type PraiseMsg struct {
	CoverId     int    `json:"cover_id"`
	SongName    string `json:"song_name"`
	CreatedAt   string `json:"created_at"`
	ID          int    `json:"id"`
	SelectionId int    `json:"selection_id"`
	Likes       int    `json:"likes"`
}
type PraiseMsgV2 struct {
	CoverId     int       `json:"cover_id"`
	SelectionId int       `json:"selection_id"`
	SongName    string    `json:"song_name"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int       `json:"id"`
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
	db := setting.MysqlConn()
	usr := UserMsg{}
	resp := make(map[string]interface{})
	db.Table("user").Select("id,avatar,nickname,school,signature,sex").Where("id=?", id).Scan(&usr)
	resp["message"] = usr
	resp["mySelections"] = getSelections(usr.ID, "selection", "user_id=?")
	resp["mySongs"] = getCovers("cover", "user_id=?", usr.ID, 0)
	resp["moments"] = getMoments(usr.ID, "moment", "user_id=?")
	resp["myLikes"] = getPraises(usr.ID, "praise", "praise.user_id=?")

	return resp
}

func UpdateBackground(openid string, background string) error {
	db := setting.MysqlConn()
	err := db.Table("user").Where("openid=?", openid).Update("background", background).Error
	return err
}

func GetCallee(id int) interface{} {
	db := setting.MysqlConn()
	usr := UserMsg{}
	resp := make(map[string]interface{})
	db.Table("user").Where("id=?", id).Scan(&usr)
	resp["message"] = usr
	resp["mySelections"] = getSelections(usr.ID, "selection", "user_id=?")

	resp["mySongs"] = getCovers("cover", "user_id=? and is_anon=?", usr.ID, 1)
	resp["moments"] = getMoments(usr.ID, "moment", "user_id=?")
	resp["myLikes"] = getPraises(usr.ID, "praise", "praise.user_id=?")

	return resp

}

//返回所有信息采用扫描行的方式

func getPraises(value interface{}, tableName string, condition string) interface{} {
	db := setting.MysqlConn()
	obj := PraiseMsgV2{}
	resp := PraiseMsg{}
	rows, err := db.Table(tableName).Joins("left join cover on cover.id=praise.cover_id").Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.ID = obj.ID
		resp.CoverId = obj.CoverId
		resp.SelectionId = obj.SelectionId
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		db.Table("praise").Where("cover_id=? and is_liked=?", resp.CoverId, 1).Count(&resp.Likes)
		content[index] = resp
		index++
	}
	return content
}
func getCovers(tableName string, condition string, value int, module int) interface{} {
	db := setting.MysqlConn()
	obj := CoverMsgV2{}
	resp := CoverMsg{}
	var rows *sql.Rows
	var err error
	if module == 0 {
		rows, err = db.Table(tableName).Where(condition, value).Rows()
	} else {
		rows, err = db.Table(tableName).Where(condition, value, 0).Rows()
	}

	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.ID = obj.ID
		resp.SelectionId = obj.SelectionId
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		db.Table("praise").Where("cover_id=? and is_liked=?", resp.ID, 1).Count(&resp.Likes)
		content[index] = resp
		index++
	}
	return content
}
func getSelections(value interface{}, tableName string, condition string) interface{} {
	db := setting.MysqlConn()
	obj := SelectionMsgV2{}
	resp := SelectionMsg{}
	rows, err := db.Table(tableName).Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		if err != nil {
			panic(err)
		}
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		resp.ID = obj.ID
		content[index] = resp
		index++
	}
	return content
}

func getMoments(value interface{}, tableName string, condition string) interface{} {
	db := setting.MysqlConn()
	obj := MomentMsgV2{}
	resp := MomentMsg{}
	rows, err := db.Table(tableName).Where(condition, value).Rows()
	index := 0
	content := make(map[int]interface{})
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &obj)
		resp.State = tools.DecodeStrArr(obj.State)
		resp.ID = obj.ID
		resp.CreatedAt = tools.DecodeTime(obj.CreatedAt)
		resp.SongName = obj.SongName
		resp.Content = obj.Content
		db.Table("praise").Where("cover_id=? and is_liked=?", resp.ID, 1).Count(&resp.Likes)
		if err != nil {
			panic(err)
		}
		content[index] = resp
		index++
	}
	return content
}

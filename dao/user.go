package dao

import (
	"encoding/json"
	"errors"
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
func CreateUser(user statements.User) int {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	redisCli.HSet("healing2021:avatar", user.Openid, user.Avatar)
	vuser := statements.User{}
	db.Table("user").Select("openid").Where("openid=?", user.Openid).Scan(&vuser)
	if vuser.Openid == "" {
		db.Table("user").Create(&user)
		return int(user.ID)
	} else {
		db.Table("user").Where("openid=?", user.Openid).Scan(&user)
		return int(user.ID)
	}

}
func Exist(openid string) bool {
	redisCli := setting.RedisConn()
	return redisCli.SIsMember("healing2021:openid", openid).Val()
}
func GetPhoneNumber(id int) (error, string) {
	db := setting.MysqlConn()
	user := statements.User{}
	err := db.Table("user").Where("id=? and phone_search=?", id, 0).Select("phone_number").Scan(&user).Error
	if err != nil {
		return err, ""
	}
	return nil, user.PhoneNumber
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
		Avatar:      param.Avatar,
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
	err = redisCli.HSet("healing2021:hobby", strconv.Itoa(id), value).Err()
	if err != nil {
		return err
	}
	return nil
}

type BasicMsg struct {
	Hobby          []string `json:"hobby"`
	AvatarVisible  int      `json:"avatar_visible"`
	PhoneSearch    int      `json:"phone_search"`
	RealNameSearch int      `json:"real_name_search"`
	Signature      string   `json:"signature"`
	Avatar         string   `json:"avatar"`
	Nickname       string   `json:"nickname"`
	Sex            int      `json:"sex"`
	SelectionNum   int      `json:"selection_num"`
}

func GetBasicMessage(id int) (BasicMsg, error) {
	resp := BasicMsg{}
	db := setting.MysqlConn()
	db.Table("user").Select("selection_num,avatar,nickname,avatar_visible,phone_search,real_name_search,signature").Where("id=?", id).Scan(&resp)
	redisCli := setting.RedisConn()
	value, _ := redisCli.HGet("healing2021:hobby", strconv.Itoa(id)).Bytes()
	if len(value) != 0 {

		err := json.Unmarshal(value, &resp.Hobby)
		return resp, err
	}
	return resp, nil

}
func UpdateUser(user *statements.User, id int, avatar string) (string, error) {
	db := setting.MysqlConn()
	redisCli := setting.RedisConn()
	other := statements.User{}
	db.Table("user").Where("openid=?", user.Openid).Scan(&other)
	if int(other.ID) != id && other.ID != 0 {

		return "", errors.New("error")
	}
	message := make(map[string]interface{})
	message["nickname"] = user.Nickname
	message["avatar_visible"] = user.AvatarVisible
	if user.AvatarVisible == 1 {
		switch other.Sex {
		case 1:
			avatar = tools.GetAvatarUrl(1)
		case 2:
			avatar = tools.GetAvatarUrl(0)
		case 3:
			avatar = tools.GetAvatarUrl(2)

		}

	} else {
		avatar = redisCli.HGet("healing2021:avatar", user.Openid).Val()
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

func UpdateBackground(openid string, background string) error {
	db := setting.MysqlConn()
	err := db.Table("user").Where("openid=?", openid).Update("background", background).Error
	return err
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
	SelectionId string `json:"selection_id"`
	SongName    string `json:"song_name"`
	CreatedAt   string `json:"created_at"`
	Likes       int    `json:"likes"`
	Check       int    `json:"check"`
}
type CoverMsgV2 struct {
	ID          int       `json:"id"`
	SelectionId string    `json:"selection_id"`
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
	Check       int    `json:"check"`
}
type PraiseMsgV2 struct {
	CoverId     int       `json:"cover_id"`
	SelectionId int       `json:"selection_id"`
	SongName    string    `json:"song_name"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int       `json:"id"`
}

func GetUser(id int, module int) interface{} {
	resp := make(map[string]interface{})
	switch module {
	case 1:
		resp["mySelections"] = getSelections(id)
	case 2:
		resp["mySongs"] = getCovers(id, 1)
	case 3:
		resp["myLikes"] = getPraises(id)
	case 4:
		resp["moments"] = getMoments(id)
	}
	return resp
}

func GetCallee(id int, module int) interface{} {
	resp := make(map[string]interface{})
	db := setting.MysqlConn()
	switch module {
	case 1:
		user := UserMsg{}
		db.Table("user").Select("nickname,signature,sex,avatar,school,id").Where("id=?", id).Scan(&user)
		resp["message"] = user
		resp["mySelections"] = getSelections(id)
	case 2:
		resp["mySongs"] = getCovers(id, 2)
	case 3:
		resp["myLikes"] = getPraises(id)
	case 4:
		resp["moments"] = getMoments(id)
	}
	return resp

}

//返回所有信息采用扫描行的方式

func getPraises(user_id int) interface{} {
	cover := []CoverDetails{}
	db := setting.MysqlConn()
	db.Table("cover").Select("sum(praise.is_liked) as likes,user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at ").
		Joins("inner join user on user.id=cover.user_id").
		Joins("inner join praise on cover.id=praise.cover_id").
		Group("cover_id").Order("created_at desc").
		Where("praise.user_id=?", user_id).
		Scan(&cover)
	ch := make(chan int, 15)
	for i, _ := range cover {
		//确认是否点赞
		ViolenceGetLikeheckC(user_id, cover[i], ch)
		cover[i].Check = <-ch
	}

	return cover
}
func getCovers(user_id int, anon int) interface{} {
	cover := []CoverDetails{}
	db := setting.MysqlConn()
	switch anon {
	case 1:
		db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.user_id=" + strconv.Itoa(user_id) + ")" + " as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
			Scan(&cover)
	case 2:
		db.Raw("select likes,avatar,nickname,selection_id,song_name,file,user_id,id,created_at from (select user.avatar,user.nickname,cover.selection_id,cover.song_name,cover.file,cover.user_id,cover.id,cover.created_at from cover inner join user on user.id=cover.user_id where cover.user_id=" + strconv.Itoa(user_id) + " and cover.is_anon=0)" + " as A left join (select cover_id,sum(is_liked) as likes from praise group by cover_id) as B on A.id=B.cover_id order by created_at desc").
			Scan(&cover)
	}
	ch := make(chan int, 15)
	for i, _ := range cover {
		//确认是否点赞
		go ViolenceGetLikeheckC(user_id, cover[i], ch)
		cover[i].Check = <-ch
	}

	return cover
}
func getSelections(user_id int) interface{} {
	db := setting.MysqlConn()
	selection := []SelectionMsg{}
	db.Table("selection").Select("song_name,id,created_at,remark").
		Where("user_id=?", user_id).
		Scan(&selection)
	return selection
}

func getMoments(user_id int) interface{} {
	moment2 := []MomentMsgV2{}
	moment := []MomentMsg{}
	db := setting.MysqlConn()
	db.Raw("select  sum(praise.is_liked) as likes,moment.content,moment.state,moment.song_name,moment.id,moment.created_at from moment  left join praise on moment.id=praise.moment_id where user_id=" + strconv.Itoa(user_id) + " group by moment_id;").Scan(&moment2)
	moment = make([]MomentMsg, len(moment2))
	ch := make(chan int, 15)
	for i, _ := range moment {
		//确认是否点赞
		ViolenceGetLikeheckM(user_id, moment2[i], ch)
		moment[i].Check = <-ch
		moment[i].State = tools.DecodeStrArr(moment2[i].State)
		moment[i].ID = moment2[i].ID
		moment[i].CreatedAt = tools.DecodeTime(moment2[i].CreatedAt)
		moment[i].Likes = moment2[i].Likes
		moment[i].SongName = moment2[i].SongName
		moment[i].Content = moment2[i].Content
	}
	return moment
}

type MomentMsgV2 struct {
	SongName  string    `json:"song_name"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	State     string    `json:"state"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	Check     int       `json:"check"`
}
type MomentMsg struct {
	SongName  string   `json:"song_name"`
	CreatedAt string   `json:"created_at"`
	ID        int      `json:"id"`
	State     []string `json:"state"`
	Content   string   `json:"content"`
	Likes     int      `json:"likes"`
	Check     int      `json:"check"`
}

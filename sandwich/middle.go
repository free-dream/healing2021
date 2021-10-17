package sandwich

import (
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

var (
	Sandwich *Middle
)

//redis mail,类型分为cover,moment,moment-comment
type mail struct {
	key      string //mailid
	mailtype string //点赞类型
}

//中间层主体
type Middle struct {
	mailbox chan *mail
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	Sandwich = Default()
}

func DummyRedisFile() {
	Conn := setting.RedisConn()
	Conn.RPush("test", "this is a dummy data")
}

//default下默认mailbox长度为200
func Default() *Middle {
	middle := new(Middle)
	middle.mailbox = make(chan *mail, 200)
	return middle
}

//自定义mailbox长度
func NewMiddle(mailboxlen int) *Middle {
	middle := new(Middle)
	middle.mailbox = make(chan *mail, mailboxlen)
	return middle
}

func GenerateMail() *mail {
	uuid := uuid.NewV4()
	real := uuid.String()
	return &mail{
		key: real,
	}
}

func GenerateMailTest() *mail {
	return &mail{
		key: "test",
	}
}

//将redis拆包成对应的mysql表数据并更新
//redis:
//key:uuid,value:[user_id,target_id] --->redis中的mail数据
//key:user_id,value:target_id --->点赞查重，24h
func (mail *mail) ToSql() {
	redisDb := setting.RedisClient

	//提取数据
	user_id, _ := strconv.Atoi(redisDb.LPop(mail.key).Val())
	target_id, _ := strconv.Atoi(redisDb.LPop(mail.key).Val())

	//判断写入
	dao.UpdateLikesByID(user_id, target_id, mail.mailtype)

}

func (mail *mail) ToSqlTest() {
	redisDb := setting.RedisClient
	data := redisDb.LPop(mail.key).Val()
	data2 := redisDb.LPop(mail.key).Val()
	fmt.Println(1)
	fmt.Println(data, data2)
}

//sql缓存到redis
func (middle *Middle) cache() {}

//优先对table更新
func (middle *Middle) Update() {
	for {
		select {
		case mail := <-middle.mailbox:
			mail.ToSqlTest()
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

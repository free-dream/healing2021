package sandwich

import (
	"fmt"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var (
	Sandwich *Middle
)

//redis mail,类型分为cover,moment,moment-comment
type mail struct {
	key      string //mailid
	mailtype string //点赞类型
}

//table指向对应表
type table struct {
	tabletype *gorm.Model
	request   string
}

//中间层主体
type Middle struct {
	mailbox  chan *mail
	tablebox chan *table
}

//错误处理函数
func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

//全局初始化
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

//点赞的话生成一个mail传入channel
func GenerateMail(mailtype string) *mail {
	uuid := uuid.NewV4()
	real := uuid.String()
	return &mail{
		key:      real,
		mailtype: mailtype,
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
//key:user_id,value:target_id --->点赞查重，expile:24h
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

//优先对table更新
func (middle *Middle) Update() {
	for {
		select {
		// case table := <-middle.tablebox:
		// 	 middle.Cache(table)
		case mail := <-middle.mailbox:
			mail.ToSqlTest()
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

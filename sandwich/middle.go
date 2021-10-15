package sandwich

import (
	"fmt"
	"time"
)

//redis mail
type mail struct {
	key      string
	value    interface{}
	mailtype string
}

//sqlHandler主体
type Middle struct {
	mailbox chan *mail
}

//default下默认mailbox长度为10
func Default() *Middle {
	middle := new(Middle)
	middle.mailbox = make(chan *mail, 10)
	return middle
}

func NewMiddle(mailboxlen int, tableboxlen int) *Middle {
	middle := new(Middle)
	middle.mailbox = make(chan *mail, mailboxlen)
	return middle
}

//将redis拆包成对应的mysql表数据
func (mail *mail) toSql() {}

//sql缓存到redis
func (middle *Middle) cache() {}

//优先对table更新
func (middle *Middle) Update() {
	for {
		select {
		case table := <-middle.mailbox:
			//
			fmt.Println(table)
			//
			continue
		case mail := <-middle.mailbox:
			if value, ok := mail.value.(int); ok {
				mail.toSql()
				//
				fmt.Println(value)
				//
			} else if value, ok := mail.value.(string); ok {
				mail.toSql()
				//
				fmt.Println(value)
				//
			}
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

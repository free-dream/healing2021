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
	next     *mail
}

type table struct{}

//sqlHandler主体
type Middle struct {
	mailbox  chan *mail
	tablebox chan *table
}

//将redis拆包成对应的mysql表数据
func (mail *mail) toSql() {}

//sql缓存到redis
func (middle *Middle) cache() {}

//redis持久化到sql,休眠1s后
func (middle *Middle) Update() {
	for {
		select {
		case mail := <-middle.mailbox:
			if value, ok := mail.value.(int); ok {
				mail.toSql()
				fmt.Println(value)
			} else if value, ok := mail.value.(string); ok {
				mail.toSql()
				fmt.Println(value)
			}
		case table := <-middle.mailbox:
			fmt.Println(table)
		default:
			time.Sleep(1)
		}
	}
}

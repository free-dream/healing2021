package tools

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type RedisUser struct {
	gorm.Model
	OpenId   string
	NickName string
	TrueName string
	More     string
	Campus   string
	Avatar   string
	Phone    string
	Sex      int
	Hobby    string
	Money    int
	Setting1 int
	Setting2 int
	Setting3 int
	Postbox  string
}

func GetUser(c *gin.Context) RedisUser {
	session := sessions.Default(c)
	data := session.Get("user")

	// 用于测试
	if IsDebug() {
		fmt.Println("默认提供一个已经登录好的用户,ID=1")
		return RedisUser{
			Model: gorm.Model{ID: 1},
		}
	}

	if data == nil {
		c.JSON(401, e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_AUTH)})
		c.Abort()
		return RedisUser{}
	}
	return data.(RedisUser)
}

//获取openid
func GetOpenid(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	raw := session.Get("openid")
	openid := raw.(string)
	return openid
}

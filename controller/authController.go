package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//向前端返回用户微信昵称

func FakeLogin(ctx *gin.Context) {
	user := statements.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
		// ctx.JSON(401, gin.H{
		// 	"message": "error param",
		// })
		// return
	}
	openid, err1 := dao.FakeCreateUser(&user)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称已存在，无法注册",
		})
		return
	}
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	option := sessions.Options{
		MaxAge: 3600,
	}

	session.Options(option)
	var redisUser tools.RedisUser
	session.Set("user", redisUser)
	session.Save()
	session.Set("openid", openid)
	err2 := session.Save()
	if err2 != nil {
		panic(err2)
		// return
	}
	ctx.JSON(200, "OK")
}

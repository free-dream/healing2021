package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/models/statements"
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
	id, err1 := dao.FakeCreateUser(&user)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"message": "用户不存在",
		})
		return
	}

	//建立任务表
	err = dao.GenerateTasktable(normaltasks, id)
	if err != nil {
		panic(err)
	}
	//

	session := sessions.Default(ctx)
	session.Set("avatar", user.Avatar)
	session.Set("openid", user.Openid)
	session.Set("user_id", id)
	session.Save()
	ctx.JSON(200, "OK")
}

func FakeLoginEasy(ctx *gin.Context) {
	session := sessions.Default(ctx)
	redirect, _ := ctx.GetQuery("redirect")
	session.Set("avatar", "我的头像 url")
	session.Set("openid", 123456)
	session.Set("user_id", 1)
	session.Save()
	ctx.Redirect(302, redirect)
	ctx.Abort()
}

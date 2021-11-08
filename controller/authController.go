package controller

import (
	"fmt"
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
	// 测试
	fmt.Println(id)
	//
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称重复",
		})
		return
	}

	//
	//fmt.Println("FakeLogin", id)
	//

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
	ctx.JSON(200, gin.H{
		"message": "登录成功",
	})
}

func FakeLoginEasy(ctx *gin.Context) {
	session := sessions.Default(ctx)
	//redirect, _ := ctx.GetQuery("redirect")
	session.Set("avatar", "我的头像 url")
	session.Set("openid", "123456")
	session.Set("user_id", 1)
	session.Save()
	ctx.JSON(200, "OK")
}

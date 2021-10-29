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
	session := sessions.Default(ctx)
	session.Set("openid", user.Openid)
	session.Set("user_id", id)
	session.Save()
	ctx.JSON(200, "OK")
}

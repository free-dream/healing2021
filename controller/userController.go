package controller

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)

	user := models.User{}
	err := ctx.ShouldBindJSON(&user)
	user.Openid = openid
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	err1, err2 := models.CreateUser(&user)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称已存在，无法注册",
		})
		return
	}
	if err2 != nil {
		panic(err)
		return
	}
	ctx.JSON(200, "OK")

}
func Updater(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	user := models.User{}
	err := ctx.ShouldBind(&user)
	fmt.Println(openid)

	if err != nil {
		panic(err)
		ctx.JSON(403, gin.H{
			"message": "修改失败",
		})
		return
	}
	err1 := models.UpdateUser(&user, openid)
	if err1 != nil {
		panic(err1)
		ctx.JSON(200, "OK")
	}

}
func Fetcher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	user := models.GetUser(openid)
	ctx.JSON(200, user)

}
func Refresher(ctx *gin.Context) {

}
func GetOther(ctx *gin.Context) {

}

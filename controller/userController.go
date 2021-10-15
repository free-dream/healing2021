package controller

import (
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
	err, err2 := models.CreateUser(&user)
	if err != nil {
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
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return
	}
	err = models.UpdateUser(&user, openid)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return

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

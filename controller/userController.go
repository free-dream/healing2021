package controller

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//用户注册
func Register(ctx *gin.Context) {
	//登录奖励机制尚未完成
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	headImgUrl := session.Get("headImgUrl").(string)

	user := dao.User{}
	err := ctx.ShouldBindJSON(&user)
	user.Openid = openid
	user.Avatar = headImgUrl
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	err, err2 := dao.CreateUser(&user)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称已存在，无法注册",
		})
		return
	}
	if err2 != nil {
		panic(err)
	}
	ctx.JSON(200, "OK")

}
func Updater(ctx *gin.Context) {
	//用户更新
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	user := dao.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return
	}
	err = dao.UpdateUser(&user, openid)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return

	}

}

//获取用户信息
func Fetcher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	user := dao.GetUser(openid)
	ctx.JSON(200, user)

}

//更新背景
type obj struct {
	Background string `json:"background" binding:"required"`
}

func Refresher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	obj := obj{}
	err := ctx.ShouldBindJSON(&obj)

	if err != nil {
		panic(err)
	}
	err = dao.UpdateBackground(openid, obj.Background)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "更新失败",
		})
		return

	}
	ctx.JSON(200, "OK")

}
func GetOther(ctx *gin.Context) {
	param, bool := ctx.GetQuery("calleeId")

	if !bool {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	calleeId, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		panic(err)
	}
	resp := dao.GetCallee(calleeId)
	ctx.JSON(200, resp)

}

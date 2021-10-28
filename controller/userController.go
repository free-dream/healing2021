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
	//headImgUrl := session.Get("headImgUrl").(string)

	user := dao.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	user.Openid = openid
	//user.Avatar = headImgUrl
	head, err := strconv.Atoi(user.PhoneNumber[0:3])
	length := len(user.PhoneNumber)
	if err != nil || head < 130 || head >= 200 || length != 11 {

		ctx.JSON(403, gin.H{
			"message": "手机号格式错误",
		})
		return
	}
	id, err := dao.CreateUser(&user)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称已存在，无法注册",
		})
		return
	}
	session.Set("user_id", id)
	err = session.Save()
	if err != nil {
		panic(err)

	}
	ctx.JSON(200, gin.H{
		"user_id": id,
	})

}
func Updater(ctx *gin.Context) {
	//用户更新
	session := sessions.Default(ctx)
	id := session.Get("user_id").(int)
	avatar := session.Get("avatar").(string)
	user := dao.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return
	}

	avatar, err = dao.UpdateUser(&user, id, avatar)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return

	}
	ctx.JSON(200, gin.H{
		"avatar": avatar,
	})

}

//获取用户信息
func Fetcher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id := session.Get("id").(int)
	user := dao.GetUser(id)
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
	param, bl := ctx.GetQuery("calleeId")
	if !bl {
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

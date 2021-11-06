package controller

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//用户登录
func Login(openid string) (int, int) {
	user := statements.User{
		Openid: openid,
	}
	isExisted, id := dao.CreateUser(&user)
	//根据给定的数组生成任务表
	var err1 error
	check, err1 := dao.CheckTasks(id)
	if !check && gorm.IsRecordNotFoundError(err1) {
		err := dao.GenerateTasktable(normaltasks, id)
		if err != nil {
			panic(err)
		}
	}
	return isExisted, id
}

//判断用户是否是管理员
func GodJudger(nickname string) bool {
	isAdministrator := dao.Authentication(nickname)
	return isAdministrator
}

//用户注册,新增生成任务表机制
func Register(ctx *gin.Context) {
	//登录奖励机制,需要时可实现于task内
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	//headImgUrl := session.Get("headImgUrl").(string)
	id := session.Get("user_id").(int)

	user := statements.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
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
	err = dao.RefineUser(&user, id)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "昵称已存在，无法注册",
		})
		return
	}

}

func HobbyPoster(ctx *gin.Context) {
	var hobby []string
	id := sessions.Default(ctx).Get("user_id").(int)
	err := ctx.ShouldBindJSON(&hobby)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
	}
	err = dao.HobbyStore(hobby, id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "ok")
}
func Updater(ctx *gin.Context) {
	//用户更新
	session := sessions.Default(ctx)
	id := session.Get("user_id").(int)
	avatar := session.Get("avatar").(string)
	user := statements.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
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
	id := session.Get("user_id").(int)

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
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	calleeId, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		panic(err)
	}
	resp := dao.GetCallee(calleeId)
	ctx.JSON(200, resp)

}

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
func Login(openid string, nickname string, avatar string) int {
	user := statements.User{
		Openid:   openid,
		Nickname: nickname,
		Avatar:   avatar,
	}
	id := dao.CreateUser(user)
	//根据给定的数组生成任务表
	var err1 error
	check, err1 := dao.CheckTasks(id)
	if !check && gorm.IsRecordNotFoundError(err1) {
		err := dao.GenerateTasktable(normaltasks, id)
		if err != nil {
			panic(err)
		}
	}
	return id
}
func PhoneCaller(ctx *gin.Context) {
	id, _ := ctx.GetQuery("user_id")
	user_id, _ := strconv.Atoi(id)
	err, phoneNumber := dao.GetPhoneNumber(user_id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "无电话号码",
		})
	}
	ctx.JSON(200, gin.H{
		"phone_number": phoneNumber,
	})
}

//判断用户是否是管理员,是否注册
func Judger(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(int)
	is_existed := dao.Exist(session.Get("openid").(string))
	resp, err := dao.GetBasicMessage(user_id)
	session.Set("headImgUrl", resp.Avatar)
	session.Set("nickname", resp.Nickname)
	session.Save()
	if err != nil {
		panic(err)
	}
	is_administrator := dao.Authentication(resp.Nickname)
	ctx.JSON(200, gin.H{
		"user_id":          user_id,
		"is_existed":       is_existed,
		"avatar":           resp.Avatar,
		"nickname":         resp.Nickname,
		"is_administrator": is_administrator,
		"hobby":            resp.Hobby,
		"avatar_visible":   resp.AvatarVisible,
		"phone_search":     resp.PhoneSearch,
		"real_name_search": resp.RealNameSearch,
		"signature":        resp.Signature,
	})
}

//用户注册,新增生成任务表机制
func Register(ctx *gin.Context) {
	//登录奖励机制,需要时可实现于task内
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	headImgUrl := session.Get("headImgUrl").(string)
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
	user.Avatar = headImgUrl
	if user.PhoneNumber != "" {
		body, err1 := strconv.Atoi(user.PhoneNumber)

		if err1 != nil || body <= 13000000000 || body >= 20000000000 {

			ctx.JSON(403, gin.H{
				"message": "手机号格式错误",
			})
			return
		}
	}

	err = dao.RefineUser(user, id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "OK")

}

func HobbyPoster(ctx *gin.Context) {
	hobby := dao.Hobby{}
	id := sessions.Default(ctx).Get("user_id").(int)
	err := ctx.ShouldBindJSON(&hobby)
	if err != nil {
		ctx.JSON(400, gin.H{

			"message": "error param",
		})
		return
	}
	err = dao.HobbyStore(hobby.Hobby, id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "ok")
}
func Updater(ctx *gin.Context) {
	//用户更新
	session := sessions.Default(ctx)
	id := session.Get("user_id").(int)
	openid := session.Get("openid").(string)
	avatar := session.Get("headImgUrl").(string)
	user := statements.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return
	}
	user.Openid = openid
	avatar, err = dao.UpdateUser(&user, id, avatar)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "修改失败",
			"error":   err,
		})
		return

	}
	session.Set("headImgUrl", avatar)
	session.Set("nickname", user.Nickname)
	session.Save()
	ctx.JSON(200, gin.H{
		"avatar": avatar,
	})

}

//获取用户信息
func Fetcher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	param := ctx.Query("module")
	module, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
	}
	id := session.Get("user_id").(int)
	user := dao.GetUser(id, module)
	ctx.JSON(200, user)

}

//更新背景
type Obj struct {
	Background string `json:"background" binding:"required"`
}

func Refresher(ctx *gin.Context) {
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	obj := Obj{}
	err := ctx.ShouldBindJSON(&obj)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
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
	param = ctx.Query("module")
	module, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
	}
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		panic(err)
	}
	resp := dao.GetCallee(calleeId, module)
	ctx.JSON(200, resp)

}

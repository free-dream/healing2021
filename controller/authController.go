package controller

import (
	"io"
	"os"

	"git.100steps.top/100steps/ginwechat"
	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//登录成功进行用户的数据缓存

// 微信登录
func Login() {
	r := gin.Default()
	file, _ := os.Create("log")
	gin.DefaultWriter = io.MultiWriter(file)
	ginwechat.UpdateEngine(r, &ginwechat.Config{
		Appid:     "wx293bc6f4ee88d87d",
		Appsecret: "",
		BaseUrl:   "https://healing2021.100steps.top",
		StoreSession: func(ctx *gin.Context, wechatUser *ginwechat.WechatUser) error {
			session := sessions.Default(ctx)
			//等任务出来具体修改
			/*session.Clear()
			session.Save()
			option := sessions.Options{
				MaxAge: 3600,
			}

			session.Options(option)
			var redisUser tools.RedisUser
			var user statements.User
			redisUser=wechatUser
			session.Set("user", redisUser)
			session.Save()
			redis_cli := setting.RedisClient

			// 加积分并且记录该用户本日登陆
			t := time.Now()
			t_zero := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
			t_to_tomorrow := 24*60*60 - (t.Unix() - t_zero)
			logined := !redis_cli.SetNX(fmt.Sprintf("healing2021:logined_user:%d", user.ID), 0, time.Duration(t_to_tomorrow)*time.Second).Val()
			if !logined {
				models.FinishTask("1", user.ID)
			}*/
			return session.Save()
		},
	})
	//增加一个登录时缓存用户的步骤
}

//向前端返回用户微信昵称

func WechatUser(ctx *gin.Context) {

	session := sessions.Default(ctx)
	nickname := session.Get("nickname").(string)
	ctx.JSON(200, gin.H{
		"nickname": nickname,
	})
}
func FakeLogin(ctx *gin.Context) {
	user := dao.User{}
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
	session.Set("openid", openid)
	err2 := session.Save()
	if err2 != nil {
		panic(err2)
		// return
	}
	ctx.JSON(200, "OK")
}

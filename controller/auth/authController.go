package auth

import (
	"git.100steps.top/100steps/ginwechat"
	"git.100steps.top/100steps/healing2021_be/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 微信登录
func Login() {
	r := gin.Default()
	file, _ := os.Create("log")
	gin.DefaultWriter = io.MultiWriter(file)
	//
	ginwechat.UpdateEngine(r, &ginwechat.Config{
		Appid:     "wx293bc6f4ee88d87d",
		Appsecret: "",
		BaseUrl:   "http://healing2021.100steps.top",
		StoreSession: func(ctx *gin.Context, user *ginwechat.WechatUser) error {
			session := sessions.Default(ctx)
			session.Set("openid", user.OpenID)
			session.Set("nickname", user.Nickname)
			session.Set("headImgUrl", user.HeadImgUrl)
			return session.Save()
		},
	})

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
	user := models.User{}
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		panic(err)
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	openid, err1 := models.FakeCreateUser(&user)
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
		return
	}
	ctx.JSON(200, "OK")
}

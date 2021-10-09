package router

import (
	"git.100steps.top/100steps/ginwechat"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func SetupRouter() *gin.Engine {
	// 路由组
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
			return session.Save()
		},
	})
	return r
}

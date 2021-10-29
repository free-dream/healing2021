package router

import (
	"encoding/gob"
	"git.100steps.top/100steps/ginwechat"
	"git.100steps.top/100steps/healing2021_be/models"
	"io"
	"log"
	"os"
	"time"

	"git.100steps.top/100steps/healing2021_be/controller"
	"git.100steps.top/100steps/healing2021_be/controller/middleware"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store

func SetupRouter() *gin.Engine {

	var test_prefix string

	if tools.IsDebug() {
		test_prefix = "/test"
		models.FakeData()
	} else {
		test_prefix = ""
	}
	r := gin.Default()

	f, _ := os.Create(tools.GetConfig("log", "location"))
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Timeout(time.Minute))
	r.Use(middleware.Cors())
	// 注册sessions组件，使用redis作为驱动
	gob.Register(tools.RedisUser{})
	var err error
	store, err = redis.NewStore(30, "tcp", tools.GetConfig("redis", "addr"), "", []byte("__100steps__100steps__100steps__"))
	if err != nil {
		log.Panicln(err.Error())
	}
	r.Use(sessions.Sessions("healing2021_session", store))
	ginwechat.UpdateEngine(r, &ginwechat.Config{
		Appid:     "wx293bc6f4ee88d87d",
		Appsecret: "",
		BaseUrl:   "https://healing2021.100steps.top",
		StoreSession: func(ctx *gin.Context, wechatUser *ginwechat.WechatUser) error {
			session := sessions.Default(ctx)
			session.Set("openid", wechatUser.OpenID)
			session.Set("headImgUrl", wechatUser.HeadImgUrl)
			session.Set("nickname", wechatUser.Nickname)
			ctx.JSON(200, gin.H{
				"nickname": wechatUser.Nickname,
			})
			//登录任务
			return session.Save()
		},
	})

	// ping 测试
	r.GET(test_prefix+"/ping", func(ctx *gin.Context) {
		ctx.JSON(200, e.ErrMsgResponse{Message: "pong"})
		return
	})

	//中间件验证
	if tools.IsDebug() {
		r.POST(test_prefix+"/user", controller.FakeLogin)
	} else {
		r.Use(middleware.IdentityCheck())
	}
	// 业务路由
	api := r.Group(test_prefix + "/api")

	//user 模块

	api.POST("/user", controller.Register)
	api.PUT("/user", controller.Updater)
	api.GET("/user", controller.Fetcher)
	api.POST("/background", controller.Refresher)
	api.GET("/callee", controller.GetOther)
	//qiniu
	api.GET("/qiniu/token", controller.QiniuToken)
	//经典治愈 模块
	api.GET("/healingPage", controller.HealingPageFetcher)
	api.GET("/healing/bulletin", controller.AdsPlayer)
	api.GET("/healing/selections/list", controller.SelectionFetcher)
	api.GET("/healing/covers/list", controller.CoverFetcher)
	api.POST("/healing/cover", controller.Recorder)
	api.POST("/healing/selection", controller.Selector)
	//经典治愈——抽奖箱
	api.GET("healing/lotterybox/prizes", controller.GetPrizes)
	api.GET("/healing/lotterybox/draw", controller.Draw)
	api.GET("/healing/lotterybox/lotteries", controller.GetLotteries)
	api.GET("/healing/lotterybox/tasktable", controller.GetTasktable)
	// childhood 模块
	api.GET("/childhood/rank", controller.GetRank)
	api.GET("/childhood/list", controller.GetList)
	api.GET("/childhood/original/info", controller.GetOriginalInfo)
	api.GET("/childhood/original/cover", controller.GetOriginalSingerList)
	api.GET("/healing/covers/player", controller.GetPlayer)
	api.POST("/healing/covers/jump", controller.JumpSongs)
	// 广场 模块
	api.GET("/dynamics/list/:method", controller.GetMomentList)
	api.POST("/dynamics/send", controller.PostMoment)
	api.GET("/dynamics/detail/:id", controller.GetMomentDetail)
	api.POST("/dynamics/comment", controller.PostComment)
	api.GET("/dynamics/comment/:id", controller.GetCommentList)
	api.GET("/dynamics/hot", controller.DynamicsSearchHot)
	api.GET("/dynamics/states", controller.OursStates)
	api.GET("/dynamics/song", controller.HotSong)
	// 管理员操作 模块
	api.POST("/administrators", controller.DeleteContent)
	// 通用操作 模块
	api.PUT("/like", controller.Like)
	return r
}

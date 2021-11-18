package router

import (
	"io"
	"log"
	"os"
	"time"

	"git.100steps.top/100steps/ginwechat"

	"git.100steps.top/100steps/healing2021_be/controller"
	"git.100steps.top/100steps/healing2021_be/controller/middleware"
	"git.100steps.top/100steps/healing2021_be/controller/ws"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store

func SetupRouter() *gin.Engine {
	r := gin.Default()

	f, _ := os.Create(tools.GetConfig("log", "location"))
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Timeout(time.Minute))
	r.Use(middleware.Cors())

	// 注册sessions组件，使用redis作为驱动
	//gob.Register(tools.RedisUser{})
	var err error

	store, err = redis.NewStoreWithDB(30, "tcp", tools.GetConfig("redis", "addr"), "", tools.GetConfig("redis", "db"), []byte("__100steps__100steps__100steps__"))

	if err != nil {
		log.Panicln(err.Error())
	}
	var url string
	r.Use(sessions.Sessions("healing2021_session", store))
	if tools.IsDebug() {
		url = tools.GetConfig("domain", "testUrl")
	} else {
		url = tools.GetConfig("domain", "url")
	}

	ginwechat.UpdateEngine(r, &ginwechat.Config{
		Appid:     "wx293bc6f4ee88d87d",
		Appsecret: "",
		BaseUrl:   url,
		StoreSession: func(ctx *gin.Context, wechatUser *ginwechat.WechatUser) error {
			redirect, _ := ctx.GetQuery("redirect")
			user_id := controller.Login(wechatUser.OpenID, wechatUser.Nickname, wechatUser.HeadImgUrl)
			session := sessions.Default(ctx)
			session.Set("user_id", user_id)
			session.Set("openid", wechatUser.OpenID)
			err = session.Save()
			ctx.Redirect(302, redirect)
			return err
		},
	})

	// ping 测试
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, e.ErrMsgResponse{Message: "pong"})
		return
	})

	//假登录
	if tools.IsDebug() {
		r.POST("/user", controller.FakeLogin)
		r.GET("/userEasy", controller.FakeLoginEasy)
	}

	// 业务路由

	api := r.Group("/api")
	//中间件登陆拦截
	api.Use(middleware.IdentityCheck())
	//任务初始化
	api.Use(middleware.TaskInit())
	// ws
	api.GET("/ws", ws.WsHandler)
	api.GET("/ws/history", ws.WsData)

	//user 模块
	api.GET("/phoneNumber", controller.PhoneCaller)
	api.GET("/user", controller.Judger)
	api.POST("/user", controller.Register)
	api.POST("/hobby", controller.HobbyPoster)
	api.PUT("/user", controller.Updater)
	api.GET("/userMsg", controller.Fetcher)
	api.POST("/background", controller.Refresher)
	api.GET("/callee", controller.GetOther)

	//qiniu
	api.GET("/qiniu/token", controller.QiniuToken)

	//经典治愈 模块
	api.GET("/healingPage", controller.HealingPageFetcher)
	api.GET("/healing/bulletin", controller.AdsPlayer)
	api.GET("/healing/selections/list", controller.SelectionFetcher)
	api.GET("/healing/covers/list", controller.CoverFetcher)
	api.POST("/healing/cover", controller.Recorder)     //植入任务 2021.11.1
	api.POST("/healing/selection", controller.Selector) //植入任务 2021.11.1
	api.GET("/healing/devotion", controller.DevotionPlayer)

	//经典治愈——抽奖箱
	// api.GET("healing/lotterybox/prizes", controller.GetPrizes)
	// api.GET("/healing/lotterybox/drawcheck", controller.DrawCheck)
	api.POST("/healing/lotterybox/draw", controller.Draw)
	api.GET("/healing/lotterybox/lotteries", controller.GetLotteries)
	api.GET("/healing/lotterybox/tasktable", controller.GetTasktable)
	api.GET("/healing/lotterybox/points", controller.GetUserPoints)

	//经典治愈——排行榜
	api.GET("/healing/rank/:school", controller.GetRanking)
	api.GET("/healing/rank/user", controller.GetMyRank)

	//经典治愈——每日热榜
	api.GET("/healing/dailyrank/:date", controller.GetDailyrank)
	api.GET("/healing/dailyrank/all", controller.GetAllrank)

	//经典治愈——搜索
	api.POST("/healing/search", controller.Search)

	// childhood 模块
	api.GET("/childhood/rank", controller.GetRank)
	api.GET("/childhood/list", controller.GetList)
	api.GET("/childhood/original/info", controller.GetOriginalInfo)
	api.GET("/childhood/original/covers", controller.GetOriginalSingerList)
	api.GET("/healing/covers/player", controller.GetPlayer)
	api.POST("/healing/covers/jump", controller.JumpSongs)

	// 广场 模块
	api.GET("/dynamics/list/new", controller.GetMomentNew)
	api.GET("/dynamics/list/recommend", controller.GetMomentRecommend)
	api.GET("/dynamics/list/search", controller.GetMomentSearch)
	api.POST("/dynamics/send", controller.PostMoment) //植入任务 2021.11.1
	api.GET("/dynamics/detail/:id", controller.GetMomentDetail)
	api.POST("/dynamics/comment", controller.PostComment)
	api.GET("/dynamics/comment/:id", controller.GetCommentList)
	api.GET("/dynamics/hotsearch", controller.DynamicsSearchHot)
	api.GET("/dynamics/ourstates", controller.OursStates)
	api.GET("/dynamics/hotsong", controller.HotSong)

	//通用操作 模块
	api.PUT("/like", controller.Like)

	// 管理员操作 模块
	api.Use(middleware.Authentication())
	api.POST("/administrators/comment", controller.DeleteContent)
	api.POST("/administrators/moment", controller.DeleteMoment)

	return r
}

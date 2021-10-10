package router

import (
	"encoding/gob"
	"git.100steps.top/100steps/healing2021_be/controller/childhood"
	"git.100steps.top/100steps/healing2021_be/controller/middleware"
	"git.100steps.top/100steps/healing2021_be/controller/playground"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

var store redis.Store

func SetupRouter() *gin.Engine {

	var test_prefix string

	if tools.IsDebug() {
		test_prefix = "/test"
	} else {
		test_prefix = ""
	}
	r := gin.Default()

	f, _ := os.Create(tools.GetConfig("log", "location"))
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Timeout(time.Minute))

	// 注册sessions组件，使用redis作为驱动
	gob.Register(tools.RedisUser{})
	var err error
	store, err = redis.NewStore(30, "tcp", tools.GetConfig("redis", "addr"), "", []byte("__100steps__100steps__100steps__"))
	if err != nil {
		log.Panicln(err.Error())
	}
	r.Use(sessions.Sessions("healing2021_session", store))

	if tools.IsDebug() {
		r.Use(middleware.Cors())
	}

	// ping 测试
	r.GET(test_prefix+"/ping", func(ctx *gin.Context) {
		ctx.JSON(200, e.ErrMsgResponse{Message: "pong"})
		return
	})

	// 业务路由
	api := r.Group("/api")

	// childhood 模块
	api.GET("/childhood/rank", childhood.GetRank)
	api.GET("/childhood/list", childhood.GetList)
	api.GET("/childhood/original/:name/info", childhood.GetOriginalInfo)
	api.GET("/childhood/original/:name/cover", childhood.GetOriginalSingerList)
	api.POST("/healing/player", childhood.LoadSongs)

	// 广场 模块
	api.GET("/dynamics/list/:method", playground.GetMomentList)
	api.GET("/dynamics/send", playground.PostMoment)
	api.GET("/dynamics/detail/:id", playground.GetMomentDetail)
	api.POST("/dynamics/comment", playground.PostComment)
	api.GET("/dynamics/comment/:id", playground.GetCommentList)
	api.PUT("/laud/:type/:id", playground.PriseOrNot)

	return r
}

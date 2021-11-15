package controller

//dailyrank除了当日热榜需要更新之外其它可以直接缓存
import (
	"encoding/json"
	"log"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GET /healing/dailyrank/all
func GetAllrank(ctx *gin.Context) {
	// 读取userid
	UserId := sessions.Default(ctx).Get("user_id").(int)
	var (
		respCovers []respModel.HotResp
		err        error
	)
	// 提取缓存
	tempdata := sandwich.GetDailyRankByDate("all")

	if tempdata != "" {
		data := []byte(tempdata)
		err = json.Unmarshal(data, &respCovers)
		if err != nil {
			panic(err) //这种错误出现的话直接挂掉，绝对是代码之外的问题
		}
	} else {
		respCovers, err = dao.GetCoversByLikes()
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
			return
		}
		//更新点赞确认即可,视前端需求
		for i, item := range respCovers {
			bo, err := dao.PackageCheck(UserId, "cover", item.Cover_Id)
			if err != nil {
				log.Printf(err.Error())
				respCovers[i].Check = 0
			} else if bo {
				respCovers[i].Check = 1
			} else {
				respCovers[i].Check = 0
			}
		}
		//缓存
		jsondata, _ := json.Marshal(respCovers)
		cache := string(jsondata)
		err = sandwich.CacheDailyRank("all", cache)
		if err != nil {
			log.Fatal("redis缓存出错")
		}
	}

	ctx.JSON(200, respCovers)
	return
}

// GET /healing/dailyrank/:date
func GetDailyrank(ctx *gin.Context) {
	date := ctx.Param("date")
	// 读取userid
	UserId := sessions.Default(ctx).Get("user_id").(int)
	var (
		respCovers []respModel.HotResp
		err        error
	)
	// 提取缓存
	tempdata := sandwich.GetDailyRankByDate(date)
	if tempdata != "" {
		data := []byte(tempdata)
		err = json.Unmarshal(data, &respCovers)
		if err != nil {
			panic(err) //这种错误出现的话直接挂掉，绝对是代码之外的问题
		}
	} else {
		respCovers, err = dao.GetCoversByDate(date)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
			return
		}
		//更新点赞确认即可,视前端需求
		for i, item := range respCovers {
			bo, err := dao.PackageCheck(UserId, "cover", item.Cover_Id)
			if err != nil {
				log.Printf(err.Error())
				respCovers[i].Check = 0
			} else if bo {
				respCovers[i].Check = 1
			} else {
				respCovers[i].Check = 0
			}
		}
		//缓存
		jsondata, _ := json.Marshal(respCovers)
		cache := string(jsondata)
		err = sandwich.CacheDailyRank("all", cache)
		if err != nil {
			log.Fatal("redis缓存出错")
		}
	}

	ctx.JSON(200, respCovers)
}

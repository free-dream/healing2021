package controller

//dailyrank除了当日热榜需要更新之外其它可以直接缓存
import (
	"encoding/json"
	"log"
	"regexp"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GET /healing/dailyrank/all
func GetAllrank(ctx *gin.Context) {
	//读取userid
	UserId := sessions.Default(ctx).Get("user_id").(int)
	// //提取缓存
	// tempdata := sandwich.GetDailyRankByDate("all")
	// var respCovers []resp.HotResp
	// if tempdata != "" {
	// 	data := []byte(tempdata)
	// 	err := json.Unmarshal(data, &respCovers)
	// 	if err == nil {
	// 		ctx.JSON(200, respCovers)
	// 		return
	// 	} else {
	// 		panic(err)
	// 	}
	// }

	//读取数据库
	raws, likes, err := dao.GetCoversByLikes()
	respCovers := make([]resp.HotResp, 0)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		return
	}
	for i, cover := range raws {
		nickname, err := dao.GetUserNickname(cover.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
			return
		}
		//点赞确认
		coverid := likes[i].CoverId
		boolean, err1 := dao.PackageCheckMysql(UserId, "cover", coverid)
		if err1 != nil {
			log.Printf(err1.Error()) //一般不会出问题
		}
		var check int
		if boolean {
			check = 1
		}
		check = 0
		respCover := resp.HotResp{
			CoverId:  coverid,
			Avatar:   cover.Avatar,
			Nickname: nickname,
			Posttime: cover.CreatedAt.String(),
			Likes:    likes[i].Likes,
			Songname: cover.SongName,
			Check:    check,
		}
		respCovers = append(respCovers, respCover)
	}

	//缓存
	jsondata, _ := json.Marshal(respCovers)
	cache := string(jsondata)
	err = sandwich.CacheDailyRank("all", cache)
	if err != nil {
		log.Fatal("redis缓存出错")
	}

	ctx.JSON(200, respCovers)
}

// GET /healing/dailyrank/:date
func GetDailyrank(ctx *gin.Context) {
	date := ctx.Param("date")

	//参数检查
	model := regexp.MustCompile("[0-1][0-9]-[0-3][0-9]")
	if !model.MatchString(date) {
		ctx.JSON(400, e.ErrMsgResponse{Message: "查询参数出错"})
		return
	}

	// //提取缓存
	// tempdata := sandwich.GetDailyRankByDate(date)
	// var respCovers []resp.HotResp
	// if tempdata != "" {
	// 	data := []byte(tempdata)
	// 	err := json.Unmarshal(data, &respCovers)
	// 	if err == nil {
	// 		ctx.JSON(200, respCovers)
	// 		return
	// 	} else {
	// 		log.Fatal("日榜redis缓存读取错误")
	// 	}
	// }

	//没有缓存，读取数据库
	raws, likes, err := dao.GetCoversByDate(date)
	respCovers := make([]resp.HotResp, 0)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		return
	}
	for i, cover := range raws {
		nickname, err := dao.GetUserNickname(cover.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
			return
		}
		respCover := resp.HotResp{
			CoverId:  likes[i].CoverId,
			Avatar:   cover.Avatar,
			Nickname: nickname,
			Posttime: cover.CreatedAt.String(),
			Songname: cover.SongName,
			Likes:    likes[i].Likes,
		}
		respCovers = append(respCovers, respCover)
	}

	//缓存
	jsondata, _ := json.Marshal(respCovers)
	cache := string(jsondata)
	err = sandwich.CacheDailyRank(date, cache)
	if err != nil {
		log.Fatal("redis缓存出错")
	}

	ctx.JSON(200, respCovers)
}

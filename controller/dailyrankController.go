package controller

//dailyrank除了当日热榜需要更新之外其它可以直接缓存
import (
	"regexp"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

// GET /healing/dailyrank/all
func GetAllrank(ctx *gin.Context) {
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
		respCover := resp.HotResp{
			CoverId:  likes[i].CoverId,
			Avatar:   cover.Avatar,
			Nickname: nickname,
			Posttime: cover.CreatedAt,
			Songname: cover.SongName,
			Likes:    likes[i].Likes,
		}
		respCovers = append(respCovers, respCover)
	}
	ctx.JSON(200, respCovers)
}

// GET /healing/dailyrank/:date
func GetDailyrank(ctx *gin.Context) {
	date := ctx.Param("date")
	model := regexp.MustCompile("[0-1][0-9]-[0-3][0-9]")
	if !model.MatchString(date) {
		ctx.JSON(403, e.ErrMsgResponse{Message: "查询参数出错"})
		return
	}
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
			Posttime: cover.CreatedAt,
			Songname: cover.SongName,
			Likes:    likes[i].Likes,
		}
		respCovers = append(respCovers, respCover)
	}
	ctx.JSON(200, respCovers)
}

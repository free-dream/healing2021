package controller

import (
	"regexp"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

// GET /healing/dailyrank/all
func GetAllrank(ctx *gin.Context) {
	date := ctx.Param("date")
	model := regexp.MustCompile("[0-1][0-9]-[0-3][0-9]")
	if !model.MatchString(date) {
		ctx.JSON(403, e.ErrMsgResponse{Message: "查询参数出错"})
	}
	raws, err := dao.GetCoversByDate(date)
	respCovers := make([]resp.HotResp, 10)
	// errHandler(err)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, cover := range raws {
		nickname, err := dao.GetUserNickname(cover.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		respCover := resp.HotResp{
			Avatar:   cover.Avatar,
			Nickname: nickname,
			Posttime: cover.CreatedAt,
			Songname: cover.SongName,
		}
		respCovers = append(respCovers, respCover)
	}
	ctx.JSON(200, respCovers)
}

// GET /healing/dailyrank/:date
func GetDailyrank(ctx *gin.Context) {
	date := ctx.Param("date")
	raws, err := dao.GetCoversByDate(date)
	respCovers := make([]resp.HotResp, 10)
	// errHandler(err)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, cover := range raws {
		nickname, err := dao.GetUserNickname(cover.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		respCover := resp.HotResp{
			Avatar:   cover.Avatar,
			Nickname: nickname,
			Posttime: cover.CreatedAt,
			Songname: cover.SongName,
		}
		respCovers = append(respCovers, respCover)
	}
	ctx.JSON(200, respCovers)
}

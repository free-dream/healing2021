package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

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
		errHandler(err)
		respCover := new(resp.HotResp)
		respCover.Avatar = cover.Avatar
		respCover.Likes = cover.Likes
		respCover.Nickname = nickname
		respCover.Posttime = cover.CreatedAt
		respCover.Songname = cover.SongName
		respCovers = append(respCovers, *respCover)
	}
	ctx.JSON(200, respCovers)
}

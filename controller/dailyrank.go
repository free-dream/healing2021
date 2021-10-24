package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

// GET /healing/dailyrank/:date
func GetDailyrank(ctx *gin.Context) {
	date := ctx.Param("date")
	raws, err := dao.GetCoversByDate(date)
	respCovers := make([]resp.HotResp, 10)
	errHandler(err)
	for _, cover := range raws {
		respCover := new(resp.HotResp)
		respCover.Avatar = cover.Avatar
		respCover.Likes = cover.Likes
		respCover.Nickname = ""
		respCover.Posttime = cover.CreatedAt
		respCover.Songname = cover.SongName
		respCovers = append(respCovers, *respCover)
	}
	ctx.JSON(200, respCovers)
}

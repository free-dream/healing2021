package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-gonic/gin"
)

// 推荐歌曲，根据click数降序获取10项(大家都在听)
func GetRank(ctx *gin.Context) {
	RankResp, err := dao.GetTop10()
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
		return
	}
	ctx.JSON(200, RankResp)
}

// 获取歌曲列表
func GetList(ctx *gin.Context) {
	ListResp, err := dao.GetLIst()
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
		return
	}
	ctx.JSON(200, ListResp)
}


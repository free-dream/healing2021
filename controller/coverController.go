package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-gonic/gin"
)

type PlayerParam struct {
	Jump    int `json:"jump"`
	Check   int `json:"check"`
	CoverId int `json:"cover_id"`
}

// 加载歌曲(翻唱),歌曲页
func LoadSongs(ctx *gin.Context) {
	var Param PlayerParam
	err := ctx.ShouldBind(&Param)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数不完整"})
		return
	}

	if Param.Check == 0 {
		Player, err := dao.GetPlayerNormal(Param.Jump, Param.CoverId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
			return
		}
		ctx.JSON(200, Player)
		return
	} else if Param.Check == 1{
		Player, err := dao.GetPlayerChild(Param.Jump, Param.CoverId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
			return
		}
		ctx.JSON(200, Player)
		return
	}
	ctx.JSON(403, e.ErrMsgResponse{Message: "Check参数错误"})
}

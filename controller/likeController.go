package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LikeParams struct {
	Todo int `json:"todo"`
	Type int `json:"type"`
	Id   int `json:"id"`
}

// 全局统一的点赞操作接口
func Like(ctx *gin.Context) {
	LikeParam := LikeParams{}
	err := ctx.ShouldBind(&LikeParam)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	// 参数准备
	UserId := sessions.Default(ctx).Get("user_id").(int)
	Type := ""
	if LikeParam.Type == 1 {
		Type = "moment"
	} else if LikeParam.Type == 2 {
		Type = "momentcomment"
	} else if LikeParam.Type == 3 {
		Type = "cover"
	} else {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	ok := dao.UpdateLikesByID(UserId, LikeParam.Id, LikeParam.Todo, Type)
	if ok != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "点赞成功"})
	return
}

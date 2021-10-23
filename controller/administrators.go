package controller

import (
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

type DeleteParams struct {
	Type string `json:"type"`
	Id   int    `json:"id"`
}

func DeleteMessage(ctx *gin.Context)  {
	// 参数获取
	param := DeleteParams{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	if UserId > 5{
		ctx.JSON(403, e.ErrMsgResponse{Message: "当前用户并非管理员"})
		return
	}

	// 分模式进行删除内容
	if param.Type != "momentcommont" && param.Type != "moment"{
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数错误"})
		return
	}


	ctx.JSON(200, e.ErrMsgResponse{Message: "删除操作成功"})
}
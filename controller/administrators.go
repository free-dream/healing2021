package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

type DeleteParams struct {
	Type int `json:"type"`
	Id   int `json:"id"`
}

func DeleteContent(ctx *gin.Context) {
	// 参数获取
	param := DeleteParams{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	// 不妨将 userid 小于5的账号预留，充当管理员
	if UserId > 5 {
		ctx.JSON(403, e.ErrMsgResponse{Message: "当前用户并非管理员"})
		return
	}

	// 分模式进行删除内容
	if param.Type != 1 && param.Type != 2 {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数错误"})
		return
	}
	err = dao.DeleteContent(param.Type, param.Id)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "数据库操作失败"})
		return
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "删除操作成功"})
}

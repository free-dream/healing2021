package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteContent(ctx *gin.Context) {
	// 参数获取
	param, bl := ctx.GetQuery("id")
	if !bl {
		ctx.JSON(400, e.ErrMsgResponse{Message: "参数非法"})
		return
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(400, e.ErrMsgResponse{Message: "error param"})
	}
	err = dao.DeleteContent(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, e.ErrMsgResponse{Message: "删除操作成功"})
}

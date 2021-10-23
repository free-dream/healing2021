package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-gonic/gin"
)

// 获取原唱相关信息
func GetOriginalInfo(ctx *gin.Context) {
	Name := ctx.Param("name")

	OriginInfo, err := dao.GetOriginInfo(Name)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
	}
	ctx.JSON(200, OriginInfo)
	return
}

// 获取用户翻唱列表并排序
func GetOriginalSingerList(ctx *gin.Context) {
	Name := ctx.Param("name")

	CoverList, err := dao.GetCoverList(Name)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
	}
	ctx.JSON(200, CoverList)
	return
}

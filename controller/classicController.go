package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// 获取原唱相关信息
func GetOriginalInfo(ctx *gin.Context) {
	ClassicIdStr := ctx.Query("classic_id")
	ClassicId, err := strconv.Atoi(ClassicIdStr)
	if err != nil {
		ctx.JSON(400, e.ErrMsgResponse{Message: "传入参数非法"})
		return
	}

	OriginInfo, err := dao.GetOriginInfo(ClassicId)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
	}
	ctx.JSON(200, OriginInfo)
	return
}

// 获取用户翻唱列表并排序
func GetOriginalSingerList(ctx *gin.Context) {
	ClassicIdStr := ctx.Query("classic_id")
	ClassicId, err := strconv.Atoi(ClassicIdStr)
	if err != nil {
		ctx.JSON(400, e.ErrMsgResponse{Message: "传入参数非法"})
		return
	}

	UserId := sessions.Default(ctx).Get("user_id").(int) // 获取当前用户 id
	CoverList, err := dao.GetCoverList(UserId, ClassicId)
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(500, e.ErrMsgResponse{Message: "未找到相关记录"})
		return
	}
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败"})
		return
	}
	ctx.JSON(200, CoverList)
	return
}

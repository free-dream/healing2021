package controller

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/task"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 治愈详情页，返回相关歌曲信息
func HealingPageFetcher(ctx *gin.Context) {
	param, verify := ctx.GetQuery("selectionId")
	//resp:=make(map[string]interface{})
	if !verify {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	selectionId, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
		// return
	}
	resp, err := dao.GetHealingPage(selectionId)
	if err != nil {
		panic(err)
		// return
	}

	ctx.JSON(200, gin.H{
		"resp": resp,
	})
}

//获取广告
func AdsPlayer(ctx *gin.Context) {
	resp, err := dao.GetAds()
	if err != nil {
		panic(err)
		// return
	}
	ctx.JSON(200, resp)

}

//点歌接口
//@@@@@@@任务模块已接入此接口@@@@@@@
func Selector(ctx *gin.Context) {
	param := statements.Selection{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		panic(err)
		// return
	}
	userid := sessions.Default(ctx).Get("user_id").(int)

	param.UserId = userid
	resp, err := dao.Select(param)
	if err != nil {
		panic(err)
		// return
	}
	//接入任务模块
	thistask := task.ST
	thistask.AddRecord(userid)
	//
	ctx.JSON(200, resp)
}

//首页控制
func SelectionFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	id := sessions.Default(ctx).Get("user_id").(int)
	if tag.Page == 1 {

		resp, err := dao.GetSelections(strconv.Itoa(1), id, tag)
		if err != nil {
			panic(err)

		}
		ctx.JSON(200, resp)

	} else {
		resp, err := dao.Pager("home"+strconv.Itoa(id), tag.Page)
		if err != nil {
			panic(err)

		}
		ctx.JSON(200, resp)
	}

}
func CoverFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	id := sessions.Default(ctx).Get("user_id").(int)
	if tag.Page == 1 {
		resp, err := dao.GetCovers(strconv.Itoa(1), id, tag)
		if err != nil {
			panic(err)
			// return
		}
		ctx.JSON(200, resp)

	} else {
		resp, err := dao.Pager("home"+strconv.Itoa(id), tag.Page)
		if err != nil {
			panic(err)
			// return
		}
		ctx.JSON(200, resp)
	}

}

type RecordParams struct {
	SelectionId string   `json:"selection_id" binding:"required"`
	Record      []string `json:"record" binding:"required"`
	Module      int      `json:"module"`
}

//唱歌接口
//@@@@@@@任务模块已植入此接口@@@@@@@
func Recorder(ctx *gin.Context) {
	params := RecordParams{}
	userID := sessions.Default(ctx).Get("user_id").(int)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	url, err := convertMediaIdArrToQiniuUrl(params.Record)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}

	resp, err := dao.CreateRecord(params.Module, params.SelectionId, url, userID)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	//任务模块植入 2021.11.1
	thistask := task.HT
	thistask.AddRecord(userID)
	//
	ctx.JSON(200, resp)
}

package controller

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/models/statements"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 治愈详情页，返回相关歌曲信息
func HealingPageFetcher(ctx *gin.Context) {
	param, verify := ctx.GetQuery("selectionId")
	//resp:=make(map[string]interface{})
	if !verify {
		ctx.JSON(401, gin.H{
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
func Selector(ctx *gin.Context) {
	param := statements.Selection{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		panic(err)
		// return
	}
	param.UserId = sessions.Default(ctx).Get("user_id").(int)
	param.Module = 1
	resp, err := dao.Select(param)
	if err != nil {
		panic(err)
		// return
	}
	ctx.JSON(200, resp)
}

//首页控制
func SelectionFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(401, gin.H{
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
		ctx.JSON(401, gin.H{
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

func Recorder(c *gin.Context) {
	params := RecordParams{}
	userID := tools.GetUser(c).ID
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	url, err := convertMediaIdArrToQiniuUrl(params.Record)
	if err != nil {
		c.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}

	resp, err := dao.CreateRecord(params.Module, params.SelectionId, url, int(userID))
	if err != nil {
		c.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}

	c.JSON(200, resp)
}

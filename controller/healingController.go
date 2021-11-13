package controller

import (
	"strconv"

	"git.100steps.top/100steps/healing2021_be/controller/ws"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"

	"git.100steps.top/100steps/healing2021_be/controller/task"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
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
	id := sessions.Default(ctx).Get("user_id").(int)
	selectionId, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
		// return
	}
	resp, err := dao.GetHealingPage(selectionId, id)
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
//----------任务模块已接入此接口----------
func Selector(ctx *gin.Context) {
	param := statements.Selection{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		panic(err)
		// return
	}
	userid := sessions.Default(ctx).Get("user_id").(int)

	param.UserId = userid
	num, resp, err := dao.Select(param)
	if err != nil {
		ctx.JSON(403, gin.H{
			"message": "今日次数已用完",
		})
		return
	}
	//接入任务模块
	thistask := task.ST
	thistask.AddRecord(userid)
	//
	ctx.JSON(200, gin.H{
		"selection_num": num,
		"resp":          resp,
	})
}

//首页控制
func SelectionFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	var err1 error
	tag.Page, err1 = strconv.Atoi(ctx.Query("page"))
	if err1 != nil || tag.Page <= 0 {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	tag.RankWay, _ = strconv.Atoi(ctx.Query("rankWay"))
	tag.Label = ctx.Query("label")
	/**if err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}*/
	id := sessions.Default(ctx).Get("user_id").(int)
	if tag.Page == 1 {

		resp, err := dao.GetSelections(id, tag)
		if err != nil {
			ctx.JSON(416, gin.H{
				"message": "out of range",
			})
			return

		}
		ctx.JSON(200, resp)

	} else {
		resp, err := dao.Pager("healing2021:home."+strconv.Itoa(id), tag.Page)
		if err != nil {
			ctx.JSON(416, gin.H{
				"message": "out of range",
			})
			return

		}
		ctx.JSON(200, resp)
	}

}
func CoverFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	var err1 error
	tag.Page, err1 = strconv.Atoi(ctx.Query("page"))
	if err1 != nil || tag.Page <= 0 {
		ctx.JSON(400, gin.H{
			"message": "error param",
		})
		return
	}
	tag.RankWay, _ = strconv.Atoi(ctx.Query("rankWay"))
	tag.Label = ctx.Query("label")
	id := sessions.Default(ctx).Get("user_id").(int)
	if tag.Page == 1 {
		//传入userid
		resp, err := dao.GetCovers(id, tag)
		//
		if err != nil {
			ctx.JSON(416, gin.H{
				"message": "out of range",
			})
			return

		}
		ctx.JSON(200, resp)

	} else {
		resp, err := dao.Pager("healing2021:home."+strconv.Itoa(id), tag.Page)
		if err != nil {
			ctx.JSON(416, gin.H{
				"message": "out of range",
			})
			return

		}
		ctx.JSON(200, resp)
	}

}

type RecordParams struct {
	SelectionId int      `json:"selection_id" binding:"required"`
	Record      []string `json:"record" binding:"required"`
	Module      int      `json:"module"`
	IsAnon      bool     `json:"is_anon"`
}

//唱歌接口
//----------任务模块已植入此接口----------
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
	id, resp, err := dao.CreateRecord(params.Module, params.SelectionId, url, userID, params.IsAnon)
	//推送到点歌用户
	conn := ws.GetConn()
	usrMsg := respModel.UsrMsg{}
	usrMsg.Url = url
	usrMsg.Song = resp.SongName
	usrMsg.SongId = uint(params.SelectionId)
	usrMsg.Message = ""
	usrMsg.FromUser = uint(userID)
	usrMsg.ToUser = uint(id)
	err = conn.SendUsrMsg(usrMsg)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	//任务模块植入 2021.11.1
	thistask := task.HT
	err = thistask.AddRecord(userID)

	ctx.JSON(200, resp)
}

//献唱接口
func DevotionPlayer(ctx *gin.Context) {
	resp, err := dao.PlayDevotion(sessions.Default(ctx).Get("user_id").(int))
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, resp)
}

package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/sandwich"

	// "git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	PRIZES = 100 //暂定
	DRAW   = 200 //抽奖的代价

	//设计奖项概率
	PRIZE1 = 0.01
	PRIZE2 = 0.04
	PRIZE3 = 0.1
)

/*
索引时按概率大小排序,小--大
先在redis里找，没有再索引数据库
*/
//GET /healing/lotterybox/lotteries
func GetLotteries(ctx *gin.Context) {
	lotteries := make([]resp.LotteryResp, 0)
	raws, err := dao.GetAllLotteries()
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, raw := range raws {
		lottery := resp.LotteryResp{
			Name:        raw.Name,
			Picture:     raw.Picture,
			Possibility: float32(raw.Possibility),
		}
		lotteries = append(lotteries, lottery)
	}
	ctx.JSON(200, lotteries)
}

type draws struct {
	Tel string `json:"tel"`
}

//POST /healing/lotterybox/draw
func Draw(ctx *gin.Context) {
	userid := sessions.Default(ctx).Get("user_id").(int)

	//确认是否已经抽奖
	check, err := dao.DrawCheck(userid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		return
	}
	var msg *resp.DrawResp = nil
	switch check {
	case 0:
		msg = &resp.DrawResp{
			Msg: resp.Msg0,
		}
	case 1:
		msg = &resp.DrawResp{
			Msg: resp.Msg1,
		}
		ctx.JSON(200, msg)
		return
	}

	//提取手机号
	ret := new(draws)
	err = ctx.ShouldBindJSON(ret)
	if err != nil {
		ctx.JSON(e.INVALID_PARAMS, e.ErrMsgResponse{Message: e.GetMsg(400)})
		return
	}
	err = dao.CreatePrize(userid, ret.Tel)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	if msg != nil {
		ctx.JSON(200, *msg)
		return
	}
	ctx.JSON(200, e.ErrMsgResponse{Message: e.GetMsg(200)})
}

// //GET /healing/lotterybox/drawcheck
// func DrawCheck(ctx *gin.Context) {
// 	userid := sessions.Default(ctx).Get("user_id").(int)
// 	check, err := dao.DrawCheck(userid)
// 	if err != nil {
// 		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
// 		return
// 	}
// 	var msg resp.DrawResp
// 	if check == 0 {
// 		msg = resp.DrawResp{
// 			Msg: resp.Msg0,
// 		}
// 	} else if check == 1 {
// 		msg = resp.DrawResp{
// 			Msg: resp.Msg1,
// 		}
// 	} else {
// 		msg = resp.DrawResp{
// 			Msg: resp.Msg2,
// 		}
// 	}
// 	ctx.JSON(200, msg)
// }

//GET /healing/lotterybox/points
func GetUserPoints(ctx *gin.Context) {
	points := resp.PointsResp{}
	userid := sessions.Default(ctx).Get("user_id").(int)
	data, err := dao.GetUserPoints(userid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	points.Points = data
	ctx.JSON(200, points)
}

//GET /healing/lotterybox/tasktable
func GetTasktable(ctx *gin.Context) {
	//返回结构体
	respTasks := make([]resp.TaskTableResp, 0)
	//userid
	userid := sessions.Default(ctx).Get("user_id").(int)
	//获取任务表
	task_table, err := dao.GetTasktables(userid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	//读取任务信息并拼合成对应数据体系
	for _, table := range task_table {
		//获取对应task
		//生成任务返回表
		tasktableresp := resp.TaskTableResp{
			Counter: sandwich.GetCacheTaskPoints(userid, table.ID),
			Task:    table,
		}
		respTasks = append(respTasks, tasktableresp)
	}
	ctx.JSON(200, respTasks)
}

// //GET /healing/lotterybox/prizes
// func GetPrizes(ctx *gin.Context) {
// 	prizes := make([]resp.PrizesResp, 10)
// 	//获取userid
// 	userid := sessions.Default(ctx).Get("user_id").(int)
// 	//获取中奖数据
// 	raws, err := dao.GetPrizesById(userid)
// 	if err != nil {
// 		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
// 	}
// 	for _, prize := range raws {
// 		res := new(resp.PrizesResp)
// 		res.Name = prize.Name
// 		res.Picture = prize.Picture
// 		prizes = append(prizes, *res)
// 	}
// 	ctx.JSON(200, prizes)
// }

// //抽奖算法，落到最后分区直接没中
// func methods(possibilities ...float64) int {
// 	base := 0
// 	phase := make([]int, 5)
// 	phase = append(phase, 0)
// 	for _, possibility := range possibilities {
// 		base += int(possibility * 500)
// 		phase = append(phase, base)
// 	}
// 	phase = append(phase, 500)
// 	draw := tools.GetRandomNumbers(500) + 1
// 	var i int = 1
// 	for ; i < len(phase); i++ {
// 		if draw <= phase[i] && phase[i-1] < draw {
// 			break
// 		}
// 	}
// 	//落到最后一个分区就是不中，其它的直接返回对应id,另行确认奖品归属
// 	if i == len(phase)-1 {
// 		return -1
// 	} else {
// 		return draw
// 	}
// }

//线上抽奖，废案
// func draw(ctx *gin.Context) {
// 	//返回格式
// 	drawResp := resp.DrawResp{}
// 	//获取userid和points
// 	userid := sessions.Default(ctx).Get("user_id").(int)
// 	points, err := dao.GetPoints(userid)
// 	if err != nil {
// 		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
// 	}
// 	//判断是否可抽
// 	if points < DRAW {
// 		drawResp.Check = 2
// 		ctx.JSON(200, drawResp)
// 		return
// 	}
// 	//抽卡，中或不中
// 	prizeid := methods()
// 	//不中
// 	if prizeid < 0 {
// 		drawResp.Check = 0
// 		ctx.JSON(200, drawResp)
// 		return
// 	}
// 	//中
// 	result, err := dao.Draw(prizeid)
// 	ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
// 	drawResp.Check = 1
// 	drawResp.Name = result.Name
// 	drawResp.Picture = result.Picture
// 	ctx.JSON(200, drawResp)
// }

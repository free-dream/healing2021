package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//错误码补全

const (
	PRIZES = 100 //暂定
	DRAW   = 200 //抽奖的代价

)

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

//抽奖算法
//在等奖品的安排和可能性
func methods(possibilities ...float64) int {
	base := 0
	phase := make([]int, 10)
	for _, possibility := range possibilities {
		base += int(possibility * 1000)
		phase = append(phase, base)
	}
	phase = append(phase, 1000)
	draw := tools.GetRandomNumbers(1000)
	var i int = 1
	for ; i < len(phase); i++ {
		if draw < phase[i] && phase[i-1] < draw {
			break
		}
	}
	switch i {

	}
	return -1
}

/*
索引时按概率大小排序,小--大
先在redis里找，没有再索引数据库
GET /healing/lotterybox/lotteries
*/
func GetLotteries(ctx *gin.Context) {
	//从数据库里调，redis爆炸的备选方案
	lotteries := make([]resp.LotteryResp, 10)
	raws, err := dao.GetAllLotteries()
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, raw := range raws {
		lottery := resp.LotteryResp{
			Name:        raw.Name,
			Picture:     raw.Picture,
			Possibility: int(raw.Possibility),
		}
		lotteries = append(lotteries, lottery)
	}
	ctx.JSON(200, lotteries)
}

//GET /healing/lotterybox/draw
//先写mysql的版本,防爆,暂时没有redis
func Draw(ctx *gin.Context) {
	//返回格式
	drawResp := resp.DrawResp{}
	//获取userid和points
	userid := sessions.Default(ctx).Get("user_id").(int)
	points, err := dao.GetPoints(userid)
	errHandler(err)
	//判断是否可抽
	if points < DRAW {
		drawResp.Check = 2
		ctx.JSON(200, drawResp)
		return
	}
	//抽卡，中或不中
	prizeid := methods()
	//不中
	if prizeid < 0 {
		drawResp.Check = 0
		ctx.JSON(200, drawResp)
		return
	}
	//中
	result, err := dao.Draw(prizeid)
	ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	drawResp.Check = 1
	drawResp.Name = result.Name
	drawResp.Picture = result.Picture
	ctx.JSON(200, drawResp)
}

//GET /healing/lotterybox/prizes
//同理，先写mysql版本的防爆
func GetPrizes(ctx *gin.Context) {
	prizes := make([]resp.PrizesResp, 10)
	//获取openid和userid
	// openid := tools.GetOpenid(ctx)
	userid := sessions.Default(ctx).Get("user_id").(int)
	//获取中奖数据
	raws, err := dao.GetPrizesById(userid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, prize := range raws {
		res := new(resp.PrizesResp)
		res.Name = prize.Name
		res.Picture = prize.Picture
		prizes = append(prizes, *res)
	}
	ctx.JSON(200, prizes)
}

//GET /healing/lotterybox/tasktable
func GetTasktable(ctx *gin.Context) {
	//返回结构体
	respTasks := make([]resp.TaskTableResp, 10)
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
		taskid := table.TaskId
		task, err := dao.GetTasks(taskid)
		errHandler(err)
		//生成任务返回
		taskresp := resp.TaskResp{
			ID:     taskid,
			Target: task.Target,
			Text:   task.Text,
		}
		//生成任务返回表
		tasktableresp := resp.TaskTableResp{
			Check:   table.Check,
			Counter: table.Counter,
			Task:    taskresp,
		}

		respTasks = append(respTasks, tasktableresp)
	}
	ctx.JSON(200, respTasks)
}

//更新任务状态，视情况更新用户积分数,由于任务未定，这里建议单独拉出来
//POST /healing/lotterybox/task
func UpdateTask(ctx *gin.Context) {}

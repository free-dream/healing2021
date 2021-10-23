package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

const (
	PRIZES = 100 //暂定
	DRAW   = 200 //抽奖的代价
)

var (
	openid string = "test"
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
	errHandler(err)
	for _, raw := range raws {
		lottery := new(resp.LotteryResp)
		lottery.Name = raw.Name
		lottery.Picture = raw.Picture
		lottery.Possibility = int(raw.Possibility)
		lotteries = append(lotteries, *lottery)
	}
	ctx.JSON(200, lotteries)
}

//GET /healing/lotterybox/draw
//先写mysql的版本,防爆,暂时没有写回mysql
func Draw(ctx *gin.Context) {
	//返回格式
	drawResp := new(respModel.DrawResp)
	//获取openid和userid和points
	openid := tools.GetOpenid(ctx)
	userid, err := dao.GetUserid(openid)
	errHandler(err)
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
	}
	//中
	result, err := dao.Draw(prizeid)
	errHandler(err)
	drawResp.Check = 1
	drawResp.Name = result.Name
	drawResp.Picture = result.Picture
	ctx.JSON(200, drawResp)
}

//GET /healing/lotterybox/prizes
//同理，先写mysql版本的防爆
func GetPrizes(ctx *gin.Context) {
	prizes := make([]respModel.PrizesResp, 10)
	//获取openid和userid
	openid := tools.GetOpenid(ctx)
	userid, err := dao.GetUserid(openid)
	errHandler(err)
	//获取中奖数据
	raws, err := dao.GetPrizesById(userid)
	errHandler(err)
	for _, prize := range raws {
		res := new(respModel.PrizesResp)
		res.Name = prize.Name
		res.Picture = prize.Picture
		prizes = append(prizes, *res)
	}
	ctx.JSON(200, prizes)
}

//GET /healing/lotterybox/tasktable
func GetTasktable(ctx *gin.Context) {}

//更新任务状态，视情况更新用户积分数
//POST /healing/lotterybox/task
func UpdateTask(ctx *gin.Context) {}

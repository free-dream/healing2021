package tradition

import (
	"fmt"

	"git.100steps.top/100steps/healing2021_be/dao"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

const (
	PRIZES = 100 //暂定
)

func test() {
	fmt.Println("test")
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

//抽奖算法
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
	return 0
}

//获取用户积分数
func getPoints() int {
	return 0
}

//索引时按概率大小排序,小--大
//先在redis里找，没有再索引数据库
//GET /healing/lotterybox/lotteries
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

//POST /healing/lotterybox/draw
func Draw(ctx *gin.Context) {
	id := methods()
	result, err := dao.Draw(id)
	errHandler(err)
	ctx.JSON(200, result)
}

//GET /healing/lotterybox/prizes
func GetPrizes(ctx *gin.Context) {}

//GET /healing/lotterybox/tasktable
func GetTasktable(ctx *gin.Context) {}

//更新任务状态，视情况更新用户积分数
//POST /healing/lotterybox/task
func UpdateTask(ctx *gin.Context) {}

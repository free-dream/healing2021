package tradition

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test() {
	fmt.Println("test")
}

//抽奖算法
func methods() {}

//获取用户积分数
func getPoints() int {
	return 0
}

//索引时按概率大小排序
//GET /healing/lotterybox/lotteries
func GetLotteries(ctx *gin.Context) {}

//POST /healing/lotterybox/draw
func Draw(ctx *gin.Context) {
	methods()
}

//GET /healing/lotterybox/prizes
func GetPrizes(ctx *gin.Context) {}

//GET /healing/lotterybox/tasktable
func GetTasktable(ctx *gin.Context) {}

//更新任务状态，视情况更新用户积分数
//POST /healing/lotterybox/task
func UpdateTask(ctx *gin.Context) {}

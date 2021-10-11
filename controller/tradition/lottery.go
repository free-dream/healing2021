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

//索引时按概率大小排序
//GET /healing/lotterybox/lotteries
func GetLotteries(ctx *gin.Context) {}

//POST /healing/lotterybox/draw

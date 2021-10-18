package tradition

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

//治愈详情页，返回相关歌曲信息
func HealingPageFetcher(ctx *gin.Context) {
	param, bool := ctx.GetQuery("selectionId")
	//resp:=make(map[string]interface{})
	if !bool {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	selectionId, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
		return
	}
	resp, err := dao.GetHealingPage(selectionId)
	if err != nil {
		panic(err)
		return
	}

	ctx.JSON(200, gin.H{
		"resp": resp,
	})
}

func AdsPlayer(ctx *gin.Context) {
	resp, err := dao.GetAds()
	if err != nil {
		panic(err)
		return
	}
	ctx.JSON(200, resp)

}

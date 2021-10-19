package tradition

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

//治愈详情页，返回相关歌曲信息
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

//获取广告
func AdsPlayer(ctx *gin.Context) {
	resp, err := dao.GetAds()
	if err != nil {
		panic(err)
		return
	}
	ctx.JSON(200, resp)

}

//点歌控制
func SelectionFetcher(ctx *gin.Context) {
	tag := dao.Tags{}
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	resp, err := dao.GetSelections(tag)
	if err != nil {
		panic(err)
		return
	}
	ctx.JSON(200, resp)

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
	resp, err := dao.GetCovers(tag)
	if err != nil {
		panic(err)
		return
	}
	ctx.JSON(200, resp)

}

//唱歌接口
//录音操作还在写
func Recorder(ctx *gin.Context) {
	/*selection_id,verify:=ctx.GetQuery("selection_id")
	if !verify{
		ctx.JSON(401,gin.H{
			"message":"error param",
		})

	}*/

}

//对经典治愈系的录音点赞
func LikePoster(ctx *gin.Context) {

}

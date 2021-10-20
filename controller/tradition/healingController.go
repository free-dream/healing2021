package tradition

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-contrib/sessions"
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

type RecordParams struct {
	SelectionId string   `json:"selection_id" binding:"required"`
	Record      []string `json:"record" binding:"required"`
}

//唱歌接口

func Recorder(c *gin.Context) {
	var params RecordParams
	userID := tools.GetUser(c).ID
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	fmt.Println(params.Record)
	url, err := convertMediaIdArrToQiniuUrl(params.Record)
	if err != nil {
		c.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}
	err = dao.CreateRecord(params.SelectionId, url, int(userID))
	if err != nil {
		c.JSON(403, e.ErrMsgResponse{Message: err.Error()})
		return
	}

	c.JSON(200, "OK")
}

//对经典治愈系的录音点赞
func LikePoster(ctx *gin.Context) {
	id, verify := ctx.GetQuery("covers_id")
	if !verify {
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	coverId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
		ctx.JSON(401, gin.H{
			"message": "error param",
		})
		return
	}
	session := sessions.Default(ctx)
	openid := session.Get("openid").(string)
	err = dao.Like(coverId, openid)
}

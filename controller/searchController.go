package controller

import (
	"regexp"
	"strings"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

type keyword struct {
	Keyword string `json:"keyword"`
}

//关键字清洗
//当关键字数超过3时，只取前三个
//为了保证性能，暂时只支持单个词的查询,下述方法雪藏
func clean(keyword string) []string {
	temp := strings.Split(keyword, "	")
	keywords := make([]string, 3)
	for _, word := range temp {
		if len(keywords) == 3 {
			break
		}
		if word != "" {
			keywords = append(keywords, word)
		}
	}
	return keywords
}

//电话号码检查
func IsTelSearch(keyword string) bool {
	model := regexp.MustCompile("1[3-6][0-9][0-9]{8}")
	return model.MatchString(keyword)
}

//POST /heaing/search
func Search(ctx *gin.Context) {
	var key keyword
	respAll := make([]interface{}, 0)
	respCovers := make([]respModel.CoversResp, 0)
	respSelections := make([]respModel.SelectionResp, 0)
	// respUsers := make([]respModel.UserResp, 0)
	respLen := new(respModel.SumResp)

	//提取关键字
	if err := ctx.ShouldBind(&key); err != nil {
		ctx.JSON(e.INVALID_PARAMS, e.ErrMsgResponse{Message: e.GetMsg(400)})
		return
	}
	keyword := key.Keyword

	// //特殊的字符处理
	// if keyword == "" {
	// 	respAll = append(respAll, respLen, respUsers, respSelections, respCovers)
	// 	ctx.JSON(200, respAll)
	// 	return
	// }

	//确认是否是电话查询
	if IsTelSearch(keyword) {
		rawUsers, lenuser, err := dao.SearchUserByTel(keyword)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		respLen.LenUser = lenuser
		respAll = append(respAll, respLen, rawUsers, respSelections, respCovers)
		ctx.JSON(200, respAll)
		return
	}

	//查询翻唱
	rawCovers, lencover, err := dao.SearchCoverByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenCover = lencover

	//查询点歌
	rawSelections, lenselec, err := dao.SearchSelectionByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenSelection = lenselec

	//查询用户
	rawUsers, lenuser, err := dao.SearchUserByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenUser = lenuser

	respAll = append(respAll, respLen, rawUsers, rawSelections, rawCovers)
	ctx.JSON(200, respAll)
}

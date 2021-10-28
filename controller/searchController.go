package controller

import (
	"strings"

	"git.100steps.top/100steps/healing2021_be/dao"
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

//POST heaing/search
func Search(ctx *gin.Context) {
	var key keyword
	respAll := make([]interface{}, 4)
	respCovers := make([]respModel.CoversResp, 5)
	respSelections := make([]respModel.SelectionResp, 5)
	respUsers := make([]respModel.UserResp, 5)
	respLen := new(respModel.SumResp)
	if err := ctx.ShouldBind(&key); err != nil {
		panic(err)
	}
	keyword := key.Keyword

	//查询翻唱
	rawCovers, lencover, err := dao.SearchCoverByKeyword(keyword)
	errHandler(err)
	respLen.LenCover = lencover
	for _, cover := range rawCovers {
		temp := new(respModel.CoversResp)
		temp.Avatar = cover.Avatar
		temp.Coverid = int(cover.ID)
		nickname, err := dao.GetUserNickname(cover.UserId)
		errHandler(err)
		temp.Nickname = nickname
		temp.Posttime = cover.CreatedAt
		respCovers = append(respCovers, *temp)
	}

	//查询点歌
	rawSelections, lenselec, err := dao.SearchSelectionByKeyword(keyword)
	errHandler(err)
	respLen.LenSelection = lenselec
	for _, selection := range rawSelections {
		temp := new(respModel.SelectionResp)
		// temp.Avatar = selection.Avatar
		temp.Selectionid = int(selection.ID)
		nickname, err := dao.GetUserNickname(selection.UserId)
		errHandler(err)
		temp.Nickname = nickname
		temp.Posttime = selection.CreatedAt
		respSelections = append(respSelections, *temp)
	}

	//查询用户
	rawUsers, lenuser, err := dao.SearchUserByKeyword(keyword)
	errHandler(err)
	respLen.LenUser = lenuser
	for _, user := range rawUsers {
		temp := new(respModel.UserResp)
		temp.Avatar = user.Avatar
		temp.Userid = int(user.ID)
		temp.Nickname = user.Nickname
		temp.Slogan = user.Signature
		respUsers = append(respUsers, *temp)
	}

	respAll = append(respAll, respLen, respSelections, respCovers)
	ctx.JSON(200, respAll)
}

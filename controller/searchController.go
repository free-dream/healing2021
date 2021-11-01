package controller

import (
	"strings"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"github.com/gin-gonic/gin"
)

//构建一个词库，用于区分不同的关键词
var ()

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

	//提取关键字
	if err := ctx.ShouldBind(&key); err != nil {
		ctx.JSON(e.INVALID_PARAMS, e.ErrMsgResponse{Message: e.GetMsg(400)})
		return
	}
	keyword := key.Keyword

	//查询翻唱
	rawCovers, lencover, err := dao.SearchCoverByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenCover = lencover

	for _, cover := range rawCovers {
		nickname, err := dao.GetUserNickname(cover.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		temp := respModel.CoversResp{
			Avatar:   cover.Avatar,
			Coverid:  int(cover.ID),
			Nickname: nickname,
			Posttime: cover.CreatedAt,
		}
		respCovers = append(respCovers, temp)
	}

	//查询点歌
	rawSelections, lenselec, err := dao.SearchSelectionByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenSelection = lenselec
	for _, selection := range rawSelections {
		nickname, err := dao.GetUserNickname(selection.UserId)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		temp := respModel.SelectionResp{
			Selectionid: int(selection.ID),
			Nickname:    nickname,
			Posttime:    selection.CreatedAt,
		}
		respSelections = append(respSelections, temp)
	}

	//查询用户
	rawUsers, lenuser, err := dao.SearchUserByKeyword(keyword)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	respLen.LenUser = lenuser
	for _, user := range rawUsers {
		temp := respModel.UserResp{
			Avatar:   user.Avatar,
			Userid:   int(user.ID),
			Nickname: user.Nickname,
			Slogan:   user.Signature,
		}
		respUsers = append(respUsers, temp)
	}

	respAll = append(respAll, respLen, respSelections, respCovers)
	ctx.JSON(200, respAll)
}

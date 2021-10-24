package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

//GET /healing/rank/:school
func GetRanking(ctx *gin.Context) {
	//取出参数
	school := ctx.Param("school")
	//生成返回模块
	rankresps := make([]resp.RankingResp, 10)
	//提取数据
	raws, err := dao.GetRankingBySchool(school)
	errHandler(err)
	for _, user := range raws {
		temp := new(resp.RankingResp)
		temp.Avatar = user.Avatar
		temp.Nickname = user.Nickname
		rankresps = append(rankresps, *temp)
	}
	ctx.JSON(200, rankresps)
}

//GET /healing/rank/user
//获取用户当前排名
func GetMyRank(ctx *gin.Context) {
	//获取openid和userid
	openid := tools.GetOpenid(ctx)
	userid, err := dao.GetUserid(openid)
	errHandler(err)
}

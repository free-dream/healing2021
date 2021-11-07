package controller

import (
	"encoding/json"
	"log"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//GET /healing/rank/:school
func GetRanking(ctx *gin.Context) {
	//取出参数
	school := ctx.Param("school")

	//读取可能有的redis缓存,若有直接返回
	var rankresps []resp.RankingResp
	datatemp := sandwich.GetPointsRanking(school)
	if datatemp != "" {
		data := []byte(datatemp)
		err := json.Unmarshal(data, &rankresps)
		if err == nil {
			ctx.JSON(200, rankresps)
			return
		} else {
			log.Fatal("pointsrank缓存数据格式化失败")
		}
	}

	//生成返回模块
	rankresps = make([]resp.RankingResp, 0)
	//提取数据
	raws, err := dao.GetRankingBySchool(school)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
	}
	for _, user := range raws {
		temp := new(resp.RankingResp)
		temp.Userid = int(user.ID)
		temp.Avatar = user.Avatar
		temp.Nickname = user.Nickname
		rankresps = append(rankresps, *temp)
	}

	//缓存json化的数据,json化的报错无视
	//缓存出错不影响可用性，所以仅log
	jsondata, _ := json.Marshal(rankresps)
	cache := string(jsondata)
	err = sandwich.CachePointsRanking(school, cache)
	if err != nil {
		log.Fatal("redis缓存出错")
	}

	ctx.JSON(200, rankresps)
}

//GET /healing/rank/user
func GetMyRank(ctx *gin.Context) {
	//获取userid
	ruresp := resp.RankingUResp{}
	userid := sessions.Default(ctx).Get("user_id").(int)
	test := sandwich.GetCURanking(userid)
	if test == "" {
		rank, err := dao.GetRankByCUserId(userid)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作出错"})
		}
		//缓存，只log
		err = sandwich.CacheCURanking(userid, rank)
		if err != nil {
			log.Fatal("redis缓存出错")
		}
		ruresp.Rank = rank
		ctx.JSON(200, ruresp)
		return
	}
	ruresp.Rank = test
	ctx.JSON(200, ruresp)
}

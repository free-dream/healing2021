package playground

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-gonic/gin"
)

// 拉取广场动态列表
func GetMomentList(ctx *gin.Context) {
	AllComment, ok := dao.GetAllMoment()
	if !ok {
		ctx.JSON(403, e.ErrMsgResponse{Message: "动态不存在"})
		return
	}
	ctx.JSON(200, AllComment)
}

// 发布动态
func PostMoment(ctx *gin.Context) {

}

// 查看动态的详情
func GetMomentDetail(ctx *gin.Context) {

}

// 给动态添加评论
func PostComment(ctx *gin.Context) {

}

// 拉取动态的评论列表
func GetCommentList(ctx *gin.Context) {

}

// 给动态或评论点赞（取消点赞）
func PriseOrNot(ctx *gin.Context) {

}

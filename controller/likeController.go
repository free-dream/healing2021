package controller

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LikeParams struct {
	Todo int `json:"todo"`
	Type int `json:"type"`
	Id   int `json:"id"`
}

// 全局统一的点赞操作接口
func Like(ctx *gin.Context) {
	LikeParam := LikeParams{}
	err := ctx.ShouldBind(&LikeParam)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	// 参数准备
	UserId := sessions.Default(ctx).Get("user_id").(int)
	Type := ""
	if LikeParam.Type == 1 {
		Type = "moment"
	} else if LikeParam.Type == 2 {
		Type = "momentcomment"
	} else if LikeParam.Type == 3 {
		Type = "cover"
	} else {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	// 写入点赞表
	ok := dao.UpdateLikesByID(UserId, LikeParam.Id, LikeParam.Todo, Type)
	if ok != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入错误"})
		return
	}

	// 发送对应的系统消息
	// 发送相应的系统消息[有 实际评论写入成功，但是系统消息发送失败 的不一致风险]
	//conn := ws.GetConn()
	//userId, err := dao.GetMomentSenderId(NewComment.DynamicsId)
	//if err != nil {
	//	ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	//	return
	//}
	//err = conn.SendSystemMsg(respModel.SysMsg{
	//	Uid: uint(userId),
	//	Type:
	//	ContentId:
	//	Song:
	//	Time: time.Now(),
	//	IsSend:
	//})
	//if err != nil {
	//	ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	//	return
	//}

	ctx.JSON(200, e.ErrMsgResponse{Message: "操作成功"})
	return
}

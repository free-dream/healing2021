package controller

import (
	"time"

	"git.100steps.top/100steps/healing2021_be/controller/ws"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
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

	//防止空表
	if LikeParam.Todo == 0 {
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	// 参数准备
	UserId := sessions.Default(ctx).Get("user_id").(int)
	Type := ""
	switch LikeParam.Type {
	case 1:
		Type = "moment"
	case 2:
		Type = "momentcomment"
	case 3:
		Type = "cover"
	default:
		ctx.JSON(403, e.ErrMsgResponse{Message: "参数非法"})
		return
	}

	// 写入点赞表
	ok := dao.UpdateLikesByID(UserId, LikeParam.Id, LikeParam.Todo, Type)

	//重复点赞判断/数据库写入检查
	if ok != nil {
		if dao.IsLikesExistError(ok) {
			ctx.JSON(405, e.ErrMsgResponse{Message: "不允许重复点赞"})
			return
		} else if dao.IsZeroUserIdError(ok) {
			ctx.JSON(406, e.ErrMsgResponse{Message: "传入用户id为0"})
			return
		}
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入错误"})
		return
	}

	//发送相应的系统消息[有 实际评论写入成功，但是系统消息发送失败 的不一致风险]
	if LikeParam.Todo == 1 {
		conn := ws.GetConn()
		sysMsg := respModel.SysMsg{}

		switch Type {
		case "moment":
			SenderId, err := dao.GetMomentSenderId(LikeParam.Id)
			if err != nil {
				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
				return
			}
			//fmt.Println("====")
			//fmt.Println(SenderId, LikeParam.Id)
			//fmt.Println("====")
			sysMsg = respModel.SysMsg{
				Uid:       uint(SenderId),
				Type:      2,
				ContentId: uint(LikeParam.Id),
				Time:      time.Now(),
			}
		case "momentcomment":
			SenderId, err := dao.GetCommentSenderId(LikeParam.Id)
			if err != nil {
				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
				return
			}
			//fmt.Println("====")
			//fmt.Println(SenderId, LikeParam.Id)
			//fmt.Println("====")
			sysMsg = respModel.SysMsg{
				Uid:       uint(SenderId),
				Type:      4,
				ContentId: uint(LikeParam.Id),
				Time:      time.Now(),
			}
		case "cover":
			singerId, songName, err := dao.GetCoverInfo(LikeParam.Id)
			if err != nil {
				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
				return
			}
			//fmt.Println("====")
			//fmt.Println(singerId, songName)
			//fmt.Println("====")
			sysMsg = respModel.SysMsg{
				Uid:       uint(singerId),
				Type:      1,
				Song:      songName,
				ContentId: uint(LikeParam.Id),
				Time:      time.Now(),
			}
		}

		err = conn.SendSystemMsg(sysMsg)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
			return
		}
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "操作成功"})
	return
}

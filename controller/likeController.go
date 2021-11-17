package controller

import (
	"log"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LikeParams struct {
	Todo int `json:"todo"`
	Type int `json:"type"`
	Id   int `json:"id"`
}

// 全局统一的点赞操作接口
// redis确认完了就直接回复
// 拉两个协程分别发消息和更新数据库
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

	//把重复判断拉到这里
	//redis爆炸直接500挂掉,牺牲安全性但是保证效率
	check, _ := sandwich.Check(LikeParam.Id, Type, UserId)
	if LikeParam.Todo == -1 && check {
		err = sandwich.CancelLike(LikeParam.Id, Type, UserId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			ctx.JSON(200, e.ErrMsgResponse{Message: "ok"})
		}
	} else if LikeParam.Todo == 1 && !check {
		err = sandwich.AddLike(LikeParam.Id, Type, UserId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			ctx.JSON(200, e.ErrMsgResponse{Message: "ok"})
		}
	} else {
		ctx.JSON(405, e.ErrMsgResponse{Message: "不允许重复点赞"})
		return
	}

	//交给后台更新点赞表
	updatemsg := make([]interface{}, 0)
	updatemsg = append(updatemsg, UserId, LikeParam.Id, LikeParam.Todo, Type)
	updatelikechan <- updatemsg

	//扔给后台发送消息
	if LikeParam.Todo == 1 {
		//干脆缓存用户nickname,这个地方未来可能会成为性能瓶颈
		nickname, err := dao.GetUserNickname(UserId)
		//
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
			return
		}
		likemsg := make([]interface{}, 0)
		likemsg = append(likemsg, nickname, LikeParam.Id, Type)
		likemsgchan1 <- likemsg
	}
	// go func() {
	// 	defer wg.Done()
	// 	dao.UpdateLikesByID(UserId, LikeParam.Id, LikeParam.Todo, Type)
	// }()

	//发送相应的系统消息[有 实际评论写入成功，但是系统消息发送失败 的不一致风险]
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	nickname, err := dao.GetUserNickname(UserId)
	// 	if err != nil {
	// 		ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	// 		return
	// 	}

	// 	if LikeParam.Todo == 1 {
	// 		conn := ws.GetConn()
	// 		sysMsg := respModel.SysMsg{}

	// 		switch Type {
	// 		case "moment":
	// 			SenderId, err := dao.GetMomentSenderId(LikeParam.Id)
	// 			if err != nil {
	// 				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	// 				return
	// 			}
	// 			sysMsg = respModel.SysMsg{
	// 				Uid:       uint(SenderId),
	// 				Type:      2,
	// 				ContentId: uint(LikeParam.Id),
	// 				Time:      time.Now(),
	// 				FromUser:  nickname,
	// 			}
	// 		case "momentcomment":
	// 			SenderId, err := dao.GetCommentSenderId(LikeParam.Id)
	// 			if err != nil {
	// 				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	// 				return
	// 			}
	// 			sysMsg = respModel.SysMsg{
	// 				Uid:       uint(SenderId),
	// 				Type:      4,
	// 				ContentId: uint(LikeParam.Id),
	// 				Time:      time.Now(),
	// 				FromUser:  nickname,
	// 			}
	// 		case "cover":
	// 			singerId, songName, err := dao.GetCoverInfo(LikeParam.Id)
	// 			if err != nil {
	// 				ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	// 				return
	// 			}
	// 			sysMsg = respModel.SysMsg{
	// 				Uid:       uint(singerId),
	// 				Type:      1,
	// 				Song:      songName,
	// 				ContentId: uint(LikeParam.Id),
	// 				Time:      time.Now(),
	// 				FromUser:  nickname,
	// 			}
	// 		}

	// 		err = conn.SendSystemMsg(sysMsg)
	// 		if err != nil {
	// 			ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
	// 			return
	// 		}
	// 	}
	// }()

	// wg.Wait()
	// ctx.JSON(200, e.ErrMsgResponse{Message: "操作成功"})
	return
}

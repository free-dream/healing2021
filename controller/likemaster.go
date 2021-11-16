package controller

import (
	"log"
	"time"

	"git.100steps.top/100steps/healing2021_be/controller/ws"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
)

var (
	normaltasks    []int //声明不同任务表，以对用户做区分,目前仅有普通用户，假任务使用
	updatelikechan chan interface{}
	likemsgchan1   chan interface{}
	likemsgchan2   chan interface{}
	likemsgchan3   chan interface{}
)

func init() {
	normaltasks = []int{1, 2, 3} //任务初始化标注于中间件
	updatelikechan = make(chan interface{}, 30)
	likemsgchan1 = make(chan interface{}, 30)
	likemsgchan2 = make(chan interface{}, 30)
	likemsgchan3 = make(chan interface{}, 30)
}

//挂在后台专门用来处理like表更新和发like消息
//考虑到点赞是发信息的两倍，分开处理发信息和点赞
//这样的话高压下消息会有延迟，不过这是可以接受的

//发送消息专门拉一个函数出来
//消息发送失败就不回显了，直接log

func sendMsg(target int, Type string, nickname string) {
	conn := ws.GetConn()
	sysMsg := respModel.SysMsg{}
	var err error
	switch Type {
	case "moment":
		SenderId, err := dao.GetMomentSenderId(target)
		if err != nil {
			log.Printf("系统消息发送失败")
			return
		}
		sysMsg = respModel.SysMsg{
			Uid:       uint(SenderId),
			Type:      2,
			ContentId: uint(target),
			Time:      time.Now(),
			FromUser:  nickname,
		}
	case "momentcomment":
		SenderId, err := dao.GetCommentSenderId(target)
		if err != nil {
			log.Printf("系统消息发送失败")
			return
		}
		sysMsg = respModel.SysMsg{
			Uid:       uint(SenderId),
			Type:      4,
			ContentId: uint(target),
			Time:      time.Now(),
			FromUser:  nickname,
		}
	case "cover":
		singerId, songName, err := dao.GetCoverInfo(target)
		if err != nil {
			log.Printf("系统消息发送失败")
			return
		}
		sysMsg = respModel.SysMsg{
			Uid:       uint(singerId),
			Type:      1,
			Song:      songName,
			ContentId: uint(target),
			Time:      time.Now(),
			FromUser:  nickname,
		}
	}
	err = conn.SendSystemMsg(sysMsg)
	if err != nil {
		log.Printf("系统消息发送失败")
		return
	}
}

//必须顺序处理，否则有可能先取消后点赞
func LikeDaemon() {
	for {
		select {
		case data := <-updatelikechan:
			temp, _ := data.([]interface{})
			UserId, _ := temp[0].(int)
			Id, _ := temp[1].(int)
			Todo, _ := temp[2].(int)
			Type, _ := temp[3].(string)
			dao.UpdateLikesByID(UserId, Id, Todo, Type)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

//看起来好丑，不过效率还不错
func MsgDaemon() {
	for {
		select {
		case data := <-likemsgchan1:
			temp, _ := data.([]interface{})
			nickname, _ := temp[0].(string)
			target, _ := temp[1].(int)
			Type, _ := temp[2].(string)
			sendMsg(target, Type, nickname)
		case data := <-likemsgchan2:
			temp, _ := data.([]interface{})
			nickname, _ := temp[0].(string)
			target, _ := temp[1].(int)
			Type, _ := temp[2].(string)
			sendMsg(target, Type, nickname)
		case data := <-likemsgchan2:
			temp, _ := data.([]interface{})
			nickname, _ := temp[0].(string)
			target, _ := temp[1].(int)
			Type, _ := temp[2].(string)
			sendMsg(target, Type, nickname)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

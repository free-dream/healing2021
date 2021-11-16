package ws

import (
	"fmt"
	//"time"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
)

var (
	Conn *Connection
	//TestUid int = 1
)

func wsInit(w http.ResponseWriter, r *http.Request, wsConn *websocket.Conn, id string) bool {
	var err error
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return false
	}

	if Conn, err = initConnection(wsConn); err != nil {
		log.Println(err)
		Conn.Close()
		return false
	}

    Conn.uid = id
	Conn.storageAndRecovery()
	return true
}

func GetConn() *Connection {
	return Conn
}

func WsHandler(ctx *gin.Context) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
	)
	// get uid
	session := sessions.Default(ctx)
	id := session.Get("user_id").(int)
	uid := strconv.Itoa(id)
	// fake uid
	//uid := strconv.Itoa(TestUid)
	//TestUid++
	//fmt.Printf("uid:%v\n",uid)

	if isInit := wsInit(ctx.Writer, ctx.Request, wsConn, uid); isInit != true {
		return
	}
	conn := GetConn()
	conn.writeMessage([]byte("Hello, ws!"))

	for {
		if data, err = conn.readMessage(); err != nil {
			conn.Close()
		}
		conn.heartBeatCheck(data)
		conn.chatWatcher(data)
	}
}

type WsDataResp struct {
	Sys []respModel.Sysmsg `json:"sys"`
	Usr []respModel.Usrmsg `json:"usr"`
}

// 返回用户历史消息

func WsData(ctx *gin.Context) {
	// get uid
	session := sessions.Default(ctx)
	id := session.Get("user_id").(int)
	uid := uint(id)
	// fake uid
	//uid := uint(2)
	//AddFakeData()

	var resp WsDataResp
	var err error

	resp.Sys, err = dao.GetAllSysMsg(uid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "can not get data"})
		return
	}
	resp.Usr, err = dao.GetAllUsrMsg(uid)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "can not get data"})
		return
	}
	ctx.JSON(200, resp)
    conn := GetConn()
    load, ok := ConnMap.Load(uid)
    load2, ok2 := ConnMap.Load(conn.uid)
    fmt.Printf("load:%v  ok:%v\nload2:%v   ok:%\nuid:%v\nid:%v\n",load ,ok ,load2 ,ok2 , conn.uid, uid)
	return
}

func AddFakeData() {
	msg := respModel.SysMsg{}
	msg.Uid = 2
	msg.Type = 1
	msg.ContentId = 3
	conn := GetConn()
	conn.SendSystemMsg(msg)
}

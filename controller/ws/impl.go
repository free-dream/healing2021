package ws

import (
	//"fmt"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan []byte

	mutex    sync.Mutex
	isClosed bool
}

var (
	MsgBuffer = struct {
		sync.RWMutex
		m map[string]*DataBuffer
	}{m: make(map[string]*DataBuffer)}
	ConnMap sync.Map
	//MsgBuffer sync.Map
)

type DataBuffer struct {
	Type      string
	Sysmsg    respModel.SysMsg
	Usrmsg    respModel.UsrMsg
	FrontNode *DataBuffer
	NextNode  *DataBuffer
}

func uBufferInit(uid string) {
	if MsgBuffer.m[uid] == nil {
		MsgBuffer.Lock()
		head := DataBuffer{
			Type:      "head",
			FrontNode: nil,
			NextNode:  nil,
		}
		MsgBuffer.m[uid] = &head
		MsgBuffer.Unlock()
		return
	}
	queueSend(uid)
}

func appendNode(uid string, bufType string, msg interface{}) {
	MsgBuffer.Lock()
	// 头插法, buf.m[uid]存的是headnode
	if MsgBuffer.m[uid] == nil {
		MsgBuffer.Unlock()
		uBufferInit(uid)
		MsgBuffer.Lock()
	}
	bufCache := MsgBuffer.m[uid]
	newBuf := DataBuffer{}
	if bufType == "sys" {
		newBuf.Type = bufType
		newBuf.Sysmsg = msg.(respModel.SysMsg)
		if bufCache.NextNode != nil {
			newBuf.NextNode = bufCache.NextNode
			MsgBuffer.m[uid].NextNode.FrontNode = &newBuf
		}
		newBuf.FrontNode = bufCache
		MsgBuffer.m[uid].NextNode = &newBuf
	}
	if bufType == "usr" {
		newBuf.Type = bufType
		newBuf.Usrmsg = msg.(respModel.UsrMsg)
		if bufCache.NextNode != nil {
			newBuf.NextNode = bufCache.NextNode
			MsgBuffer.m[uid].NextNode.FrontNode = &newBuf
		}
		newBuf.FrontNode = bufCache
		MsgBuffer.m[uid].NextNode = &newBuf
	}
	MsgBuffer.Unlock()
}

func queueSend(uid string) {
	MsgBuffer.Lock()
	bufCache := MsgBuffer.m[uid]
	for bufCache.NextNode != nil {
		bufCache = bufCache.NextNode
	}
	for bufCache.Type != "head" {
		if bufCache.Type == "sys" {
			conn, _ := getNewConn(uid)
			msg, _ := json.Marshal(bufCache.Sysmsg)
			conn.writeMessage(msg)
		}
		if bufCache.Type == "usr" {
			conn, _ := getNewConn(uid)
			msg, _ := json.Marshal(bufCache.Usrmsg)
			conn.writeMessage(msg)
		}
		bufCache = bufCache.FrontNode
		bufCache.NextNode.FrontNode = nil
		bufCache.NextNode = nil
	}
	MsgBuffer.Unlock()
}

func initConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan []byte, 1),
	}

	go conn.readLoop()
	go conn.writeLoop()
	return
}

func (conn *Connection) storage(uid string) {
	ConnMap.Store(uid, conn)
	uBufferInit(uid)
	//fmt.Println(ConnMap)
	//fmt.Printf("buffer:%v\n",MsgBuffer)
}

func getNewConn(uid string) (*Connection, bool) {
	newConn, ok := ConnMap.Load(uid)
	if !ok {
		return &Connection{}, ok
	}
	return newConn.(*Connection), ok
}

func (conn *Connection) readMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	//fmt.Printf("read:%v\n",data)
	return
}

func (conn *Connection) writeMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	conn.wsConnect.Close()

	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			conn.Close()
			return
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			conn.Close()
			return
		}
	}
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			conn.Close()
			return
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			conn.Close()
			return
		}
	}
}

func (conn *Connection) heartBeatCheck(data []byte) {
	if reflect.DeepEqual(data, []byte("heartbeat")) {
		if err := conn.writeMessage(data); err != nil {
			conn.Close()
			return
		}
	}
}

func (conn *Connection) chatWatcher(data []byte, uid string) {
	newUsrMsg := respModel.UsrMsg{}
	if ok := json.Unmarshal(data, &newUsrMsg); ok == nil {
		if !filter(uint2str(newUsrMsg.FromUser), "^[1-9][0-9]*$") {
			conn.writeMessage([]byte("error format"))
			return
		}
		if !filter(uint2str(newUsrMsg.ToUser), "^[1-9][0-9]*$") {
			conn.writeMessage([]byte("error format"))
			return
		}
		if uint2str(newUsrMsg.FromUser) != uid {
			conn.writeMessage([]byte("error format"))
			return
		}
		if newUsrMsg.Message == "" {
			conn.writeMessage([]byte("error format"))
			return
		}

		newConn, isConn := getNewConn(uint2str(newUsrMsg.ToUser))
		if !isConn {
			if bErr := dao.UsrBackUp(newUsrMsg, 2); bErr != nil {
				//conn.writeMessage([]byte(bErr.Error()))
				conn.writeMessage([]byte("Fail to storage data"))
				return
			}
			appendNode(uint2str(newUsrMsg.ToUser), "usr", newUsrMsg)
			conn.writeMessage([]byte("ok"))
			return
		}
		if bErr := dao.UsrBackUp(newUsrMsg, 1); bErr != nil {
			conn.writeMessage([]byte("Fail to storage data"))
			return
		}
		conn.writeMessage([]byte("ok"))
		newConn.writeMessage(data)
	}
}

type SysMsg struct {
	Uid       uint      `json:"uid"`
	Type      int       `json:"type"`
	ContentId uint      `json:"contentId"`
	Song      string    `json:"song"`
	Time      time.Time `json:"time"`
}

type UsrMsg struct {
	FromUser uint   `json:"fromUser"`
	ToUser   uint   `json:"toUser"`
	Url      string `json:"user"` //录音url
	Song     string `json:"song"` //歌名
	Message  string `json:"message"`
}

func (conn *Connection) SendSystemMsg(sysMsg respModel.SysMsg) error {
	id := uint2str(sysMsg.Uid)
	newConn, ok := getNewConn(id)
	if !ok {
		if bErr := dao.SysBackUp(sysMsg, 2); bErr != nil {
			return bErr
		}
		appendNode(uint2str(sysMsg.Uid), "sys", sysMsg)
		return nil
	}
	if bErr := dao.SysBackUp(sysMsg, 1); bErr != nil {
		return bErr
	}
	msg, _ := json.Marshal(sysMsg)
	newConn.writeMessage(msg)
	return nil
}

func (conn *Connection) SendUsrMsg(usrMsg respModel.UsrMsg) error {
	id := uint2str(usrMsg.ToUser)
	newConn, ok := getNewConn(id)
	if !ok {
		if bErr := dao.UsrBackUp(usrMsg, 2); bErr != nil {
			return bErr
		}
		appendNode(uint2str(usrMsg.ToUser), "usr", usrMsg)
		return nil
	}
	if bErr := dao.UsrBackUp(usrMsg, 1); bErr != nil {
		return bErr
	}
	msg, _ := json.Marshal(usrMsg)
	newConn.writeMessage(msg)
	return nil
}

func uint2str(u uint) string {
	i := int(u)
	return strconv.Itoa(i)
}

func filter(param string, pattern string) bool {
	if ok, _ := regexp.Match(pattern, []byte(param)); !ok {
		return false
	}
	return true
}

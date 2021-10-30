### 食用说明:

```
import "healing2021/controller/ws"
//建议在controller完成
conn := ws.GetConn()
API:
// 用于发送系统消息和用户消息，详见产品文档
conn.SendSysMsg(respModel.SysMsg) error
conn.SendUsrMsg(respModel.UsrMsg) error


对外接口
//请求ws连接
GET("/ws")
// 获取用户历史消息
GET("/ws/history")

前端需要完成心跳检测，发送内容为"heartbeat"，后端会返回"heartbeat"
聊天信息需要按下面的格式编码成字符串
 {
	fromUser uint //自己的id
	toUser uint //对面的id
	message string //消息
}
发送信息成功会返回ok，不成功会给错误

完成连接的同时，后端会发送关机时间收到的消息，注意查收
```



### 基本流程：

初始化：

​			连接客户端，把连接存进hash表，用户信息缓存初始化

流程：

1. websocket.conn注册
2. impl.Conn注册
3. 开启读写线程
4. 提取uid（从业务），保存Conn连接，创建用户的buffer（如果已经有数据就全部发送，先进先出）
5. 读线程收到数据，先心跳过滤，然后聊天过滤，丢弃
6. 发现是心跳检测，原路返回同样的数据
7. 发现是聊天信息，根据来去的uid作转发
8. 如果对方连接未开，放buffer里
9. 业务需要发送系统信息和用户信息，封装好信息直接插入发送

要点：
1. 循环读取客户端信息，
2. 回复心跳检测，
3. 聊天信息监测，提取对方id从hash表拿连接发送聊天信息
4. 两个接口给内部业务，一个是系统信息发送，一个是用户信息发送
5. 所有已发送信息存进数据库，作为历史记录（把待发送也存进去，但是要注意发送的时候别冲突了）
6. 给客户端两个接口，一个是ws连接，一个是历史记录
7. ws连接成功，会立刻发送用户缓存消息队列的消息，并清空缓存

异常情况：
1. 对方用户关闭客户端，把消息缓存到对方用户的消息队列
2. 热更新，需要对消息队列备份（redis，**mysql**或者本地？）

### 内部组织：

1. 连接部分
  websocket包demo完成，连接后Conn作为全局变量，可提供给整个业务各个部分
2. Conn
  Conn的核心websocket基本结构体的指针，设计上还集成了消息进出的两个管道，一个控制关闭的管道，两个负责控制关闭管道的并发安全

```

type  Connection struct {
		wsConnect *websocket.Conn
		inChan chan []byte
		outChan chan []byte
		closeChan chan []byte
		mutex sync.Mutex
		isClosed bool
}
ConnMap sync.Map

主要函数:

// 连接初始化阶段由wshandler调用
initConnection(*websocket.Conn)(*Connection)

// 读写函数
(*Connection)readMessage()([]byte, error)
(*Connection)writeMessage([]byte) error

//两个线程完成连续读写操作，里面循环调用上面两个读写函数
(*Connection)readLoop()
(*Connection)writeLoop()

//心跳检测，收到"heartbeat"就回复"heartbeat"
(*Connection)heartBeatCheck([]byte)

//聊天监视，转发聊天
(*Connection)chatWatcher([]byte)

// 业务需要的发送系统信息与发送用户信息
(*Connection)sendSystemMsg(SysMsg)
(*Connection)sendUsrMsg(UsrMsg)

//用户连接存入hash表
(*Connection)storage(string)

//从hash表拿用户连接
getNewConn(string)
```

 3.  buffer

     map+双链表的形式储存，`map[uid]userbufferheadnode`，用户uid对应的是链表的head

```
MsgBuffer {
		sync.RWMutex
		m map[string]DataBuffer
}

type DataBuffer struct {
		Type string
		Sysmsg SysMsg
		Usrmsg UsrMsg
		NextNode *DataBuffer
}

// 用户buffer初始化，逻辑应该是create or update
uBufferInit(string)

//添加缓存, 用头插法方便函数里的局部变量逃逸
(buf MsgBuffer)appendNode(string, string, interface{})

//清除缓存，按道理，把指针指向改了，丢掉的变量应该会被自动清理
(MsgBuffer)deleteNode(DataBuffer, front)

// 批量发送, 每发送完就调用清除函数， 这个应该在init里执行
(MsgBuffer)queueSend(uid string)
```


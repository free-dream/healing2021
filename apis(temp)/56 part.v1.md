# <span id="">5. 消息推送</span>

## <span id="">5.1 拉取消息列表</span>

GET /message/list/{id}  HTTP1.1

id 指当前用户的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "target"：{    						//目标对象的个人信息 (系统消息的话填什么呢)
            "id": integer,
            "nickname": string,
			"avatar": text(url)				// 头像
        }, 
        "last":{    						// 最后一条的信息
            "id": integer,  				// 最后一条信息的id
            "from": integer,  				// 最后一条信息来自哪个用户
            "to": integer,					// 最后一条消息是发给哪个用户
            "content": string,				// 最后一条信息的内容
            "time": datatime,				// 最后一条信息的时间
            "type": integer,  				// 消息类型 1:用户消息, 2:系统消息
        },
        "number": 5   						// 未读消息数量
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "消息不存在!"}`

## <span id="">5.1 查看消息详情</span>

GET /message/detail/{id}  HTTP1.1

id 指要查看的消息的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "id":integer,  					// 信息的id
		"from":integer,  				// 发送该信息的用户id
        "user":{    					// 发送该信息的用户的信息
            "id": integer,
            "nickname": string,			// 用户昵称
			"avatar": text(url) 		// 用户头像
        }, 
        "content":text,					// 信息的内容
        "time":datatime,				// 信息的时间
        "type":integer,  				// 消息类型 0:普通消息 1:系统消息
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "消息不存在!"}`

## <span id="">5.2 获取用户消息的前两条信息内容</span>

GET /message/first/{id}  HTTP1.1

id 指要查看的消息的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "title" : string,			// 歌名
    "song" : string(url)		// 录音
}
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "消息不存在!"}`

## <span id="">5.3 发送信息</span>

POST/message/send  HTTP1.1

id 指要回复的消息的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "message_id" : integer,  		// 信息对应的消息id
	"from": integer,  				// 发送该信息的用户id
    "to": integer,					// 接收该信息的用户id
    "content": string,				// 内容
}
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "消息不存在!"}`

# <span id="">6. 广场主页</span>

## <span id="">5.1 拉取广场动态列表</span>

GET /dynamics/list/{method}  HTTP1.1

其中 method 可取： "new"/"recommend"/"search"

当选用的 method 为 "search" 时,在 url 后加上 ?keyword=xxx 即可拉取含有关键字的 状态、歌曲名 的动态列表

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "dynamics_id": integer,
        "content": string,								// 动态的内容
        "created_at": timestamp,
        "img" : [string(url)...],						// 多张图片的 url 地址，这里在展示时应该只用得上一张
        "song" : string,								// 要分享的歌名
        "lauds" : integer,								// 动态的点赞数
        "lauded": integer(0/1),							// 当前用户是否点赞该动态
        "comments" : integer,							// 动态的评论数
        "status" : ["status1", "status2" ...],			// 状态列表 元素都是string
        "creator": {
            "id": integer,	
            "nackname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 是否设置了头像（0代表没设置）
        }
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "动态不存在!"}`

## <span id="">5.2 发布动态</span>

POST  /dynamics/send  HTTP1.1

Content-Type: application/json

```js
{
    "content": string,								// 动态的内容
    "img" : [string(url)...],						// 上传的多张图片的url
    "song" : string,								// 要分享的歌名
    "status" : ["status1", "status2" ...]			// 状态列表 元素都是string
}
```

成功时：

HTTP/1.1 200 OK

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "动态发布失败"}`

## <span id="">5.3 查看动态的详情</span>

GET /dynamics/detail/{id}  HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "dynamics_id": integer,
    "content": string,								// 动态的内容
    "created_at": timestamp,
    "img" : [string(url)...],						// 多张图片的 url 地址
    "song" : string,								// 要分享的歌名
    "lauds" : integer,								// 动态的点赞数
    "lauded": integer(0/1),							// 当前用户是否点赞该动态
    "comments" : integer,							// 动态的评论数
    "status" : ["status1", "status2" ...],			// 状态列表 元素都是string
    "creator": {
        "id": integer,	
        "nackname" : string,						// 用户名
        "avatar": string(url),						// 头像
    	"avatar_visible": integer(0/1)				// 是否设置了头像（0代表没设置）
	}
}

```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## <span id="">5.4 给动态添加评论</span>

POST /dynamics/comment  HTTP1.1

id 为动态对应的 id

Content-Type: application/json

```json
{
    "dynamics_id": integer,		// 评论对应的 动态id
    "content": string
}
```

成功时：

HTTP/1.1 200 OK

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## <span id="">5.5 拉取动态的评论列表</span>

GET /dynamics/comment/{id}  HTTP1.1

id 为动态对应的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "comment_id": integer,
        "content": string,
        "creator": {
            "id": integer,	
            "nackname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 是否设置了头像（0代表没设置）
        },
        "created_at": timestamp,
        "lauds" : integer,								// 动态的点赞数
    	"lauded": integer(0/1)							// 当前用户是否点赞该动态
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## <span id="">5.6 给动态或评论点赞</span>（取消点赞）

PUT  /laud/{type}/{id}  HTTP1.1

其中 type 可取： "comment"/"dynamics"

id 为对应的 动态/评论 的 id

成功时：

HTTP/1.1 200 OK

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "点赞失败"}`

`{"message" : "取消点赞失败"}`




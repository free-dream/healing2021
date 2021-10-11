[toc]



# 1. 微信授权

## 1.1 授权

GET /auth/jump2[?redirect={encoded_uri}] HTTP/1.1

成功：

Content-Type application/json

200 OK
```json
{
  "nickname": "string",
}
```

需要先访问此接口接受redirect参数

失败:

`可能遇到401,用户未登录强制重定向进行授权登录`
## 1.2 jssdk的接口

GET /api/wxconfig HTTP/1.1

成功：

HTTP/1.1 200 OK

```json
{
    "appId": "string",
    "timestamp": timestamp,
    "nonceStr": "string",
    "signature": "string"
}
```

失败：

HTTP/1.1 500 Internal Server Error

Content-Type: application/json

```json
{"message" : "ip is not in whitelist"}
```

#  -----下面的接口都要带上前缀"/api"-----

# 2. 用户模块

## 2.1 用户注册

POST /user HTTP/1.1

成功：

Content-Type: application/json

200 OK
```json
{
  "nickname": "string",  
  "real_name": "string", //选填
  "phone_number": "string", //选填
  "sex": int,// 1:男 2:女 3:其他
  "school": "string" //可以传缩写过来 scut

}
```
失败：

Content-Type: application/json

403 Forbidden
```json
{
  "message": "昵称/手机号已存在,无法注册"
}
```
## 2.2 个人信息更新

PUT /user HTTP/1.1

成功：

Content-Type: application/json

200 OK

**下列全为可选项。**

```json
{
    "avatar": "string" //头像url,
    "nickname": "string",
    "avatar_visible": int,     	// 1：隐藏头像，0：不隐藏
    "phone_search": int,     	// 1：允许通过手机号查找，0：不允许
    "real_name_search": int,      	// 1：允许通过姓名查找，0：不允许
    "signature": "string"  		//个性签名（可不填）
}
```
失败:

403 Forbidden

Content-Type: application/json

```json
{"message" : "修改失败"}
```

## 2.3 获取自己信息(用户个人页信息拉取)

GET /user HTTP1.1

成功：

200 OK

Content-Type: application/json

```json
{
  "avatar": "string",
  "nickname": "string",
  "school": "string",
  "signature": "string",
  "mySelections": {
    "model": "string" //模块名 治愈或是投递箱
    "song_name": "string",
    "post_time": "string", //"yyyy-mm-dd
    "anonymous": int, //1:匿名 2:不匿名
    "healingId": int, //所点歌对应的治愈模块id
  },
  "mySongs": {
    "model": "string" //模块名 治愈或是投递箱
    "post_time": "string",
    "song_name": "string",
  },
  "myLikes": {
    "model": "string",
    "post_time": "string",
    "song_name": "string",
    "likeId": int, //对应点赞的id
    "likeNum": int //对应点赞数
  },
  "moments": {
    "post_time": "string",
    "state": "string", //状态:摸鱼
    "content": "string", //动态内容
    "momentId": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likeNum": int 
  }
}
```

失败：

403 Forbidden

Content-Type: application/json

```json
{"message" : "获取头像失败"}
```

## 2.4 更新个人背景
POST /background HTTP/1.1

成功：

200 OK

Content-Type: application/json

```json
{
  "background": "string" //背景图片url
}
```
失败：

500 Internal Server Error
```json
{"message" : "更新头像失败"}
```
## 2.5 治愈系详情页
GET /healingPage HTTP/1.1

Query
```json
{
  "healingId": int
}
```
成功：

Content-Type: application/json

```json
{
  "songName": "string",
  "songId": int,
  "selector": {
    "name": "string",//点歌用户名
    "style": "string",//风格
    "post_time": "string", //"yyyy-mm-dd"
    "remark": "string" //30字以内
  },
  "singers": {
    "singer": "string",
    "songId": int,
    "likeId":int,
    "song": "string" //歌曲url
  }
}
```
失败:

```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```




## 2.6 获取他人信息
GET /callee HTTP/1.1

Query

```json
{
  "calleeId": int //被访问用户的id
}
```

成功：

200 OK

Content-Type: application/json

```json
{
  "avatar": "string",
  "nickname": "string",
  "school": "string",
  "signature": "string",
  "mySelections": {
    "model": "string" //模块名 治愈或是投递箱
    "song_name": "string",
    "post_time": "string", //"yyyy-mm-dd
  },
  "mySongs": {
    "model": "string" //模块名 治愈或是投递箱
    "post_time": "string",
    "song_name": "string",
    "likeNum": int,
    "songId": int, //受访者所唱歌曲的id
  },
  "myLikes": {
    "model": "string",
    "post_time": "string",
    "song_name": "string",
    "likeId": int, //对应点赞的id
    "likeNum": int //对应点赞数
  },
  "moments": {
    "post_time": "string",
    "state": "string", //状态:摸鱼
    "content": "string", //动态内容
    "momentId": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likeNum": int 
  }
}
```

失败:

```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```

## 2.7 获取二维码
GET /QR_code HTTP/1.1

Query
```json
{
  "userId": int,//发起请求的用户id
  "songId": int,//对应歌曲id
}
```
成功：

```json
{
  "QR_code": "string",//二维码url
}
```
失败:

```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```

# 3 经典治愈

所有初始化调用的接口最好直接将部分数据缓存在redis里，若有必要的更新将数据写入mysql

全部接口在有cookie的情况下进行

## 3.1 治愈页面关联接口

### 3.1.1 轮播图接口

***轮播图接口更新可能依赖于多于一张表***

包括可能存在的广告商(笑)

GET /healing/bulletin HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "picture":text(url),	//广告或曲目对应的图片
        "address":text(url)		//对应的链接，广告外链或者翻唱界面
    }
	...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.2 点歌(唱歌)请求获取

GET /healing/selections/list HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
	{
		"nickname":string,	//可匿名
        "remark":text,	//备注
        "song_name":string,	
        "language":string,	//以下两项为有限选项
        "style":string,
        "user_id":integer,	//点歌用户的id
        "post_time":string(datetime),	//时间，排序用
        "model":string,	//所属模块名，有限选项，索引用
	}
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.3 翻唱列表获取

GET /healing/covers/list HTTP 1.1

***根据点赞表和翻唱时间综合排序***

***初始化就从mysql检索，准备好两幅表缓存在redis里***

***设置一个更新器，若表格发生了更新，先写入redis,每隔一段时间将数据录入mysql一次***

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "nickname":string,	//翻唱者，允许匿名
        "post_time":string(datetime),	//翻唱时间
        "avatar":text(url),	//用于加载的用户头像url
        "selection_url":text(url)	//用于跳转的治愈详情页面url
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.4 索引和不同的排序

GET /healing/{a}/{model}/{rankby} HTTP 1.1

***a指selections和covers其中之一***

***label指所有给定的model,all也是model之一***

***rankby指所有给定的排序方法,default即默认检索序也是rankby之一***

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//if a==covers
    {
        "nickname":string,	//翻唱者，允许匿名
        "post_time":string(datetime),	//翻唱时间
        "avatar":text(url),	//用于加载的用户头像url
        "selection_url":text(url)	//用于跳转的用户个人页面url
    }
    ...
]
```



```json
[//if a==selections
	{
		"nickname":string,	//可匿名
        "remark":text,	//备注
        "song_name":string,	
        "language":string,	//以下两项为有限选项
        "style":string,
        "user_id":integer,	//点歌用户的id
        "post_time":string(datetime),	//时间，排序用
        "model":string,	//所属模块名，有限选项，索引用
	}
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`



### 3.1.5 听歌点赞

POST /healing/covers/like HTTP 1.1

***注意到点赞的话对应的排行会更新，此处应该对缓存数据进行更新并决定是否持久化***

Content-Type: application/json

```json
{
    "covers_id":int
}
```

成功:

HTTP/1.1 200 OK

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "点赞失败惹~"}`

## 3.2 功能性接口(包括抽奖、排行榜、热榜)

### 3.2.1 抽奖

//前提是持有包含用户信息的cookie

#### 3.2.1.1 奖池信息获取

GET /healing/lotterybox/lotteries HTTP 1.1

Content-Type: application/json

```json
[
    {
        "picture":string(url),	//奖品图像
        "name":string,	//奖品名
        "possibility":double	//概率，百分号由前端补充
    }
    ...
]
```

成功:

HTTP/1.1 200 OK

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "获取奖池信息错误"}`

#### 3.2.1.2 抽奖

POST /healing/lotterybox/draw HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "check":int,	//1或0或2，对应中或者没中或者积分不足一抽，前端设置对应文案
    "name":string,	//奖品名,default=null
    "picture":string(url)	//奖品图片，或者没中的话指向一些默认图片url
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "抽奖失败"}`

#### 3.2.1.3 拉取用户中奖记录

GET /healing/lotterybox/prizes

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//已按照时间戳进行排序
    {
        "name":string,	//奖品名
        "picture":string	//奖品图片url
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "拉取中奖信息失败"}`

#### 3.2.1.4 拉取对应用户的任务列表

GET /healing/lotterybox/tasktable

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "check": integer,	//0未完成，1已完成	
        "task": {
            "id": integer,	
            "text" : string,	// 任务描述					
            "target": integer	//目标次数				
        },
        "counter":integer	//已经进行的次数，交付前端表示进度
    },
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "加载任务列表失败"}`

### 3.2.2 排行榜

GET /healing/rank/{school}

***school指学校名称，中间有一次对换，例如华工==华南理工大学==scut***

***特别地，全部==all也视为学校名***



成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10
    {
        "avatar":string,	//用户头像url
        "nickname":string,	//用户名
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "加载排行榜失败"}`



### 3.2.3 每日热榜

POST /healing/dailyrank/{date}

***date指日期，default为当天日期,在索引前有一个转换***

点赞调用 /healing/covers/like 接口

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10
    {
        "avatar":string,	//用户头像url
        "nickname":string,	//用户名
        "post_time":string(datetime),	//时间
        "likes":int,	//点赞数
        "song_name":string	//歌曲名
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "加载排行榜失败"}`

### 3.2.4 搜索页面的相关接口

1. 搜索历史是否可以交付前端缓存？//后端先拉了一个表项
2. 搜索要求不同关键字之间以空格分开
3. 热榜建议缓存到redis里，调用的时候直接抓取若干项 //已经有了

#### 3.2.4.1 搜索接口

POST /healing/search

Content-Type: application/json

//此处前端或后端应对搜索记录和对应的页面url作一保存

```json
{
    "keyword":string,	//以空格或tab分开，便于检索
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//返回一个包含三个长列表和一个含一个json的短列表,按返回顺序:
    [{//依次是以下三个列表的长度
        "user":int,
        "selections":int,
        "covers":int
    }],
    [//用户表
        {
            "user_id":int,	//用于组合跳转个人页面，可改为个人页面的url
            "avatar":string(url),	//用于索引头像
            "nickname":string,	//用户名
            "slogan":string	//个性签名
        }
        ...
    ],
    [//点歌(唱歌)表
        {
            "selection_id":int,	//用于跳转治愈详情页，可换成url
            "nickname":string,	//点歌用户名
            "song_name":string,	//歌曲名
            "post_time":string(datetime),	//点歌时间
            "avatar":string(url)	//用户头像索引
        }
        ...
    ],
    [//听歌表
        {
            "cover_id":int,
            "nickname":string,
            "song_name":string,
            "post_time":string(datetime),
            "avatar":string(url)	
        }
    ]
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "搜索失败"}`

#### 3.2.4.2 搜索历史 (保留，可选，视前端需求)

GET /healing/search/history

***建议搜索历史缓存在redis里,用户退出时持久化于mysql***

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "keyword":string,	//搜索关键字
        "result":string	//搜索结果
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "搜索历史获取失败"}`

#### 3.4.2.3 热榜

调用 GET /healing/dailyrank/{default}

前端取用需求的数据

## 3.3 点歌页接口

POST /healing/selection

//上传一条点歌需求

Content-Type: application/json

```json
{
    {
        "remark":text,	//备注
        "song_name":string,	
        "language":string,	
        "style":string,
        "model":string,	//所属模块名，有限选项，索引用
	}
}
```

成功:

HTTP/1.1 200 OK

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "搜索历史获取失败"}`



## 3.4 录音接口

POST /healing/recording

**//录音接口按去年应该另有途径，这里姑且使用http表单上传**

Content-Type: multipart/form-data

```multipart/form-data
......
Content-Type: multipart/form-data; boundary=---------------------------58731222010402	
//http表单会自己生成boundary分割不同数据，此处随便拿了一个
......
-----------------------------1100116873645
Content-Disposition: form-data; name="public" 	//0为私有，1为公开

0
-----------------------------1100116873645
Content-Disposition: form-data; name="childhood"	//0为经典，1为童年
 
1
-----------------------------1100116873645
Content-Disposition: form-data; name="视录音名称,由前端决定"	
Content-Type:audio/mp3
//传递录音文件,目前不知道是什么格式，姑且视为mp3
 
-----------------------------1100116873645
```

成功:

HTTP/1.1 200 OK

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "上传录音失败"}`

# 4.追忆童年

## 4.1 追忆童年主页相关接口

### 4.1.1 推荐歌曲，根据click数降序获取10项(大家都在听)

GET /childhood/rank HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10，若出现等click数的情况按name字母序
	{
    	"name":string,	//歌曲名
    	"icon":text(url),	//歌曲图标,此处是大图标
        "click":int	//这里可以做一个小小的换算，比如1 click = 500 热度
	}
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "列表不存在"}`

### 4.1.2 获取歌曲列表

GET /childhood/list HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表name字母序
    {
        "name":string,
        "avatar":text(url), //小图标
        "discription":text,	//简单描述
        "work_name":string	//音乐出处
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "列表不存在"}

## 4.2 原翻唱页相关接口

### 4.2.1 获取原唱相关信息

GET /childhood/original/{name}/info HTTP 1.1

***name指对应歌曲的名字***

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "song_id":integer,	//歌曲id
    "name": string,
    "singer": string,
    "icon":text(url)   //歌曲图标
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "不存在对应的歌曲"}`

### 4.2.2  获取用户翻唱列表并排序

GET /childhood/original/{name}/cover HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {	//此处应该是不需要用户id的，对应的记录索引后应位于歌曲id下
        "nickname": string,
        "avatar": text(url),   //用户头像
        "post_time": datetime,	//可以按时间顺序排序，不过这里最好还是按喜爱数排序
        //"like": integer(uint) ,可选，设计上没有提到
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "不存在对应的歌曲"}`

### 4.2.3 录音接口

调用 POST /healing/recording 接口，有一个属性用于确认归属

## 4.3 歌曲页相关接口

### 4.3.1 点赞接口

调用 POST /healing/covers/like 接口

### 4.3.2  加载歌曲(翻唱)

POST /healing/player

//不知道这里有没有播放列表的设计

//这里关于童年的设计挺奇怪的，不过我自己先填上了

```json
{
    "jump":integer,	//0为上一首,1为下一首,2为跳转，跳转则传回对应歌曲的id
    "check":integer,	//0为经典治愈，1为童年
    "cover_id":integer	//若jump=2，则传回对应的翻唱id
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{//童年模式下的返回值
    "song_name":string,	//歌曲名
    "file":string,	//翻唱文件url
    "lyrics":string,	//如果有的话，默认是null
    "icon":string,	//图标url
    "work_name":string,	//作品名
    "nickname":string	//翻唱者
}
```



```json
{//常规模式下的返回值
    "song_name":string,	//歌曲名
    "file":string,	//翻唱文件url
    "avatar":string,	//头像大图url
    "nickname":string	//翻唱者
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "不存在对应的歌曲"}`

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

## <span id="">6.1 拉取广场动态列表</span>

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

## <span id="">6.2 发布动态</span>

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

## <span id="">6.3 查看动态的详情</span>

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

## <span id="">6.4 给动态添加评论</span>

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

## <span id="">6.5 拉取动态的评论列表</span>

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

## <span id="">6.6 给动态或评论点赞</span>（取消点赞）

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




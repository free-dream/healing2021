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
        "thumbnail":text(url), //小图标
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

调用 POST /healing/record 接口，有一个属性用于确认归属

这个接口应该做的事：

1. 更新covers(翻唱表)
2. 更新songs(歌曲表)

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
    "icon":string,	//大图标url
    "thumbnail":string,	//小图标url
    "work_name":string,	//作品名
    "nickname":string	//翻唱者
}
```



```json
{//常规模式下的返回值
    "song_name":string,	//歌曲名
    "file":string,	//翻唱文件url
    "avatar":string,	//头像大图url
    "User_thumbnail":string,	//头像小图url
    "nickname":string	//翻唱者
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "不存在对应的歌曲"}`

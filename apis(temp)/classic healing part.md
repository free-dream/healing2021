# 3. 经典治愈

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
        "user_thumbnail":text(url),	//用于加载的用户小图url
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
        "user_thumbnail":text(url),	//用于加载的用户小图url
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
        "thumbnail":string,	//用户头像url
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
        "thumbnail":string,	//用户头像url
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
            "thumbnail":string(url),	//用于索引头像
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
            "thumbnail":string(url)	//用户头像索引
        }
        ...
    ],
    [//听歌表
        {
            "cover_id":int,
            "nickname":string,
            "song_name":string,
            "post_time":string(datetime),
            "thumbnail":string(url)	
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


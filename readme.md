* [1\. 微信授权](#1-微信授权)
  * [1\.1 授权](#11-授权)
  * [1\.2 jssdk的接口](#12-jssdk的接口)
* [\-\-\-\-\-下面的接口都要带上前缀"/api"\-\-\-\-\-](#-----下面的接口都要带上前缀api-----)
* [2\. 用户模块](#2-用户模块)
  * [2\.1 用户注册](#21-用户注册)
  * [2\.2 个人信息更新](#22-个人信息更新)
  * [2\.3 获取自己信息(我的点歌与个人信息)](#23-获取自己信息我的点歌与个人信息)
  * [2\.4 更新个人背景](#24-更新个人背景)
  * [2\.5 治愈系详情页](#25-治愈系详情页)
  * [2\.6 获取他人信息](#26-获取他人信息)
* [3 经典治愈](#3-经典治愈)
  * [3\.1 治愈页面关联接口](#31-治愈页面关联接口)
    * [3\.1\.1 轮播图接口](#311-轮播图接口)
    * [3\.1\.2 点歌(唱歌)请求获取](#312-点歌唱歌请求获取)
    * [3\.1\.3 翻唱列表获取](#313-翻唱列表获取)
    * [3\.1\.4 治愈系翻唱接口](#314-治愈系翻唱接口)
    * [3\.1\.5 听歌点赞](#315-听歌点赞)
  * [3\.2 功能性接口(包括抽奖、排行榜、热榜)](#32-功能性接口包括抽奖排行榜热榜)
    * [3\.2\.1 抽奖](#321-抽奖)
      * [3\.2\.1\.1 奖池信息获取](#3211-奖池信息获取)
      * [3\.2\.1\.2 抽奖](#3212-抽奖)
      * [3\.2\.1\.3 拉取用户中奖记录](#3213-拉取用户中奖记录)
      * [3\.2\.1\.4 拉取对应用户的任务列表](#3214-拉取对应用户的任务列表)
      * [3\.2\.1\.5 任务更新/领取积分](#3215-任务更新领取积分)
    * [3\.2\.2 排行榜](#322-排行榜)
    * [3\.2\.3 每日热榜](#323-每日热榜)
    * [3\.2\.4 搜索页面的相关接口](#324-搜索页面的相关接口)
      * [3\.2\.4\.1 搜索接口](#3241-搜索接口)
      * [3\.2\.4\.2 搜索历史 (保留，可选，视前端需求)](#3242-搜索历史-保留可选视前端需求)
      * [3\.4\.2\.3 热榜](#3423-热榜)
  * [3\.3 点歌页接口](#33-点歌页接口)
* [4\.追忆童年](#4追忆童年)
  * [4\.1 追忆童年主页相关接口](#41-追忆童年主页相关接口)
    * [4\.1\.1 推荐歌曲，根据click数降序获取10项(大家都在听)](#411-推荐歌曲根据click数降序获取10项大家都在听)
    * [4\.1\.2 获取歌曲列表](#412-获取歌曲列表)
  * [4\.2 原翻唱页相关接口](#42-原翻唱页相关接口)
    * [4\.2\.1 获取原唱相关信息](#421-获取原唱相关信息)
    * [4\.2\.2  获取用户翻唱列表并排序](#422--获取用户翻唱列表并排序)
    * [4\.2\.3 录音接口](#423-录音接口)
  * [4\.3 歌曲页相关接口](#43-歌曲页相关接口)
    * [4\.3\.1 点赞接口](#431-点赞接口)
    * [4\.3\.2  加载歌曲(翻唱)](#432--加载歌曲翻唱)
* [<span id="">5\. 消息推送</span>](#5-消息推送)
  * [<span id="">5\.1 拉取消息列表</span>](#51-拉取消息列表)
  * [<span id="">5\.1 查看消息详情</span>](#51-查看消息详情)
  * [<span id="">5\.2 获取用户消息的前两条信息内容</span>](#52-获取用户消息的前两条信息内容)
  * [<span id="">5\.3 发送信息</span>](#53-发送信息)
* [<span id="">6\. 广场主页</span>](#6-广场主页)
  * [<span id="">6\.1 拉取广场动态列表</span>](#61-拉取广场动态列表)
  * [<span id="">6\.2 发布动态</span>](#62-发布动态)
  * [<span id="">6\.3 查看动态的详情</span>](#63-查看动态的详情)
  * [<span id="">6\.4 给动态添加评论</span>](#64-给动态添加评论)
  * [<span id="">6\.5 拉取动态的评论列表</span>](#65-拉取动态的评论列表)
  * [<span id="">6\.6 给动态或评论点赞</span>（取消点赞）](#66-给动态或评论点赞取消点赞)



# 1. 微信授权

## 1.1 授权

GET /wx/jump/?redirect=  HTTP/1.1

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

#  -----下面的接口都要带上前缀"/api"-----

# 2. 用户模块

## 2.1 用户注册

POST /user HTTP/1.1

成功：

Content-Type: application/json

```json
{
"nickname": "string",  
"real_name": "string", //选填
"phone_number": "string", //选填
"sex": int,// 1:男 2:女 3:其他
"school": "string" //可以传缩写过来 scut
"hobby":[]string

}
```

200 OK

```json
{
  "user_id": int
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
    "phone_search": int,     	// 0：允许通过手机号查找，1：不允许
    "real_name_search": int,      	// 0：允许通过姓名查找，1：不允许
    "signature": "string"  		//个性签名（可不填）
}
```
失败:
```json
{
  "avatar": string
}
```
403 Forbidden

Content-Type: application/json

```json
{"message" : "修改失败"}
```

## 2.3 获取自己信息(我的点歌与个人信息)

GET /user HTTP1.1

成功：

200 OK

Content-Type: application/json

//index项不止一条，index从0开始

```json
{
   "message":{
  "avatar": "string",
  "nickname": "string",
  "school": "string",
  "signature": "string",
},
  "mySelections": {
      index:{
    "song_name": "string",
    "created_at": "string", //“2006-01-02 15:04:05”
    "anonymous": int, //1:匿名 2:不匿名
      }
  },
  "mySongs": {
      index:{

    "created_at": "string",
    "song_name": "string",
      }
  },
  "myLikes": {
      index:{
    "created_at": "string",
    "song_name": "string",
    "id": int, //对应点赞的id
    "likeNum": int //对应点赞数
      }
  },
  "moments": {
      index:{
    "created_at": "string",
    "state": "string", //状态:摸鱼
    "content": "string", //动态内容
    "id": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likeNum": int 
      }
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
  "selectionId": int
}
```
成功：

Content-Type: application/json

```json
{
    user：{
  "songName": "string",
  "selectionId": int,
  "name": "string",//点歌用户名
  "style": "string",//风格
  "created_at": "string", //“2006-01-02 15:04:05”
  "remark": "string" //30字以内
}
  "singers": {
      index{
    "singer": "string",
    "songId": int,
    "likeId":int,
    "song": "string" //歌曲url
  }
  }
}
```
失败:

```json
{
  "statusCode": 401,
  "message": "param error"
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
    message：{
  "avatar": "string",
  "nickname": "string",
  "school": "string",
  "signature": "string",
}
  "mySelections": {
      index:{
    "model": "string", //模块名 治愈或是投递箱
    "song_name": "string",
    "created_at": "string", //“2006-01-02 15:04:05”
      },
  },
  "mySongs": {
      index：{
    "model": "string",//模块名 治愈或是投递箱
    "created_at": "string",
    "song_name": "string",
    "likeNum": int,
    "songId": int, //受访者所唱歌曲的id
  },
  },
  "myLikes": {
      index:{
    "model": "string",
    "created_at": "string",
    "song_name": "string",
    "likeId": int, //对应点赞的id
    "likeNum": int //对应点赞数
      }
  },
  "moments": {
      index:{
    "created_at": "string",
    "state": "[]string", //状态:摸鱼
    "content": "string", //动态内容
    "momentId": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likeNum": int 
      }
  }
}
```

失败:

```json
{
  "statusCode": 401,
  "message": "param error"
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
        "address":text(url),	//对应的链接，广告外链或者翻唱界面
        "weight":int,//权重
    }
	
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.2 点歌(唱歌)请求获取

GET /healing/selections/list HTTP 1.1

```
{
label:string //recommend all 或对应风格，语言
"rankWay":int //1综合排序，2最新binding:"required"`
"page":int//页数,当获取个数小于10时，页数不再加一
//用户第一次到页面为1，后每次上拉逐次加一
}
```

成功:



Content-Type: application/json

```json
[
	{
		"nickname":string,	//可匿名
        "id":int,//对应点歌id
        "song_name":string,
        "user_id":integer,	//点歌用户的id
        "created_at":string(datetime),	//“2006-01-02 15:04:05”
        "avatar":string
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

GET /healing/selections/list HTTP 1.1

```
{
label：string //recommend all 或对应风格，语言
"rankWay":int //1综合排序，2最新binding:"required"`
"page":int //页数
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
       "nickname":string,	//可匿名
        "id":int,//对应翻唱歌id
        "song_name":string,
        "user_id":integer,	//翻唱用户的id
        "created_at":string(datetime),	//“2006-01-02 15:04:05”
        "avatar":string,
        "file":string//歌曲url
    }
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.4 治愈系翻唱接口

POST /healing/cover

```json
{
    "selection_id":int,//点歌id
    "record":[]string,//拼接的录音url
    module:int 1表示治愈系翻唱
}
```

成功：

```json
{		
    "nickname":string,	//可匿名
    "id":int,//对应翻唱歌id
    "song_name":string,
    "user_id":integer,	//翻唱用户的id
    "created_at":string(datetime),	//“2006-01-02 15:04:05”
    "avatar":string,
    "file":string//歌曲url
}
```



200 "OK"

失败：

HTTP/1.1 401 Forbidden

Content-Type: application/json

`{"message" : "error param"}`



### 3.1.5 听歌点赞/取消点赞

GET /like HTTP 1.1

统一使用通用模块的点赞操作，详见接口8.1

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

GET /healing/lotterybox/draw HTTP 1.1

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
        "picture":string,	//奖品图片url
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

#### 3.2.1.5 任务更新/领取积分

### 3.2.2 排行榜

#### 3.2.2.1 学校积分排名

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

#### 3.2.2.2 用户当前排名

GET /healing/rank/user

**200人以内有详细排名，接下来依次是:**

**200--500**

**500--1000**

**>1000**

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "rank":string
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "加载用户排名失败"}`

### 3.2.3 每日热榜

#### 3.2.3.1 获取按日的热榜

Get /healing/dailyrank/{date}
***date的日期遵循统一格式 mm-dd*** 
***日期不早于上线当日***

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

#### 3.2.3.2 总体热榜

Get /healing/dailyrank/{date}

**全部时间获赞数最多的翻唱**

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

```json
 {
       	"remark":text,	//备注
        "song_name":string,	
        "language":string,	
        "style":string,
}
//童年歌曲点歌的格式
//"style":童年
//"language":中文
```



Content-Type: application/json

```json
	{
        "nickname": string,
        "id": int,//点歌的id
        "song_name": string,
        "user_id": int,
        "created_at": "2021-10-26T16:30:27+08:00",
        "avatar":string
    } 
```

成功:

HTTP/1.1 200 OK

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "搜索历史获取失败"}`



# 4.追忆童年

## 4.1 追忆童年主页相关接口

### 4.1.1 推荐歌曲，根据click数降序获取10项(大家都在听)

GET /childhood/rank HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10，click数降序
	{
        "classic_id":int, //用于跳转到对应的原翻唱页
    	"name":string,	//歌曲名
    	"icon":text(url),	//歌曲图标
        "click":int	    //听歌人数
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
[
    {
       	"classic_id":int, 	//用于跳转到对应的原翻唱页
    	"name":string,		//歌曲名
    	"icon":text(url),	//歌曲图标
        "work_name":string	//作品名
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

GET /childhood/original/info HTTP 1.1

```json
{
	"classic_id":int
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "classic_id":int, 	//用于播放原唱
    "song_name": string,
    "singer": string,
    "icon":text(url)   //歌曲图标
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "不存在对应的歌曲"}`

### 4.2.2  获取用户翻唱列表并排序

GET /childhood/original/covers HTTP 1.1

```json
{
	"classic_id":int
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "cover_id":int,			// 用于进入歌曲页
        "nickname": string,		// 用户昵称
        "avatar": text(url),   	// 用户头像
        "post_time": datetime,	// “2006-01-02 15:04:05”
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

### 4.3.2 当前歌曲的信息获取

GET /healing/covers/player

```
{
	"cover_id":int
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
	"cover_id":int,
	"file":  url,		//歌曲录音的
	"name":string,		//歌曲名
	"nickname":string,	//翻唱者
    "icon":text(url),	//歌曲图标
    "work_name":string	//作品名
}
```

### 4.3.2  歌曲跳转(翻唱)

POST /healing/covers/jump

```json
{
    "jump":integer,		//0为上一首,1为下一首
    "check":integer,	//0为经典治愈，1为童年
    "cover_id":integer	//若jump=2，则传回对应的翻唱id
    
    // 分享，直接拿着"classic_id"参数去发布动态的接口即可
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{//童年模式下的返回值
    "cover_id":int,
	"file":  url,		//翻唱文件url
	"name":string,		//歌曲名
    "icon":text(url),	//歌曲图标
    "nickname":string,	//翻唱者
    "work_name":string	//作品名
}
```

```json
{//常规模式下的返回值
    "cover_id":int,
    "file":string,			//翻唱文件url
    "name":string,			//歌曲名
    "icon":string,			//头像大图url
    "nickname":string,		//翻唱者
    "work_name":string(空串) //作品名
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
        "time":datatime,				// 信息的时间“2006-01-02 15:04:05”
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

为实现分页功能（10条动态为一页），url 后**必须**加上 ?page=xx 即可取得第 xx 页的动态列表（page=0为第一页）

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    {
        "dynamics_id": integer,
        "content": string,								// 动态的内容
        "created_at": "2006-01-02 15:04:05",
        
        "song" : string,								// 要点的歌名
        "module" : int,									// 1为经典，2为童年
        "lauds" : integer,								// 动态的点赞数
        "lauded": integer(0/1),							// 当前用户是否点赞
        "comments" : integer,							// 动态的评论数
        "status" : ["status1", "status2" ...],			// 状态列表

        "creator": {
            "id": integer,	
            "nackname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 0代表没设置头像
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
    "content": string,							// 动态的内容
    "status": ["status1", "status2"...],		// 状态列表元素都是string
        
    "have_selection":int,						// 点歌了为1，否则为0
    "is_normal":int,							// 经典点歌为0，童年为1
   
    // 经典点歌填这四个
    "song_name":string,	
    "language":string,	
    "style":string,
    "remark":text,								//点歌的备注
     
     // 童年点歌填这一个
    "classic_id":int
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
    "created_at": "2006-01-02 15:04:05",
    
    "song" : string,								// 歌名
    "selection_id" : string, 						// 点歌id
    "module" : int,									// 1为经典，2为童年
    "lauds" : integer,								// 动态的点赞数
    "lauded": integer(0/1),							// 当前用户是否点赞该动态
    "comments" : integer,							// 动态的评论数
    "status" : ["status1", "status2" ...],			// 状态列表

    "creator": {
        "id": integer,	
        "nackname" : string,						// 用户名
        "avatar": string(url),						// 头像
    	"avatar_visible": integer(0/1)				// 0代表没设置头像
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
        "created_at": "2006-01-02 15:04:05",
        "lauds" : integer,								// 动态的点赞数
    	"lauded": integer(0/1),							// 当前用户是否点赞
    
        "creator": {
            "id": integer,	
            "nackname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 是0代表没设置头像
        }
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## 6.6 动态、评论点赞

使用 8.1 通用点赞接口进行点赞

## 6.7 动态热门搜索（最多十条）

GET /dynamics/hot HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    string1, string2,...
]
```

失败时(例子)：

HTTP/1.1 500 Forbidden

Content-Type: application/json

```
{"message" : "服务端出错"}
```

## 6.8 大家的状态（最多十八条）

GET /dynamics/states HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    state1, state2,...
]
```

失败时(例子)：

HTTP/1.1 500 Forbidden

Content-Type: application/json

```
{"message" : "服务端出错"}
```

## 6.9 点歌页歌曲推荐(最多30条)

GET /dynamics/songt HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    {
    	"song_name":string,	
    	"language":string,	
    	"style":string,
	}
	...
}
```

失败时(例子)：

HTTP/1.1 500 Forbidden

Content-Type: application/json

```
{"message" : "服务端出错"}
```

# 管理员相关

## 7.1 动态、评论删除

仅管理员可用，现不妨将 userId 为 0-5 的账号预留，当管理员账号

POST /administrators/content HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "type" : int,	//1代表动态，2代表评论
    "id" : int,		//对应动态、评论的Id
}
```

失败时(例子)：

HTTP/1.1 500 Forbidden

Content-Type: application/json

```
{"message" : "服务端出错"}
```

# 通用功能

## 8.1点赞

POST /praise HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "todo" : int, 	//1代表点赞，-1代表取消点赞
    "type" : int, 	//1代表动态，2代表评论,3代表翻唱
    "id" : int 		//对应动态、评论、翻唱的Id
}
```

失败时(例子)：

HTTP/1.1 500 Forbidden

Content-Type: application/json

```
{"message" : "服务端出错"}
```

## 

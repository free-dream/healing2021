* [1\. 微信授权](#1-微信授权)
  * [1\.1 授权](#11-授权)
  * [1\.2 假登录](#12-假登录)
* [\-\-\-\-\-下面的接口都要带上前缀"/api"\-\-\-\-\-](#-----下面的接口都要带上前缀api-----)
* [2\. 用户模块](#2-用户模块)
  * [2\.1 用户注册](#21-用户注册)
    * [2\.1\.1 用户爱好选择](#211-用户爱好选择)
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
    * [3\.1\.5 听歌点赞/取消点赞](#315-听歌点赞取消点赞)
  * [3\.2 功能性接口(包括抽奖、排行榜、热榜)](#32-功能性接口包括抽奖排行榜热榜)
    * [3\.2\.1 抽奖](#321-抽奖)
      * [3\.2\.1\.1 奖池信息获取](#3211-奖池信息获取)
      * [3\.2\.1\.2 抽奖](#3212-抽奖)
      * [3\.2\.1\.4 拉取对应用户的任务列表](#3214-拉取对应用户的任务列表)
      * [3\.2\.1\.5 获取当前用户的积分](#3215-获取当前用户的积分)
    * [3\.2\.2 排行榜](#322-排行榜)
      * [3\.2\.2\.1 学校积分排名](#3221-学校积分排名)
      * [3\.2\.2\.2 用户当前排名](#3222-用户当前排名)
    * [3\.2\.3 每日热榜](#323-每日热榜)
      * [3\.2\.3\.1 获取按日的热榜](#3231-获取按日的热榜)
      * [3\.2\.3\.2 总体热榜](#3232-总体热榜)
    * [3\.2\.4 搜索页面的相关接口](#324-搜索页面的相关接口)
      * [3\.2\.4\.1 搜索接口](#3241-搜索接口)
      * [3\.2\.4\.2 搜索历史 (废案)](#3242-搜索历史-废案)
      * [3\.4\.2\.3 热榜](#3423-热榜)
  * [3\.3 点歌页接口](#33-点歌页接口)
* [4\.追忆童年](#4追忆童年)
  * [4\.1 追忆童年主页相关接口](#41-追忆童年主页相关接口)
    * [4\.1\.1 推荐歌曲，根据click数降序获取10项(大家都在听)](#411-推荐歌曲根据click数降序获取10项大家都在听)
    * [4\.1\.2 获取歌曲列表](#412-获取歌曲列表)
  * [4\.2 原翻唱页相关接口](#42-原翻唱页相关接口)
    * [4\.2\.1 获取原唱相关信息](#421-获取原唱相关信息)
    * [4\.2\.2  获取用户翻唱列表并排序](#422--获取用户翻唱列表并排序)
  * [4\.3 歌曲页相关接口](#43-歌曲页相关接口)
    * [4\.3\.1 点赞接口](#431-点赞接口)
    * [4\.3\.2 当前歌曲的信息获取](#432-当前歌曲的信息获取)
    * [4\.3\.2  歌曲跳转(翻唱)](#432--歌曲跳转翻唱)
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
  * [6\.6 动态、评论点赞](#66-动态评论点赞)
  * [6\.7 动态热门搜索（最多十条）](#67-动态热门搜索最多十条)
  * [6\.8 大家的状态（最多十八条）](#68-大家的状态最多十八条)
  * [6\.9 点歌页歌曲推荐(最多30条)](#69-点歌页歌曲推荐最多30条)
* [管理员相关](#管理员相关)
  * [7\.1 动态、评论删除](#71-动态评论删除)
* [通用功能](#通用功能)
  * [8\.1点赞](#81点赞)
      * [3\.2\.1\.3 拉取用户中奖记录(废案)](#3213-拉取用户中奖记录废案)
      * [3\.2\.1\.3 抽奖确认(废案)](#3213-抽奖确认)

# 1. 微信授权

## 1.1 授权

GET /wx/jump2wechat/?redirect=  HTTP/1.1

成功：

Content-Type application/json

200 OK

需要先访问此接口接受redirect参数

失败:

`可能遇到401,用户未登录强制重定向进行授权登录`
## 1.2 假登录

POST /user  HTTP/1.1
```json
{
	"nickname": "string",
	"openid": string,
  	"avatar": string
}
```

成功：

Content-Type application/json

200 OK


需要先访问此接口接受redirect参数

失败:

```json
{
  "message": "用户不存在"
}
```

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
### 2.1.1 用户爱好选择

POST /hobby HTTP/1.1

成功：

Content-Type: application/json

```json
{
"hobby":[]string
}
```

200 OK



失败：

Content-Type: application/json

400 
```json
{
  "message": "error param"
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
    "avatar_visible": string,     	// 1：隐藏头像，0：不隐藏
    "phone_search": string,     	// 0：允许通过手机号查找，1：不允许
    "real_name_search": string,      	// 0：允许通过姓名查找，1：不允许
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

GET /userMsg HTTP1.1
query
"module": int//点歌1 唱歌2 点赞3 动态4

成功：

200 OK

Content-Type: application/json

//index项不止一条，index从0开始

```json
{
  "mySelections": {
      index:{
    "song_name": "string",
    "created_at": "string", //“2006-01-02 15:04:05”
      }
  },
  "mySongs": {
      index:{
	"id":int,
    "likes":int,
    "created_at": "string",
    "song_name": "string",
      }
  },
  "myLikes": {
      index:{
    "cover_id":int,
    "likes":int,
    "created_at": "string",
    "song_name": "string",
    "id": int, //对应点赞的id
    "likes": int //对应点赞数
      }
  },
  "moments": {
      index:{
    "created_at": "string",
    "state": "string", //状态:摸鱼
    "content": "string", //动态内容
    "id": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likes": int 
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
### 2.3.1 获取用户登录信息
GET /user
200 OK
Content-Type: application/json
```json
        "user_id":          int,
		"is_existed":       int,
		"avatar":           string,
		"nickname":         string,
		"is_administrator": bool,
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
  "song_name": "string",
  "selectionId": int,
  "name": "string",//点歌用户名
  "style": "string",//风格
  "created_at": "string", //“2006-01-02 15:04:05”
  "remark": "string" //30字以内
}
  "singers": {
      index{
    "nickname": "string",
    "id": int, //对应翻唱的id
    "user_id":int,
    "file": "string" ,//歌曲url
    "likes":int
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
  "module": int//点歌1 唱歌2 点赞3 动态4
}
```

成功：

200 OK

Content-Type: application/json

```json
{
  "message": {
    "avatar": "string",
    "nickname": "string",
    "school": "string",
    "signature": "string",
  },
  "mySelections": {
      index:{
    "song_name": "string",
    "created_at": "string", //“2006-01-02 15:04:05”
      }
  },
  "mySongs": {
      index:{
	"id":int,
    "likes":int,
    "created_at": "string",
    "song_name": "string",
     "likes":int
      }
  },
  "myLikes": {
      index:{
    "cover_id":int,
    "likes":int,
    "created_at": "string",
    "song_name": "string",
    "id": int, //对应点赞的id
    "likes": int //对应点赞数
      }
  },
  "moments": {
      index:{
    "created_at": "string",
    "state": "string", //状态:摸鱼
    "content": "string", //动态内容
    "id": int, //对应动态的id
    "song_name": string, //分享的歌曲名
    "likes": int 
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

### 3.1.1 轮播图歌手接口

***轮播图接口更新可能依赖于多于一张表***

GET /healing/devotion HTTP 1.1
query


成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
   "阿细":{index:
    {
      "devotion_id": int
      "song_name": string,
      "file": string,
      "likes": int
      
    }},
  "梁山山":{
    index: {
      "devotion_id": int
      "song_name": string,
      "file": string,
      "likes": int
    }
  }
]
```
//index不止一个

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "请求列表失败"}`

### 3.1.2 点歌列表获取

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
	selection_list:[{
		"nickname":string,	//可匿名
        "id":int,//对应点歌id
        "song_name":string,
        "user_id":integer,	//点歌用户的id
        "created_at":string(datetime),	//“2006-01-02 15:04:05”
        "avatar":string
	}],
  "page_num":int,
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
"page":int //页数 页数为1会刷新，若列表长度小于10，则到底
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
   cover_list:[{
       "nickname":string,	//可匿名
        "id":int,//对应翻唱歌id
        "song_name":string,
        "user_id":integer,	//翻唱用户的id
        "created_at":string(datetime),	//“2006-01-02 15:04:05”
        "avatar":string,
        "file":string//歌曲url
    }],
  page_num:int
    
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
    module:int, 1表示治愈系翻唱,2表示童年
    "is_anon": bool,      
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

### 3.1.6 电话治愈
GET /phoneNumber HTTP 1.1
query
"user_id"
成功：

Content-Type: application/json

`{"phone_number" : "string"}`

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

POST  /healing/lotterybox/draw HTTP 1.1

```json
{
    "tel":string
}
```

成功:

HTTP/1.1 200 OK

Content-Type: application/json

`{"message" : "抽奖成功，请耐心等待开奖"}`

`{"message" : "不能重复抽奖"}`

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "抽奖失败"}`

#### 3.2.1.4 拉取对应用户的任务列表

GET /healing/lotterybox/tasktable

成功:

HTTP/1.1 200 OK

Content-Type: application/json

**前端的任务上限确认由前端自行完成,只要用max和counter进行比较，若max<=counter,则任务已完成**

```json
[
    {
        "task": {
            "id": integer,	
            "text" : string,	// 任务描述					
            "max": integer	//上限积分				
        },
        "counter":integer	//已经获得的积分，交付前端表示进度
    },
    ...
]
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "加载任务列表失败"}`

#### 3.2.1.5 获取当前用户的积分

GET /healing/lotterybox/points

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "points":integer， //当前用户的积分
}
```

失败(例)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "获取当前用户积分失败"}`

### 3.2.2 排行榜

#### 3.2.2.1 学校积分排名

GET /healing/rank/:school

***全部使用学校全名***

***特别地，全部==All***



成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10
    {
        “userid”:integer, 	//userid
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

Get /healing/dailyrank/:date
***date的日期遵循统一格式 mm-dd*** 
***日期不早于上线当日***

点赞调用 /healing/covers/like 接口

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10
    {
        "cover_id":integer,	//跳转翻唱id
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

Get /healing/dailyrank/all

**全部时间获赞数最多的翻唱**

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[//列表长度为10
    {
        "cover_id":integer,	//翻唱id，跳转
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

#### 3.2.4.2 搜索历史 (废案)

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
//童年歌曲点歌的格式（分享也按点歌这个进行操作）
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

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

### 4.1.2 获取歌曲列表

GET /childhood/list HTTP 1.1

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
[// 一次性把所有的童年歌曲信息都发出去了
    {
       	"classic_id":int, 	//用于跳转到对应的原翻唱页
    	"name":string,		//歌曲名
    	"icon":text(url),	//歌曲图标
        "work_name":string	//作品名
    }
    ...
]
```

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

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
{// 成功就一定会返回且仅返回一条记录
    "classic_url":string, 	//用于播放原唱
    "song_name": string,
    "singer": string,
    "icon":text(url),   //歌曲图标
    "work_name":string	//作品名
}
```

失败（说明参数缺失或者非法）：

HTTP/1.1 400 

Content-Type: application/json

`{"message" : "传入参数非法"}`

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

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
[// 当这首歌曲暂未有人翻唱时，返回值会是 null
    {
        "cover_id":int,			// 用于进入歌曲页
        "nickname": string,		// 用户昵称
        "avatar": text(url),   	// 用户头像
        "post_time": datetime,	// “2006-01-02 15:04:05”
    }
    ...
]
```

失败（说明参数缺失或者非法）：

HTTP/1.1 400 

Content-Type: application/json

`{"message" : "传入参数非法"}`

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

## 4.3 歌曲页相关接口

### 4.3.1 点赞接口

调用 8.1 通用点赞接口

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
    "work_name":string,	//作品名
	"check":int		//0表示当前用户未点赞，1表示当前用户已经点赞
}
```

失败（说明参数缺失或者非法）：

HTTP/1.1 400 

Content-Type: application/json

`{"message" : "传入参数非法"}`

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

### 4.3.2  歌曲跳转(翻唱)

POST /healing/covers/jump

```json
{
    "jump":integer,		// 0为上一首,1为下一首
    "check":integer,	// 0为经典治愈，1为童年
    "cover_id":integer	// 当前的翻唱 id
    
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

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "Check参数错误"}`

失败（说明参数缺失或者非法）：

HTTP/1.1 400 

Content-Type: application/json

`{"message" : "参数不完整"}`

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}`

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

HTTP/1.1 403 

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

HTTP/1.1 403 

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

HTTP/1.1 403 

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

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "消息不存在!"}`

# <span id="">6. 广场主页</span>

## <span id="">6.1 拉取广场动态列表</span>

GET /dynamics/list/{method}  HTTP1.1

其中 method 可取： "new"(时间排序)/"recommend"（点赞数+评论数排序）/"search"（动态中含有关键字）

当选用的 method 为 "search" 时,在 url 后加上 ?keyword=xxx 即可拉取含有关键字的 状态、歌曲名 的动态列表

为实现分页功能（10条动态为一页），url 后**必须**加上 ?page=xx 即可取得第 xx 页的动态列表（page=0为第一页）

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[// 当根本没有任何人发过动态时响应为空 null, 没有动态含有关键字时 search 为空 null
    {
        "dynamics_id": integer,
        "content": string,								// 动态的内容
        "created_at": "2006-01-02 15:04:05",
        
        "song" : string,								// 要点的、分享的歌名
        "song_id": int,									// 经典点歌时为selection_id;童年分享时为classic_id
        "module" : int,									// 0为经典，1为童年，2为无歌曲
        
        "lauds" : integer,								// 动态的点赞数
        "lauded": integer(0/1),							// 当前用户是否点赞
        "comments" : integer,							// 动态的评论数
        "status" : ["status1", "status2" ...],			// 状态列表

        "creator": {
            "id": integer,	
            "nickname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 0代表没设置头像
        }
    },
    ...
]
```

失败时：

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "page参数非法"}`

HTTP/1.1 404 (method参数不正确)

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}



## <span id="">6.2 发布动态</span>

POST  /dynamics/send  HTTP1.1

Content-Type: application/json

```js
{
    "content": string,							// 动态的内容
    "status": ["status1", "status2"...],		// 状态列表元素都是string
        
    "have_song":int,						    // 经典点歌为0，，童年分享为1，无歌曲为2
   
    // 经典点歌填这四个
    "song_name":string,	
    "language":string,	
    "style":string,
    "remark":text,								//点歌的备注
     
     // 童年分享填这一个
    "classic_id":int
}
```

成功时：

HTTP/1.1 200 OK

`{"message" : "评论发布成功"}`

失败时：

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "评论参数不完整"}`

失败（说明后台挂了）：

HTTP/1.1 500

Content-Type: application/json

`{"message" : "数据库操作失败"}

`{"message" : "系统消息发送失败"}

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
    "song_id" : string, 							// 歌曲id
    "module" : int,									// 1为经典，2为童年，0为无歌曲
    "lauds" : integer,								// 动态的点赞数
    "lauded": integer(0/1),							// 当前用户是否点赞该动态
    "comments" : integer,							// 动态的评论数
    "status" : ["status1", "status2" ...],			// 状态列表

    "creator": {
        "id": integer,	
        "nickname" : string,						// 用户名
        "avatar": string(url),						// 头像
    	"avatar_visible": integer(0/1)				// 0代表没设置头像
	}
}

```

失败时：

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "id参数未传入"}`

`{"message" : "id参数非法"}`（传入的不是数字）

HTTP/1.1 500 

Content-Type: application/json

`{"message" : "数据库操作失败"}`

`{"message" : "数据库中出现非法字段"}`

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

失败时：

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## <span id="">6.5 拉取动态的评论列表</span>

GET /dynamics/comment/{id}  HTTP1.1

id 为动态对应的 id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[// 动态没有评论时为 null
    {
        "comment_id": integer,
        "content": string,
        "created_at": "2006-01-02 15:04:05",
        "lauds" : integer,								// 动态的点赞数
    	"lauded": integer(0/1),							// 当前用户是否点赞
    
        "creator": {
            "id": integer,	
            "nickname" : string,						// 用户名
            "avatar": string(url),						// 头像
            "avatar_visible": integer(0/1)				// 是0代表没设置头像
        }
    },
    ...
]
```

失败时(例子)：

HTTP/1.1 403 

Content-Type: application/json

`{"message" : "该动态不存在!"}`

## 6.6 动态、评论点赞

使用 8.1 通用点赞接口进行点赞

## 6.7 动态热门搜索（最多十条）

GET /dynamics/hotsearch HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    string1, string2,...
]
```

失败时(例子)：

HTTP/1.1 500 

Content-Type: application/json

```
{"message" : "服务端出错"}
```

## 6.8 大家的状态（最多十八条）

GET /dynamics/ourstates HTTP1.1

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

```json
[
    state1, state2,...
]
```

失败时(例子)：

HTTP/1.1 500 

Content-Type: application/json

```
{"message" : "服务端出错"}
```

## 6.9 点歌页歌曲推荐(最多30条)

GET /dynamics/hotsong HTTP1.1

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

HTTP/1.1 500 

Content-Type: application/json

```
{"message" : "服务端出错"}
```

# 管理员相关

## 7.1 动态、评论删除

POST /administrators/content/?id= HTTP1.1

query

id

成功时：

HTTP/1.1 200 OK

Content-Type: application/json

失败时(例子)：

HTTP/1.1 403 Forbidden

Content-Type: application/json

```
{"message" : "无权限"}
```

# 通用功能

## 8.1点赞

PUT/like HTTP1.1

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

HTTP/1.1 405 Forbidden

Content-Type: application/json

```
{"message" : "不允许重复点赞"}
```

## 

#### 3.2.1.3 拉取用户中奖记录(废案)

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



#### 3.2.1.3 抽奖确认(废案)

GET /healing/lotterybox/drawcheck HTTP 1.1

**作为抽奖按钮的前置接口存在**

成功:

HTTP/1.1 200 OK

Content-Type: application/json

```json
{
    "msg":"积分不足"
}
```

or

```json
{
    "msg":"请填写手机号码"
}
```

or

```json
{
    "msg":"已参与抽奖
}
```

HTTP/1.1 403 Forbidden

Content-Type: application/json

`{"message" : "抽奖失败"}

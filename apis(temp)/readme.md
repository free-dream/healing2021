- [1. 微信授权](#head1)
    - [1.1 授权](#head2)
- [ -----下面的接口都要带上前缀"/api"-----](#head3)
- [2. 用户模块](#head4)
    - [2.1 用户注册](#head5)
    - [2.2 个人信息更新](#head6)
    - [2.3 获取自己信息(用户个人页信息拉取)](#head7)
    - [2.4 更换个人背景](#head8)
    - [2.5 治愈详情页](#head9)
    - [2.6 获取他人信息](#head10)
    - [2.7 获取二维码](#head11)

# <span id="head1">1. 微信授权</span>

## <span id="head2">1.1 授权</span>

GET /auth/jump2[?redirect={encoded_uri}] HTTP/1.1

### 成功：
Content-Type application/json

200 OK
```json
{
  "nickname": "string",
}
```

需要先访问此接口接受redirect参数
### 失败:
`可能遇到401,用户未登录强制重定向进行授权登录`

# <span id="head3"> -----下面的接口都要带上前缀"/api"-----</span>

# <span id="head4">2. 用户模块</span>

## <span id="head5">2.1 用户注册</span>

POST /user HTTP/1.1
### 成功：
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
### 失败：
Content-Type: application/json

403 Forbidden
```json
{
  "message": "昵称/手机号已存在,无法注册"
}
```
## <span id="head6">2.2 个人信息更新</span>

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

## <span id="head7">2.3 获取自己信息(用户个人页信息拉取)</span>

GET /user HTTP1.1

成功时：

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

失败时：

403 Forbidden

Content-Type: application/json

```json
{"message" : "获取头像失败"}
```

## <span id="head8">2.4 更新个人背景</span>
POST /background HTTP/1.1
### 成功：

200 OK

Content-Type: application/json

```json
{
  "background": "string" //背景图片url
}
```
### 失败：
500 Internal Server Error
```json
{"message" : "更新头像失败"}
```
## <span id="head9">2.5 治愈系详情页</span>
GET /healingPage HTTP/1.1

Query
```json
{
  "healingId": int
}
```
### 成功：
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
### 失败:
```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```




## <span id="head10">2.6 获取他人信息</span>
GET /callee HTTP/1.1

Query

```json
{
  "calleeId": int //被访问用户的id
}
```

### 成功：
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


### 失败:
```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```

## <span id="head11">2.7 获取二维码</span>
GET /QR_code HTTP/1.1

Query
```json
{
  "userId": int,//发起请求的用户id
  "songId": int,//对应歌曲id
}
```
### 成功：
```json
{
  "QR_code": "string",//二维码url
}
```
### 失败:
```json
{
  "statusCode": 401,
  "message": "parameter error"
}
```
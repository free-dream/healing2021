# mysql更新和redis缓存的统一工具
* 基于提高效率的考虑，设置一个统一处理mysql读取和缓存到redis的管理层
* 后续会画一个详细的构造图出来

1. 点赞操作和数值缓存
2. 预缓存热榜/用户信息/任务
3. 此处的接口都要求前端传回带openid的cookie

## 缓存表

**()内注明golang数据类型,用于读取断言，无特别说明默认string**

### 翻唱排序缓存表

**所有的排序表都设置定时(推荐5min)更新，更新前先更新点赞表**

**为了提高效率，所以每条记录单独提出来处理了，存储了数据**

**总共是8\*2\*10条记录**

```
style={Recommend,Pop,Vintage}
language={All,Chinese,Cantonese,English,Japanese}
//style和language在排序时是平行关系，所以是5+3=8
request={composite,latest}
//标记：“Cantonese:粤语;Vintage:古风;composite:综合”
```

| key(i=1-10)                  | field1      | field2 | field3      | field4   | field5         | field6     | field7 | field8 | field9 | field10  |
| ---------------------------- | ----------- | ------ | ----------- | -------- | -------------- | ---------- | ------ | ------ | ------ | -------- |
| {style}/{request}/cover**i** | userid(int) | avatar | selectionid | songname | classicid(int) | likes(int) | file   | style  | module | language |

### 点歌排序缓存表

```
//注释和翻唱缓存表保持一致
```

| key(i=1-15)                      | f1       | f2     | f3       | f4    | f5          | f6     | f7     |
| -------------------------------- | -------- | ------ | -------- | ----- | ----------- | ------ | ------ |
| {style}/{request}/selection**i** | songname | remark | language | style | userid(int) | avatar | module |

### 用户缓存表

**用户信息表,长期使用且基本不变化，有变化直接写入数据库并缓存**

**首次登录时缓存**

**expire 7\*24h**

**key:info{userid}，例如info5**

| key          | f1          | f2       | f3       | f4        | f5     | f6          | f7   | f8     | f9         | f10              | f11              | f12                 |
| ------------ | ----------- | -------- | -------- | --------- | ------ | ----------- | ---- | ------ | ---------- | ---------------- | ---------------- | ------------------- |
| info{userid} | userid(int) | nickname | realname | signature | avatar | phonenumber | sex  | school | background | avatarvisible(0) | phonesearch(int) | realnamesearch(int) |

### 奖品缓存表

**展示所有奖品,只缓存一次**

| key              | prize  | possibility     |
| ---------------- | ------ | --------------- |
| lottery{prizeid} | string | string(float64) |



### 任务缓存表

**任务是人为设计的，所以不过期，如果有更新直接更新数据库**

**系统初始化的时候直接缓存**

**key:task{taskid}，例如:task1**

| key          | f1   | f2     |
| ------------ | ---- | ------ |
| task{taskid} | text | target |

### 校园积分排行榜

**每小时更新一次，更新前对mysql表进行更新**

**有序集合zset,zset名为学校名+rank,例如scutrank/allrank**

**有n个学校所以有n+1张表，从userid获取头像（redis或mysql）**

**可以确定，排行榜首位一定也是常使用产品的，redis大概率有其数据**

| name       | number | value  |
| ---------- | ------ | ------ |
| allrank    | score  | userid |
| （待补充） |        |        |

## 中频表

**抽奖箱** //我的接口，完成抽奖箱的时候搞定

## 高频表

### 用户积分

**用户积分表,高频更新,key是{userid}/point的字符串，例如:64/point**

**score只增不减,point有增有减，完成任务record/point增加，抽奖point减少**

**每小时更新排行榜时进行更新**

| key            | f1          | f2         |
| -------------- | ----------- | ---------- |
| {userid}/point | record(int) | point(int) |

### 任务表

**用户任务表,记录任务完成情况和目标,expile:24h**

**登录时缓存一次,之后定时更新**

| key                         | f1           | f2         |
| --------------------------- | ------------ | ---------- |
| {userid}/task/{tasktableid} | process(int) | check(int) |



### 点赞表

**三个有序集合zset，点赞即时更新**

**热榜更新时以此为增量更新关联数据库，更新之后表内数据全部置0**

**Zincrby增长redis记录**

**尝试数据库批处理**

| name    | score      | value     |
| ------- | ---------- | --------- |
| cover   | likes(int) | coverid   |
| moment  | likes(int) | momentid  |
| comment | likes(int) | commentid |

#### 点赞记录表

**记录用户点赞记录，24h不允许重复点赞**

**一个用户对应3个集合：**

**user{i}cover,user{i}moment,user{i}comment**

**sismember查找以确定是否已点赞**

**例：**

1. user{i}cover{1,2,3,4,5,6,7,8}
2. user{i}moment{2,4,6,7}

3. user{i}comment{3,5,6,7}


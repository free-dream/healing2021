# mysql更新和redis缓存的统一工具
* 基于提高效率的考虑，设置一个统一处理mysql读取和缓存到redis的管理层
* 后续会画一个详细的构造图出来

1. 点赞操作和数值缓存
2. 预缓存热榜/更新热榜
3. 此处的接口都要求前端

## 点赞表

**expire 24h**

| key    | field2    |
| ------ | --------- |
| userid | coverid   |
| userid | momentid  |
| userid | commentid |

## 记录表

即时更新

| key       | value        |
| --------- | ------------ |
| coverid   | likes(int)   |
| momentid  | likes(int)\| |
| commentid | likes(int)   |

## 缓存表

### 翻唱缓存表

```
style={Recommend,Pop,Vintage}
language={All,Chinese,Cantonese,English,Japanese}
request={composite,latest}
//标记：“Cantonese:粤语;Vintage:古风;composite:综合”
//()内注明golang数据类型,用于读取断言，无特别说明默认string
```

| key(i=1-15)                  | field1      | field2 | field3      | field4   | field5         | field6     | field7 | field8 | field9 | field10  |
| ---------------------------- | ----------- | ------ | ----------- | -------- | -------------- | ---------- | ------ | ------ | ------ | -------- |
| {style}/{request}/cover**i** | userid(int) | avatar | selectionid | songname | classicid(int) | likes(int) | file   | style  | module | language |

### 点歌缓存表

```
和翻唱缓存表保持一致
```

| key(i=1-15)                      | f1       | f2     | f3       | f4    | f5          | f6     | f7     |
| -------------------------------- | -------- | ------ | -------- | ----- | ----------- | ------ | ------ |
| {style}/{request}/selection**i** | songname | remark | language | style | userid(int) | avatar | module |


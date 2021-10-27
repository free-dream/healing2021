#  Database

# MySQL

注意使用utf8mb4编码

GORM表自动迁移时复合命名的字段会自动变为下划线,且自动生成时间为datetime类型

thumbnail可以通过url的变换来获取，但暂时保留thumbnail字段

## healing

### user

//增加了个性签名、历史最高累计积分和thumbnail

| 字段值       | 字段类型     | 说明                          |
| ------------ | ------------ | ----------------------------- |
| id           | int          | 自增主键                      |
| avatar       | varchar(255) | 头像url                       |
| signature    | varchar(255) | 个签，可选                    |
| nickname     | varchar(255) |                               |
| real_name    | varchar(255) | 选填项                        |
| phone_number | char(11)     | 选填项                        |
| sex          | int          | 1男,2女,3其他                 |
| school       | varchar(255) | 大学名，用于healing排行榜     |
| points       | int          | 初始0，用于记录抽奖，动态变化 |
| record       | int          | 记录历史最高积分              |
| background   | varchar(255) | 背景图片url                   |
| openid       | int          | 微信颁发的用户openid          |

### selection（点歌）

| 字段值         | 字段类型     | 说明               |
| -------------- | ------------ | ------------------ |
| id             | int          | 主键，自增         |
| song_name      | varchar(255) |                    |
| remark         | varchar(255) | 备注               |
| language       | varchar(255) |                    |
| style          | varchar(255) |                    |
| user_id        | int          | 点歌用户的id       |
| module         | varchar(255) | 模块名             |

### moment（动态）

| 字段值       | 字段类型     | 说明     |
| ------------ | ------------ | -------- |
| id           | int          | 自增主键 |
| user_id           | int          |         |
| like_num | int          |          |
| content      | varchar(255) |          |
| selection_id | int          |          |
| song_name    | varchar(255) |          |
| states       | varchar(255) | 状态     |
| picture      | varchar(255) | 图片url  |


### moment_comment（评论）

| 字段值    | 字段类型     | 说明     |
| --------- | ------------ | -------- |
| id        | int          | 自增主键 |
| user_id   | int          | 评论用户 |
| moment_id | int          | 外键     |
| comment   | varchar(255) | 评论内容 |

### advertisement

//可拓展的广告表，根据广告费用加权，提高展示率，用于轮播图

| 字段值  | 字段类型     | 说明                 |
| ------- | ------------ | -------------------- |
| id      | int          | 自增主键             |
| url     | varchar(255) | 广告图资源地址       |
| address | varchar(255) | 广告指向的外链(可选) |
| weight  | int          | 权重，提高轮播频率   |

### lottery

//彩票盒子

| 字段值      | 字段类型     | 说明                        |
| ----------- | ------------ | --------------------------- |
| id          | int          | 自增主键                    |
| user_id     | int          | default=-1,用于确定奖品归属 |
| name        | varchar(255) | 奖品名                      |
| picture     | varchar(255) | 奖品图像的url               |
| possibility | Double(2,2)  | 小于100，保留两位小数       |

### prize 

//用户中奖记录

| 字段值  | 字段类型 | 说明                       |
| ------- | -------- | -------------------------- |
| id      | int      | 自增主键                   |
| user_id | int      | 外键，绑定用户             |
| prize   | int      | 外键，绑定用户中的奖       |
| date    | datetime | 中奖时间，用于中奖记录排序 |

### task_table

//保证任务的可扩展性，另外拉取tasks表记录不同的任务类型和数据

//每日进行次数的更新

//初始化用户时要对任务进行分配和初始化

| 字段值  | 字段类型 | 说明                 |
| ------- | -------- | -------------------- |
| id      | int      | 自增主键             |
| task_id | int      | 外键，关联任务       |
| user_id | int      | 外键，关联用户       |
| check   | int      | 1/0表示已完成/未完成 |
| counter | int      | 非负，记录进行次数   |

### task

//记录不同任务的说明以及可能的次数要求，需要的时候可直接添加进数据库

//可根据需要设置针对特定用户的任务

| 字段值 | 字段类型     | 说明                                   |
| ------ | ------------ | -------------------------------------- |
| id     | int          | 自增主键                               |
| text   | varchar(255) | 任务描述                               |
| target | int          | 非负，记录目标数，用于确认任务是否完成 |



### classic（原唱）

| 字段值    | 字段类型     | 说明                 |
| --------- | ------------ | -------------------- |
| remark    | varchar(255) | 备注                 |
| id        | int          | 自增主键             |
| song_name | varchar(255) | 歌曲名称             |
| icon      | varchar(255) | 大图标url            |
| singer    | varchar(255) | 原唱歌手             |
| work_name | varchar(255) | 出处作品名称         |
| click     | int          | 听歌人数，每小时更新 |
| record    | varchar(255) | 原唱歌曲的录音url    |

### cover(翻唱)

//一个用户界面、经典治愈、童年界面、排行榜通用的表，基于不同接口提供不同的视图

| 字段值     | 字段类型     | 说明                               |
| ---------- | ------------ | ---------------------------------- |
| id         | int          | 自增主键                           |
| user_id    | int          | 用于用户界面                       |
| avatar     | varchar(255) | 头像url,童年/原唱界面视图          |
| nickname   | varchar(255) | 用户名,童年/原唱界面视图(可选匿名) |
| classic_id | int          | 童年/原唱界面视图索引              |
| likes      | int          | 点赞数，从redis持久化，初始为0     |
| file       | varchar(255) | 录音文件存储url                    |
| module     | int          | 1治愈系2童年                       |

### praise

| 字段值            | 字段类型 | 说明                      |
| ----------------- | -------- | ------------------------- |
| id                | int      | 自增主键                  |
| cover_id          | int      | 动态id(外键)              |
| user_id           | int      | 点赞用户的id              |
| is_liked          | int      | 点赞1/取消点赞0(保留功能) |
| moment_id         | int      |                           |
| moment_comment_id | int      |                           |

### message

| 字段值    | 字段类型     | 说明      |
| --------- | ------------ | --------- |
| id        | int          | 自增主键  |
| sender_id | int          | 0作为系统 |
| taker_id  | int          |           |
| content   | varchar(255) |           |

## Redis

### 每日热榜

| 数据类型 | key                    | score    | member               |
| -------- | ---------------------- | -------- | -------------------- |
| zset     | 按日期命名(yyyy-mm-dd) | like_num | cover_id(song_likes) |

### 排行榜

| 数据类型 | key                         | score | member  |
| -------- | --------------------------- | ----- | ------- |
| zset     | 按学校命名(总榜命名为total) | point | user_id |

`point值可用zincrby改变`

### 点赞缓存

点赞关系

| 数据类型    | key                           | field    | value |
| ----------- | ----------------------------- | -------- | ----- |
| hash(hmset) | 按对应点赞表命名(例song_like) | user_id  | int   |
|             |                               | song_id  | int   |
|             |                               | is_liked | int   |

点赞数

| 数据类型 | key                               | field    | value |
| -------- | --------------------------------- | -------- | ----- |
| hmset    | 按对应点赞表命名(例song_like_num) | song_id  | int   |
|          |                                   | like_num | int   |

`用hincrby记录点赞数`


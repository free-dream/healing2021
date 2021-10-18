# sandwich——mysql更新和redis缓存的统一管理插件
* 基于提高效率的考虑，设置一个统一处理mysql读取和缓存到redis的管理层
* 后续会画一个详细的构造图出来

1. 点赞操作和数值缓存
2. 预缓存热榜/更新热榜
3. 此处的接口都要求前端

### 点赞表

**expire 24h**

| key  | field1 | field2    |
| ---- | ------ | --------- |
| uuid | userid | coverid   |
| uuid | userid | momentid  |
| uuid | userid | commentid |

## 记录表

即时更新

| key      | value        |
| -------- | ------------ |
| coverid  | likes(int)   |
| momentid | likes(int)\| |

| key        | field1    | field2 |
| ---------- | --------- | ------ |
| comment_id | moment_id | likes  |


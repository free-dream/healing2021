# 抽奖池的机制设计

* 每次随机生成包含奖品参数在内的1000张卡牌，根据可能性进行卡牌数分配
* 用户中奖，进行信息查询，若奖品已有归属，同样返回未中奖，不修改数据库
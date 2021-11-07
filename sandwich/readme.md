# sandwich声明了redis的存取工具
    voloroloq 2021.11.7
* 对于大量的数据比如排行榜建议使用encoding/json,marshal后存为json字符串缓存于redis中，提取时unmarshal为结构体数组
* 目前的缓存采取主动更新策略，使用cron定期删除key,原因在于expile不好控制绝对时间
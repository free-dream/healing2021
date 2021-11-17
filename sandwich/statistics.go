package sandwich

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/go-redis/redis"
)

// redis 添加两个 Zset 来统计搜索热词和热门状态

// ZSet 对应的键
const ourStates = "healing2021:ourStates"
const hotSearch = "healing2021:hotSearch"
const hotSong = "healing2021:hotSong"

func init()  {
	for i := 0; i < 20; i++ {
		PutInStates("上早课")
		PutInStates("奔赴饭堂")
		PutInStates("乐跑")
		PutInStates("游泳")
		PutInStates("等待外卖")
		PutInStates("取快递")
		PutInStates("学习中")
		PutInStates("工作中")
		PutInStates("充能中")
		PutInStates("奋斗中")
		PutInStates("求锦鲤")
		PutInStates("芜湖起飞")
		PutInStates("喝奶茶")
		PutInStates("自拍")
		PutInStates("美滋滋")
		PutInStates("滋润")
		PutInStates("期待满满")
		PutInStates("元气满满")
		PutInStates("摸鱼中")
		PutInStates("纠结中")
		PutInStates("叹气")
		PutInStates("网抑云")
		PutInStates("头大中")
		PutInStates("熬夜中")
		PutInStates("Emo中")
		PutInStates("自闭中")
		PutInStates("社死")
		PutInStates("想静静")
		PutInStates("裂开")
		PutInStates("赶DDL")
		PutInStates("生病中")
		PutInStates("不想学习")
		PutInStates("睡不醒")
		PutInStates("拒绝早八")
	}
}

// 放入状态
func PutInStates(States string) {
	redisDb := setting.RedisConn()
	err := redisDb.ZIncrBy(ourStates, 1, States).Err()
	if err != nil {
		fmt.Println("PutInStates error")
	}
}

// 返回前 36 条状态
func GetStates() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前36个状态
	op := redis.ZRangeBy{
		Min:    "-inf", //最小分数
		Max:    "+inf", //最大分数
		Offset: 0,      // 类似sql的limit, 表示开始偏移量
		Count:  36,     // 一次返回多少数据
	}
	Values, err := redisDb.ZRevRangeByScore(ourStates, op).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	return Values
}

// 放入搜索词
func PutInSearchWord(Word string) {
	redisDb := setting.RedisConn()
	err := redisDb.ZIncrBy(hotSearch, 1, Word).Err()
	if err != nil {
		fmt.Println("PutInSearchWord error")
	}
}

// 取出前 10 条搜索词
func GetSearchWord() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前10条搜索记录
	op := redis.ZRangeBy{
		Min:    "-inf", //最小分数
		Max:    "+inf", //最大分数
		Offset: 0,      // 类似sql的limit, 表示开始偏移量
		Count:  10,     // 一次返回多少数据
	}
	Values, err := redisDb.ZRevRangeByScore(hotSearch, op).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	return Values
}

// 放入动态热门点歌数据(编码解码需要在外部做)
func PutInHotSong(Songinfo string) {
	redisDb := setting.RedisConn()
	err := redisDb.ZIncrBy(hotSong, 1, Songinfo).Err()
	if err != nil {
		fmt.Println("PutInHotSong error")
	}
}

// 获取动态热门点歌数据 前30条最多(编码解码需要在外部做)
func GetHotSong() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前 30 条点歌信息
	op := redis.ZRangeBy{
		Min:    "-inf", //最小分数
		Max:    "+inf", //最大分数
		Offset: 0,      // 类似sql的limit, 表示开始偏移量
		Count:  10,     // 一次返回多少数据
	}
	Values, err := redisDb.ZRevRangeByScore(hotSong, op).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	return Values
}

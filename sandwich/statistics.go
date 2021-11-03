package sandwich

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"github.com/go-redis/redis"
)

// redis 添加两个 Zset 来统计搜索热词和热门状态

// ZSet 对应的键
const ourStates = "ourStates"
const hotSearch = "hotSearch"
const hotSong = "hotSong"

// 放入状态
func PutInStates(States string) {
	redisDb := setting.RedisConn()
	err := redisDb.ZIncrBy(ourStates, 1, States).Err()
	if err != nil {
		fmt.Println("PutInStates error")
	}
}

// 返回前 18 条状态
func GetStates() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前18个状态
	op := redis.ZRangeBy{
		Min:    "-inf", //最小分数
		Max:    "+inf", //最大分数
		Offset: 0,      // 类似sql的limit, 表示开始偏移量
		Count:  18,     // 一次返回多少数据
	}
	Values, err := redisDb.ZRevRangeByScore(ourStates, op).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	// 测试
	for _, val := range Values {
		fmt.Println(val)
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

	// 测试
	for _, val := range Values {
		fmt.Println(val)
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

	// 测试
	for _, val := range Values {
		fmt.Println(val)
	}

	return Values
}

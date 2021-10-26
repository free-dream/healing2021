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

// 放入状态
func PutInStates(States string) {
	redisDb := setting.RedisConn()
	redisDb.ZIncrBy(ourStates, 1, States)
}

// 返回前18条状态
func GetStates() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前18个状态
	op := redis.ZRangeBy{
		Offset: 0,  // 类似sql的limit, 表示开始偏移量
		Count:  18, // 一次返回多少数据
	}
	values, err := redisDb.ZRevRangeByScore(ourStates, op).Result()
	if err != nil {
		panic(err)
	}

	// 测试
	for _, val := range values {
		fmt.Println(val)
	}

	return values
}

func PutInSearchWord(Word string) {
	redisDb := setting.RedisConn()
	redisDb.ZIncrBy(hotSearch, 1, Word)
}

func GetSearchWord() []string {
	redisDb := setting.RedisConn()

	// 设置查找要求并找到前10条搜索记录
	op := redis.ZRangeBy{
		Offset: 0,  // 类似sql的limit, 表示开始偏移量
		Count:  10, // 一次返回多少数据
	}
	values, err := redisDb.ZRevRangeByScore(hotSearch, op).Result()
	if err != nil {
		panic(err)
	}

	// 测试
	for _, val := range values {
		fmt.Println(val)
	}

	return values
}

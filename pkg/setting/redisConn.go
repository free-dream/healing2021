package setting

import (
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client
var TokenGetCli *redis.Client

func init() {
	addr := tools.GetConfig("redis", "addr")
	dbStr := tools.GetConfig("redis", "db")
	var db int
	if dbStr == "" {
		db = 0
	} else {
		db, _ = strconv.Atoi(dbStr)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           db,
		PoolSize:     30,
		MinIdleConns: 10,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	TokenGetCli = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 10,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func RedisConn() *redis.Client {
	return RedisClient
}
func RedisTokenConn() *redis.Client {
	return TokenGetCli
}

package dao

import (
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"time"
)

func Filter(key string, value string) bool {
	redisCli := setting.RedisConn()
	return redisCli.SetNX(key, value, time.Minute*5).Val()
}

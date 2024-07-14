package redis

import (
	"github.com/go-redis/redis/v8"
	"goAiproject/global"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password,
		DB:       global.RedisSetting.DB,
	})
	return client
}

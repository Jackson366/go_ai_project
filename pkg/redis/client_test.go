package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	ping := client.Ping(ctx)
	t.Log(ping)
}

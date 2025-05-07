package redis

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	RDB     *redis.Client
	Limiter *redis_rate.Limiter
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return &RedisClient{
		RDB:     rdb,
		Limiter: redis_rate.NewLimiter(rdb),
	}
}

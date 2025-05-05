package redis

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	RDB     *redis.Client
	Limiter *redis_rate.Limiter
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	Limiter = redis_rate.NewLimiter(RDB)
}

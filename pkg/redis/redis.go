package redis

import (
	"github.com/go-redis/redis"
	"runtime"
)

var RedisClient *redis.Client

func NewRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.137.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 20 * runtime.NumCPU(),
	})
}

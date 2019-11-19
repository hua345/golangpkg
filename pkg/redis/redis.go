package redis

import (
	"github.com/go-redis/redis"
	"runtime"
	"sync"
)

var (
	once sync.Once

	redisClient *redis.Client
)

func GetInstance() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			PoolSize: 20 * runtime.NumCPU(),
		})
	})

	return redisClient
}

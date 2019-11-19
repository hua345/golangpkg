package redigo

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"runtime"
	"sync"
	"time"
)

var redisHost = "127.0.0.1:6379"
var redisPassword = ""

var (
	once sync.Once

	instance *redis.Pool
)

func GetInstance() *redis.Pool {
	once.Do(func() {
		instance = &redis.Pool{
			MaxIdle:     20 * runtime.NumCPU(),
			IdleTimeout: 5 * time.Second,
			Dial: func() (redis.Conn, error) {
				pool, err := redis.Dial("tcp",
					redisHost,
					redis.DialPassword(redisPassword),
					redis.DialDatabase(0))
				if err != nil {
					return nil, err
				}
				return pool, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	})

	return instance
}

func Set(key string, data interface{}) error {
	// 从池里获取连接
	conn := GetInstance().Get()
	// 用完后将连接放回连接池
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
func SetStr(key string, value string) error {
	// 从池里获取连接
	conn := GetInstance().Get()
	// 用完后将连接放回连接池
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
func SetStrWithExpire(key string, value string, second int) error {
	// 从池里获取连接
	conn := GetInstance().Get()
	// 用完后将连接放回连接池
	defer conn.Close()
	_, err := conn.Do("SET", key, value, "EX", second)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	// 从池里获取连接
	conn := GetInstance().Get()
	// 用完后将连接放回连接池
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func GetStr(key string) (string, error) {
	// 从池里获取连接
	conn := GetInstance().Get()
	// 用完后将连接放回连接池
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := GetInstance().Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

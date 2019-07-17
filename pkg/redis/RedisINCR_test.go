package redis

import (
	"sync"
	"testing"
)

func redisIncrWork(wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	testOrderId := "testOrderId"
	_, err := RedisClient.IncrBy(testOrderId, 1).Result()
	if err != nil {
		t.Error(err)
	}
}
func redisIncr() (int64, error) {
	testOrderId := "testOrderId"
	orderId, err := RedisClient.IncrBy(testOrderId, 1).Result()
	if err != nil {
		return -1, err
	}
	return orderId, nil
}

// 单元测试
func TestRedisIncr(t *testing.T) {
	NewRedis()
	orderId, err := redisIncr()
	if err != nil {
		t.Error(err)
	}
	t.Log("Begin orderId :", orderId)
	wg := sync.WaitGroup{}
	wg.Add(100)
	// create two separate sessions for lock competition
	for i := 0; i < 100; i++ {
		go redisIncrWork(&wg, t)
	}
	wg.Wait()
	orderId, err = redisIncr()
	if err != nil {
		t.Error(err)
	}
	t.Log("End orderId :", orderId)
}

// 性能测试
//go test -bench=.
func BenchmarkRedisINCR(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	defaultOrderId := "defaultOrderId"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		_, err := RedisClient.IncrBy(defaultOrderId, 1).Result()
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisINCRParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	defaultOrderId := "defaultOrderId"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := RedisClient.IncrBy(defaultOrderId, 1).Result()
			if err != nil {
				b.Error(err)
			}
		}
	})
}

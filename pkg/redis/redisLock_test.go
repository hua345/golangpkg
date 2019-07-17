package redis

import (
	"sync"
	"testing"
	"time"
)

func redisLockWork(index int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	redisSession := NewRedisSession(RedisClient)
	redisLock, err := redisSession.TryLock("myLock", 10*time.Second)
	if err != nil {
		t.Log(err)
	}
	t.Log("redis Lock Success", index)
	time.Sleep(100 * time.Millisecond)
	t.Log("released lock for session", index)
	err = redisLock.Release()
	if err != nil {
		t.Log(err)
	}
}
func TestRedisLock(t *testing.T) {
	NewRedis()
	wg := sync.WaitGroup{}
	wg.Add(100)
	// create two separate sessions for lock competition
	for i := 0; i < 100; i++ {
		go redisLockWork(i, &wg, t)
	}
	wg.Wait()
}

// 性能测试
//go test -bench=.
func BenchmarkRedisLock(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	lockKey := "myLock"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		redisSession := NewRedisSession(RedisClient)
		redisLock, err := redisSession.TryLock(lockKey, 10*time.Second)
		if err != nil {
			b.Error(err)
		}
		err = redisLock.Release()
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisLockParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	lockKey := "myLock"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			redisSession := NewRedisSession(RedisClient)
			redisLock, err := redisSession.TryLock(lockKey, 10*time.Second)
			if err != nil {
				b.Error(err)
			}
			err = redisLock.Release()
			if err != nil {
				b.Error(err)
			}
		}
	})
}

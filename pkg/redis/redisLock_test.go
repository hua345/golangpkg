package redis

import (
	"sync"
	"testing"
	"time"
)

func redisLockWork(index int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	redisSession := NewRedisSession(GetInstance())
	redisLock, err := redisSession.TryLock("myLock", 10*time.Second)
	if err != nil {
		t.Log(err)
	}
	t.Log("redis Lock Success", index)
	time.Sleep(10 * time.Millisecond)
	err = redisLock.Release()
	if err != nil {
		t.Log(err)
	}
	t.Log("released lock for session", index)
}
func TestRedisLock(t *testing.T) {
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
	lockKey := "myLock"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		redisSession := NewRedisSession(GetInstance())
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
	lockKey := "myLock"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			redisSession := NewRedisSession(GetInstance())
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

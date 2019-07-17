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

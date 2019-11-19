package redis

import (
	"sync"
	"testing"
)

func redisLeafWork(leaf *RedisLeaf, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		t.Log(leaf.NextId())
	}
}
func TestRedisLeaf(t *testing.T) {
	testLeafKey := "myLeafKey"
	redisLeaf := NewLeaf(testLeafKey)
	defer redisLeaf.Close()
	wg := sync.WaitGroup{}
	wg.Add(100)
	// create two separate sessions for lock competition
	for i := 0; i < 100; i++ {
		go redisLeafWork(redisLeaf, &wg, t)
	}
	wg.Wait()
}

// 性能测试
//go test -bench=.
func BenchmarkLeaf(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testLeafKey := "myLeafKey"
	redisLeaf := NewLeaf(testLeafKey)
	defer redisLeaf.Close()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		redisLeaf.NextId()
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisLeafParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testLeafKey := "myLeafKey"
	redisLeaf := NewLeaf(testLeafKey)
	defer redisLeaf.Close()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			redisLeaf.NextId()
		}
	})
}

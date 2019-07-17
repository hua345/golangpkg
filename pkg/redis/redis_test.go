package redis

import (
	"testing"
)

// 单元测试
func TestRedisSetGet(t *testing.T) {
	NewRedis()
	key := "name"
	value := "fangfang"
	err := RedisClient.Set(key, value, 0).Err()
	if err != nil {
		t.Error(err)
	}

	result, err := RedisClient.Get(key).Result()
	if err != nil {
		t.Error(err)
	}
	if value != result {
		t.Log("key ", key, "value ", value)
		t.Log("key ", key, "result ", result)
		t.Error("redis键值对测试失败")
	}
}

// 性能测试
//go test -bench=.
func BenchmarkRedisSet(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	key := "name"
	value := "fangfang"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		err := RedisClient.Set(key, value, 0).Err()
		if err != nil {
			b.Error(err)
		}
	}
}

// 性能测试
//go test -bench=.
func BenchmarkRedisGet(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	key := "name"
	value := "fangfang"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		result, err := RedisClient.Get(key).Result()
		if err != nil {
			b.Error(err)
		}
		if value != result {
			b.Error("redis键值对测试失败")
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisGetParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	key := "name"
	value := "fangfang"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			result, err := RedisClient.Get(key).Result()
			if err != nil {
				b.Error(err)
			}
			if value != result {
				b.Error("redis键值对测试失败")
			}
		}
	})
}

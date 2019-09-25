package redis

import "testing"

// 单元测试
func TestRedisClusterSetGet(t *testing.T) {
	NewRedisCluster()
	key := "name"
	value := "fangfang"
	err := ClusterClient.Set(key, value, 0).Err()
	if err != nil {
		t.Error(err)
	}

	result, err := ClusterClient.Get(key).Result()
	if err != nil {
		t.Error(err)
	}
	if value != result {
		t.Log("key ", key, "value ", value)
		t.Log("key ", key, "result ", result)
		t.Error("redis键值对测试失败")
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkRedisClusterGetParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedisCluster()
	b.StartTimer()
	key := "name"
	value := "fangfang"
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			result, err := ClusterClient.Get(key).Result()
			if err != nil {
				b.Error(err)
			}
			if value != result {
				b.Error("redis键值对测试失败")
			}
		}
	})
}

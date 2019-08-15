package redis

import "testing"

func TestPublish(t *testing.T) {
	NewRedis()
	pubResult := RedisClient.Publish("mychannel", "hello")
	if pubResult.Err() != nil {
		t.Error(pubResult.Err())
	}
	t.Log(pubResult.Result())
}

// 性能测试
//go test -bench=.
func BenchmarkPublish(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRedis()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		pubResult := RedisClient.Publish("mychannel", "hello")
		if pubResult.Err() != nil {
			b.Error(pubResult.Err())
		}
	}
}

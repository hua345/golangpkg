package redis

import "testing"

//# 对没有订阅者的频道发送信息
//redis> publish bad_channel "can any body hear me?"
//(integer) 0
//# 向有一个订阅者的频道发送信息
//redis> publish msg "good morning"
//(integer) 1
func TestPublish(t *testing.T) {
	// 接收到信息 message 的订阅者数量
	pubResult := GetInstance().Publish("mychannel", "hello")
	if pubResult.Err() != nil {
		t.Error(pubResult.Err())
	}
	t.Log(pubResult.Result())
}

// 性能测试
//go test -bench=.
func BenchmarkPublish(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		pubResult := GetInstance().Publish("mychannel", "hello")
		if pubResult.Err() != nil {
			b.Error(pubResult.Err())
		}
	}
}

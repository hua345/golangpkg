package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

// The Conn type is the core of the kafka-go package.
// It wraps around a raw network connection to expose a low-level API to a Kafka server.
func TestConnectionProducer(t *testing.T) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaLeaderBroker, kafkaTopic, 1)

	err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		t.Error(err)
	}
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("fang")},
	)
	if err != nil {
		t.Error(err)
	}
	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

// 性能测试
//go test -bench=.
func BenchmarkConnectionProducer(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaLeaderBroker, kafkaTopic, 1)
	defer conn.Close()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte("fang")},
		)
		if err != nil {
			b.Error(err)
		}
	}
}

package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strconv"
	"testing"
	"time"
)

// To produce messages to Kafka, a program may use the low-level Conn API,
// but the package also provides a higher level Writer type which is more appropriate to use in most cases as it provides additional features:
// Automatic retries and reconnections on errors.
// Configurable distribution of messages across available partitions.
// Synchronous or asynchronous writes of messages to Kafka.
// Asynchronous cancellation using contexts.
// kafka-go是通过过批次处理来优化吞吐量，可以通过设置批次超时时间或者异步些来减少延时

func TestBatchWriter(t *testing.T) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      kafkaBrokersList,
		Topic:        kafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("love"),
			Value: []byte("fangfang"),
		},
	)

	if err != nil {
		t.Error(err)
	}
	err = writer.Close()
	if err != nil {
		t.Error(err)
	}
}

// 性能测试
//go test -bench=.
func BenchmarkBatchWriter(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	// make a writer that produces to kafkaTopic-A, using the least-bytes distribution
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      kafkaBrokersList,
		Topic:        kafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	})
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		msg := kafka.Message{
			Key:   []byte("love"),
			Value: []byte("fangfang" + strconv.Itoa(i)),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			b.Error(err)
		}
	}
}

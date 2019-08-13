package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestWriter(t *testing.T) {
	// make a writer that produces to kafkaTopic-A, using the least-bytes distribution
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaBrokersList,
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("hello world"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
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

package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestReader(t *testing.T) {
	// make a new reader that consumes from kafkaTopic-A, kafkaPartition 0, at offset 42
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   kafkaBrokersList,
		Topic:     kafkaTopic,
		Partition: kafkaPartition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	err := reader.SetOffset(26)
	if err != nil {
		t.Error(err)
	}
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))
	}
	err = reader.Close()
	if err != nil {
		t.Error(err)
	}
}

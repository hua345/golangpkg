package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

// A Reader is another concept exposed by the kafka-go package,
// which intends to make it simpler to implement the typical use case of consuming from a single topic-partition pair.
// A Reader also automatically handles reconnections and offset management,
// and exposes an API that supports asynchronous cancellations and timeouts using Go contexts.
func TestBatchReader(t *testing.T) {
	// make a new reader that consumes from kafkaTopic-A, kafkaPartition 0, at offset 42
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   kafkaBrokersList,
		Topic:     kafkaTopic,
		Partition: 1,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	// (*Reader).SetOffset will return an error when GroupID is set
	//(*Reader).Offset will always return -1 when GroupID is set
	//(*Reader).Lag will always return -1 when GroupID is set
	//(*Reader).ReadLag will return an error when GroupID is set
	//(*Reader).Stats will return a partition of -1 when GroupID is set
	err := reader.SetOffset(2)
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

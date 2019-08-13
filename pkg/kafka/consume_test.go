package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	// to consume messages
	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaLeaderBroker, kafkaTopic, kafkaPartition)

	err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		t.Error(err)
	}
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		byteLength, err := batch.Read(b)
		if err != nil {
			break
		}
		t.Log(string(b[0:byteLength]))
	}

	batch.Close()
	conn.Close()
}

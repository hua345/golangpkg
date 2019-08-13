package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaLeaderBroker, kafkaTopic, kafkaPartition)

	err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		t.Error(err)
	}
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		t.Error(err)
	}
	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

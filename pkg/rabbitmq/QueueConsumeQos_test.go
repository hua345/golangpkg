package rabbitmq

import (
	"testing"
)

func TestQueueConsumeQos(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "hello"
	QueueConsumeQos(ch, queueName)
}

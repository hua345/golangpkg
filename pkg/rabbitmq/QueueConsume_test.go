package rabbitmq

import (
	"testing"
)

func TestQueueConsume(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	var queueName = "hello"
	QueueConsume(ch, queueName)
}

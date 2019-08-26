package rabbitmq

import (
	"testing"
)

func TestExchangeFanoutConsume(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "fangQue"
	QueueConsumeQos(ch, queueName)
}
func TestExchangeFanoutConsume2(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "fangQue"
	QueueConsumeQos(ch, queueName)
}
func TestExchangeFanoutConsume3(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "fangQue2"
	QueueConsumeQos(ch, queueName)
}

package rabbitmq

import "testing"

func TestExchangeDirectConsume(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "fangDirectQue"
	QueueConsumeQos(ch, queueName)
}
func TestExchangeDirectConsume2(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "loveDirectQue"
	QueueConsumeQos(ch, queueName)
}

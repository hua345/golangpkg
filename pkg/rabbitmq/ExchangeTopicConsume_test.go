package rabbitmq

import "testing"

func TestExchangeTopicConsume(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "fangTopicQue"
	QueueConsumeQos(ch, queueName)
}
func TestExchangeTopicConsume2(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()
	var queueName = "loveTopicQue"
	QueueConsumeQos(ch, queueName)
}

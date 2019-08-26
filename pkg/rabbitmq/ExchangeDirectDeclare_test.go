package rabbitmq

import "testing"

func TestExchangeDirectDeclare(t *testing.T) {
	var exchangeName = "fangDirect"
	ExchangeDirectDeclare(exchangeName)
}
func TestDirectQueueBind(t *testing.T) {
	var exchangeName = "fangDirect"
	var queueName = "fangDirectQue"
	var topicName = "fang"
	TopicQueueBind(exchangeName, queueName, topicName)
}
func TestDirectQueueBind2(t *testing.T) {
	var exchangeName = "fangDirect"
	var queueName = "loveDirectQue"
	var topicName = "love"
	TopicQueueBind(exchangeName, queueName, topicName)
}

package rabbitmq

import "testing"

func TestExchangeTopicDeclare(t *testing.T) {
	var exchangeName = "fangTopic"
	ExchangeTopicDeclare(exchangeName)
}
func TestTopicQueueBind(t *testing.T) {
	var exchangeName = "fangTopic"
	var queueName = "fangTopicQue"
	var topicName = "fang.*"
	TopicQueueBind(exchangeName, queueName, topicName)
}
func TestTopicQueueBind2(t *testing.T) {
	var exchangeName = "fangTopic"
	var queueName = "loveTopicQue"
	var topicName = "love.*"
	TopicQueueBind(exchangeName, queueName, topicName)
}

package rabbitmq

import "testing"

func TestExchangeFanoutDeclare(t *testing.T) {
	var exchangeName = "fang"
	ExchangeFanoutDeclare(exchangeName)
}
func TestFanoutQueueBind(t *testing.T) {
	var exchangeName = "fang"
	var queueName = "fangQue"
	FanoutQueueBind(exchangeName, queueName)
}
func TestFanoutQueueBind2(t *testing.T) {
	var exchangeName = "fang"
	var queueName = "fangQue2"
	FanoutQueueBind(exchangeName, queueName)
}

package rabbitmq

import "testing"

func TestExchangeDirectDeclare(t *testing.T) {
	var exchangeName = "fangDirect"
	ExchangeDirectDeclare(exchangeName)
}

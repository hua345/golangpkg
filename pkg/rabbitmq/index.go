package rabbitmq

import (
	"github.com/streadway/amqp"
)

const rabbitmqURL = "amqp://guest:guest@192.168.137.128:5672/"

var rabbitmqClient *amqp.Connection

func NewRabbitmq() {
	var err error
	rabbitmqClient, err = amqp.Dial(rabbitmqURL)
	if err != nil {
		panic(err)
	}
}

func CloseRabbitmq() {
	err := rabbitmqClient.Close()
	if err != nil {
		panic(err)
	}
}

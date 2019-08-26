package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func QueueConsume(ch *amqp.Channel, queueName string) {
	msgChan, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer")
		panic(err)
	}

	for msg := range msgChan {
		time.Sleep(time.Second)
		log.Printf("Received a message: %s", msg.Body)
	}
}

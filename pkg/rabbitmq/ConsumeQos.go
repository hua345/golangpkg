package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func QueueConsumeQos(ch *amqp.Channel, queueName string) {
	// With a prefetch count greater than zero, the server will deliver that many
	//messages to consumers before acknowledgments are received.
	err := ch.Qos(
		5,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		panic(err)
	}
	msgChan, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // 设置auto-ack为false
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer")
		panic(err)
	}

	index := 0
	for msg := range msgChan {
		index++
		time.Sleep(100 * time.Millisecond)
		if (index % 2) == 1 {
			err := msg.Ack(true)
			if err != nil {
				panic(err)
			}
		} else {
			err = msg.Nack(true, true)
			if err != nil {
				panic(err)
			}
		}
		log.Printf("Received a message: %s", msg.Body)
	}
}

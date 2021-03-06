package rabbitmq

import (
	"github.com/streadway/amqp"
	"strconv"
	"testing"
)

func TestQueuePublish(t *testing.T) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		t.Error(err)
	}
	defer ch.Close()

	var queueName = "hello"
	body := "Hello World!"
	err = ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		t.Error(err)
	}
}

// 性能测试
//go test -bench=.
func BenchmarkQueuePublish(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		b.Error(err)
	}
	defer ch.Close()
	b.StartTimer()
	var queueName = "hello"
	body := "Hello World!"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		err = ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body + strconv.Itoa(i)),
			})
		if err != nil {
			b.Error(err)
		}
	}
}

package rabbitmq

import (
	"github.com/streadway/amqp"
	"strconv"
	"testing"
)

// 性能测试
//go test -bench=.
func BenchmarkExchangeDirectPublish(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		b.Error(err)
	}
	defer ch.Close()
	b.StartTimer()
	var exchangeName = "fangDirect"
	var fangDirect = "fang"
	var loveDirect = "love"
	body := "Hello World!"
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			err = ch.Publish(
				exchangeName, // exchange
				fangDirect,   // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body + fangDirect + strconv.Itoa(i)),
				})
			if err != nil {
				b.Error(err)
			}
		} else {
			err = ch.Publish(
				exchangeName, // exchange
				loveDirect,   // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body + loveDirect + strconv.Itoa(i)),
				})
			if err != nil {
				b.Error(err)
			}
		}
	}
}

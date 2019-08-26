package rabbitmq

func ExchangeTopicDeclare(exchangeName string) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

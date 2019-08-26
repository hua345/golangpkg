package rabbitmq

func ExchangeDirectDeclare(exchangeName string) {
	NewRabbitmq()
	ch, err := rabbitmqClient.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

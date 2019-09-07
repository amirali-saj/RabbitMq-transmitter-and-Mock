package rabbitpublisher

import (
	errorhandling2 "Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/errorhandling"
	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel

func Connect() {
	connection, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
	channel, _ = connection.Channel()

	err := channel.ExchangeDeclare(
		"rabbit_data_transmit_exchange_topic", // name
		"topic",                               // type
		true,                                  // durable
		false,                                 // auto-deleted
		false,                                 // internal
		false,                                 // no-wait
		nil,                                   // arguments
	)
	errorhandling2.HandleError(err)
}

func Send(message []byte, routingKey string) {
	err := channel.Publish(
		"rabbit_data_transmit_exchange_topic2", // exchange name
		routingKey,                            // routing key
		false,                                 // mandatory
		false,                                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	errorhandling2.HandleError(err)
}

func Close() {
	err := channel.Close()
	errorhandling2.HandleError(err)
	err = connection.Close()
	errorhandling2.HandleError(err)
}

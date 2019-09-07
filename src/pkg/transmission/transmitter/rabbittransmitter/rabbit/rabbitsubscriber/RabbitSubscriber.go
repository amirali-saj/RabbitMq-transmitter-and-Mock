package rabbitsubscriber

import (
	errorhandling2 "Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/errorhandling"
	"Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/subscribertypes"
	"github.com/streadway/amqp"
)

type Subscriber struct {
	connection        *amqp.Connection
	qChannel          *amqp.Channel
	exChannel          *amqp.Channel

	queue             amqp.Queue
	forever           chan bool
	msg               <-chan amqp.Delivery
	ReceiverFunctions subscribertypes.ReceiveFunctionsSlice
}

func (sub *Subscriber) Connect() {

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorhandling2.HandleError(err)
	ch, err := connection.Channel()
	errorhandling2.HandleError(err)
	sub.connection = connection
	sub.qChannel = ch
	//changes
	ch2,err := connection.Channel()
	sub.exChannel = ch2
	err = sub.exChannel.ExchangeDeclare(
		"rabbit_data_transmit_exchange_topic2", // name
		"topic",                               // type
		true,                                  // durable
		false,                                 // auto-deleted
		false,                                 // internal
		false,                                 // no-wait
		nil,                                   // arguments
	)
	errorhandling2.HandleError(err)

	sub.queue, _ = ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (sub *Subscriber) Bind(routingKey string) {
	err := sub.qChannel.QueueBind(
		sub.queue.Name,                        // queue name
		routingKey,                            // routing key
		"rabbit_data_transmit_exchange_topic2", // exchange name
		false,
		nil)
	errorhandling2.HandleError(err)
}

func (sub *Subscriber) Listen() {

	msg, err := sub.qChannel.Consume(
		sub.queue.Name, // queue
		"",             // consumer
		true,           // auto ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // args
	)
	errorhandling2.HandleError(err)
	sub.msg = msg

	go func() {
		for m := range msg {
			sub.ReceiverFunctions.CallAll(m.Body)
		}
	}()

}

func (sub *Subscriber) Close() {
	err := sub.qChannel.Close()
	errorhandling2.HandleError(err)
	err = sub.connection.Close()
	errorhandling2.HandleError(err)
}


package rabbittransmitter

import (
	rabbitpublisher2 "github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/rabbit/rabbitpublisher"
	rabbitsubscriber2 "github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/rabbit/rabbitsubscriber"
	subscribertypes2 "github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/subscribertypes"
)

type RabbitTransmitter struct {
	subscribers []rabbitsubscriber2.Subscriber
}

func (r *RabbitTransmitter) Start() {
	rabbitpublisher2.Connect()
	r.subscribers = make([]rabbitsubscriber2.Subscriber, 0)
}

func (r *RabbitTransmitter) Register(functionsSlice subscribertypes2.ReceiveFunctionsSlice, label string) {
	subscriber := rabbitsubscriber2.Subscriber{}
	subscriber.Connect()
	subscriber.Bind("change." + label)
	subscriber.ReceiverFunctions = functionsSlice
	subscriber.Listen()
	r.subscribers = append(r.subscribers, subscriber)
}

func (r *RabbitTransmitter) Broadcast(bytes []byte, label string) {
	rabbitpublisher2.Send(bytes, "change."+label)
}

func (r *RabbitTransmitter) Stop() {
	rabbitpublisher2.Close()
	for i := range r.subscribers {
		r.subscribers[i].Close()
	}

}

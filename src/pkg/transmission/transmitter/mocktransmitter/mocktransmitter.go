package mocktransmitter

import (
	subscribertypes2 "github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/subscribertypes"
)

type subscriber struct {
	functions subscribertypes2.ReceiveFunctionsSlice
	label     string
}

type MockTransmitter struct {
	subscribers        []subscriber
	broadcastDataQueue []message
	active             bool
}

type message struct {
	content []byte
	label   string
	readBy  int
}

func (mock *MockTransmitter) Start() {
	mock.subscribers = make([]subscriber, 0)
	mock.broadcastDataQueue = make([]message, 0)
	mock.active = true
	go func() {
		for mock.active {
			for i := range mock.subscribers {
				iterateOverDataQueue(mock, i)
			}
		}
	}()
}

func (mock *MockTransmitter) Broadcast(data []byte, label string) {
	mock.broadcastDataQueue = append(mock.broadcastDataQueue, message{content: data, readBy: 0, label: label})
}

func (mock *MockTransmitter) Register(functionsSlice subscribertypes2.ReceiveFunctionsSlice, label string) {
	mock.subscribers = append(mock.subscribers, subscriber{functions: functionsSlice, label: label})
}

func (mock *MockTransmitter) Stop() {
	mock.active = false
}

//Returns number of receivers with given label.
func (mock *MockTransmitter) numberOfReceivers(label string) int {
	count := 0
	for _, subscriber := range mock.subscribers {
		if subscriber.label == label {
			count++
		}
	}
	return count
}

func iterateOverDataQueue(mock *MockTransmitter, receiverIndex int) {
	for j := range mock.broadcastDataQueue {
		if label := mock.subscribers[receiverIndex].label;
			label == mock.broadcastDataQueue[j].label &&
				mock.broadcastDataQueue[j].readBy < mock.numberOfReceivers(label) {
			mock.subscribers[receiverIndex].functions.CallAll(mock.broadcastDataQueue[j].content)
			mock.broadcastDataQueue[j].readBy++
		}
	}
}

package transmitter

import (
	subscribertypes2 "Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/subscribertypes"
)

type Transmitter interface {
	Start()
	Register(functionsSlice subscribertypes2.ReceiveFunctionsSlice, label string)
	Broadcast([]byte, string)
	Stop()
}

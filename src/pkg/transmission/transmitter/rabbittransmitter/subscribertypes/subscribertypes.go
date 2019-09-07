package subscribertypes

import (
	"Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/errorhandling"
	gobencoding2 "Rabbitmq-data-transmission/src/pkg/transmission/transmitter/rabbittransmitter/gobencoding"
)

type Type1 struct {
	FirstName string
	LastName  string
	Age       int
}

type Type2 struct {
	ProductId uint8
	Price     float64
}

type Type3 struct {
	FileName   string
	HashString string
	Size       int64
}

type Type4 struct {
	UserName string
	Password string
}

type ReceiveFunctionsSlice interface {
	CallAll(marshalledMessage []byte)
}


type Type1ReceiveFunctions []func(type1 Type1)
type Type2ReceiveFunctions []func(type1 Type2)
type Type3ReceiveFunctions []func(type1 Type3)
type Type4ReceiveFunctions []func(type1 Type4)

func (functions Type1ReceiveFunctions) CallAll(marshalledMessage []byte) {
	var unmarshalledMessage Type1
	err := gobencoding2.Unmarshal(marshalledMessage,&unmarshalledMessage)
	errorhandling.HandleError(err)
	for _, function := range functions {
		function(unmarshalledMessage)
	}
}

func (functions Type2ReceiveFunctions) CallAll(marshalledMessage []byte) {
	var unmarshalledMessage Type2
	err := gobencoding2.Unmarshal(marshalledMessage,&unmarshalledMessage)
	errorhandling.HandleError(err)
	for _, function := range functions {
		function(unmarshalledMessage)
	}
}
func (functions Type3ReceiveFunctions) CallAll(marshalledMessage []byte) {
	var unmarshalledMessage Type3
	err:= gobencoding2.Unmarshal(marshalledMessage,&unmarshalledMessage)
	errorhandling.HandleError(err)
	for _, function := range functions {
		function(unmarshalledMessage)
	}
}
func (functions Type4ReceiveFunctions) CallAll(marshalledMessage []byte) {
	var unmarshalledMessage Type4
	err:= gobencoding2.Unmarshal(marshalledMessage,&unmarshalledMessage)
	errorhandling.HandleError(err)
	for _, function := range functions {
		function(unmarshalledMessage)
	}
}

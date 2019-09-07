package main

import (
	"github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter"
	"github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/gobencoding"
	"github.com/amirali-saj/RabbitMq-transmitter-and-Mock/src/pkg/transmission/transmitter/rabbittransmitter/subscribertypes"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")
	receiverType1Function1 := func(data subscribertypes.Type1) {
		fmt.Printf("First name:%v\nLast name:%v\nAge:%d\n\n", data.FirstName, data.LastName, data.Age)
	}
	receiverType1Function3 := func(data subscribertypes.Type1) {
		fmt.Printf("%v %v is %d years old!\n\n", data.FirstName, data.LastName, data.Age)
	}

	receiverType2Function := func(data subscribertypes.Type2) {
		fmt.Printf("Product Id:%v\nPrice:%f\n\n", data.ProductId, data.Price)
	}
	receiverType3Function := func(data subscribertypes.Type3) {
		fmt.Printf("File name:%v\nHash:%v\nFile size:%d\n\n", data.FileName, data.HashString, data.Size)
	}
	receiverType4Function := func(data subscribertypes.Type4) {
		fmt.Printf("Username:%v\nPassword:%v\n\n", data.UserName, data.Password)
	}

	rt := rabbittransmitter.RabbitTransmitter{}
	//rt := mocktransmitter.MockTransmitter{}
	rt.Start()
	rt.Register(subscribertypes.Type1ReceiveFunctions{receiverType1Function1, receiverType1Function3}, "1")
	rt.Register(subscribertypes.Type1ReceiveFunctions{receiverType1Function3}, "1")
	rt.Register(subscribertypes.Type2ReceiveFunctions{receiverType2Function}, "2")
	rt.Register(subscribertypes.Type2ReceiveFunctions{receiverType2Function}, "2")
	rt.Register(subscribertypes.Type2ReceiveFunctions{receiverType2Function}, "2")

	rt.Register(subscribertypes.Type3ReceiveFunctions{receiverType3Function}, "3")

	rt.Register(subscribertypes.Type4ReceiveFunctions{receiverType4Function}, "4")
	rt.Register(subscribertypes.Type4ReceiveFunctions{receiverType4Function}, "4")
	rt.Register(subscribertypes.Type4ReceiveFunctions{receiverType4Function}, "4")

	rt.Broadcast(gobencoding.Marshal(subscribertypes.Type1{FirstName: "Amirali", LastName: "Sajjadi", Age: 21}), "1")
	rt.Broadcast(gobencoding.Marshal(subscribertypes.Type1{FirstName: "Test", LastName: "Person", Age: 5}), "1")

	rt.Broadcast(gobencoding.Marshal(subscribertypes.Type2{ProductId: 2, Price: 1500}), "2")
	rt.Broadcast(gobencoding.Marshal(subscribertypes.Type3{FileName: "file.txt", HashString: "fRdPxWf3==", Size: 350}), "3")
	rt.Broadcast(gobencoding.Marshal(subscribertypes.Type4{UserName: "Username1", Password: "password"}), "4")
	time.Sleep(2 * time.Second)
	rt.Stop()
}

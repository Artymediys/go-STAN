package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main() {

	sc, err := stan.Connect("test-cluster", "client-artyom", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Simple Synchronous Publisher
	err = sc.Publish("foo", []byte("Hello World"))
	if err != nil {
		fmt.Println(err)
		return
	} // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	},
		stan.DeliverAllAvailable())

	// Unsubscribe
	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close connection
	defer func(sc stan.Conn) {
		err = sc.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(sc)
}

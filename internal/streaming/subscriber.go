package streaming

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"go_STAN/internal/db"
	"log"
)

func Subscribe(sc *stan.Conn, clientID *string, subject *string) *stan.Subscription {
	sub, err := (*sc).Subscribe(*subject, func(msg *stan.Msg) {
		var receivedOrder db.Order
		log.Printf("Received a message: %s\n", string(msg.Data))

		unmarshalErr := json.Unmarshal(msg.Data, &receivedOrder)
		if unmarshalErr != nil {
			log.Fatalln("Incorrect data received -> skip")
		} else {
			//cache.Push(receivedOrder)
		}
	}, stan.DurableName("DurSub"))
	if err != nil {
		log.Fatalf("%s: %v", *clientID, err)
	}

	return &sub
}

func Unsubscribe(subscriber *stan.Subscription, clientID *string) {
	err := (*subscriber).Unsubscribe()
	if err != nil {
		log.Fatalf("%s: %v", *clientID, err)
		return
	}

	log.Printf("%s: unsubscribed!", *clientID)
}

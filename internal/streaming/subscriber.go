package streaming

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"go_STAN/internal/db"
	"log"
)

func Subscribe(dataBase *db.DataBase, sc *stan.Conn, clientID *string, subject *string) *stan.Subscription {
	sub, err := (*sc).Subscribe(*subject, func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))

		var receivedOrder db.Order
		unmarshalErr := json.Unmarshal(msg.Data, &receivedOrder)
		if unmarshalErr != nil {
			log.Println("Incorrect data received -> skip")
		} else {
			dataBase.AddOrder(receivedOrder)
		}
	}, stan.DurableName("DurSub"))
	if err != nil {
		log.Printf("%s: %v", *clientID, err)
	}

	return &sub
}

func Unsubscribe(subscriber *stan.Subscription, clientID *string) {
	err := (*subscriber).Unsubscribe()
	if err != nil {
		log.Printf("%s: %v", *clientID, err)
		return
	}

	log.Printf("%s: unsubscribed!", *clientID)
}

package streaming

import (
	"github.com/nats-io/stan.go"
	"log"
)

func Publish(sc *stan.Conn, clientID *string, subject *string, jsonData *[]byte) {
	err := (*sc).Publish(*subject, *jsonData)
	if err != nil {
		log.Printf("%s: %v", *clientID, err)
		return
	}

	log.Printf("%s: the data published!", *clientID)
}

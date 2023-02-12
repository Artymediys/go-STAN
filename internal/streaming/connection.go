package streaming

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func Connect(clusterID, clientID, natsURL *string) *stan.Conn {
	sc, err := stan.Connect(
		*clusterID, *clientID, stan.NatsURL(*natsURL),
		stan.NatsOptions(nats.ReconnectWait(5*time.Second), nats.Timeout(5*time.Second)),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("%s: connection lost, reason: %v", *clientID, reason)
		}),
	)
	if err != nil {
		log.Fatalf("%s: %v", *clientID, err)
		return &sc
	}

	log.Printf("%s: connected!", *clientID)
	return &sc
}

func Disconnect(sc *stan.Conn, clientID *string) {
	err := (*sc).Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%s: disconnected!", *clientID)
}

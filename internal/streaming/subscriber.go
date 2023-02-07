package streaming

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"go_STAN/pkg"
)

func Subscribe(sc stan.Conn, subject string) stan.Subscription {
	subscriber, err := sc.Subscribe(subject, func(msg *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data)) // "Received a message: #{string(msg.Data)}\n"
	}, stan.DeliverAllAvailable())

	pkg.PrintIfError(err)

	return subscriber
}

func Unsubscribe(subscriber stan.Subscription) {
	pkg.PrintIfError(subscriber.Unsubscribe())
}

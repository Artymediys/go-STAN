package streaming

import (
	"github.com/nats-io/stan.go"
	"go_STAN/pkg"
)

func Connect(clusterID, clientID, natsURL string) stan.Conn {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	pkg.PrintIfError(err)

	return sc
}

func Disconnect(sc stan.Conn) {
	pkg.PrintIfError(sc.Close())
}

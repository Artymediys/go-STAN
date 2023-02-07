package streaming

import (
	"github.com/nats-io/stan.go"
	"go_STAN/pkg"
)

func Publish(sc stan.Conn, subject string, jsonData string) {
	pkg.PrintIfError(sc.Publish(subject, []byte(jsonData)))
}

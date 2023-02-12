package streaming

import (
	"github.com/nats-io/stan.go"
	"go_STAN/cmd/configs"
	"go_STAN/internal/testing"
	"log"
)

var stanPublisher *stan.Conn
var stanSubscriber *stan.Conn
var subscriber *stan.Subscription

func Run() {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		clusterID = configData.STAN.ClusterID
		clientID  = configData.STAN.ClientID
		natsURL   = configData.STAN.NatsURL
		subject   = configData.STAN.Subject
	)

	stanPublisher = Connect(&clusterID, &clientID.PublisherID, &natsURL)
	if stanPublisher == nil {
		return
	}

	stanSubscriber = Connect(&clusterID, &clientID.SubscriberID, &natsURL)
	if stanSubscriber == nil {
		return
	}
	subscriber = Subscribe(stanSubscriber, &clientID.SubscriberID, &subject)

	order1, order2 := testing.GetTestOrders()
	Publish(stanPublisher, &clientID.PublisherID, &subject, &order1)
	Publish(stanPublisher, &clientID.PublisherID, &subject, &order2)
}

func Finish() {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Fatalln(err)
		return
	}

	Unsubscribe(subscriber, &configData.STAN.ClientID.SubscriberID)
	Disconnect(stanSubscriber, &configData.STAN.ClientID.SubscriberID)
	Disconnect(stanPublisher, &configData.STAN.ClientID.PublisherID)

	log.Println("STAN: finished!")
}

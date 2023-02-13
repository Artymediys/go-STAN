package streaming

import (
	"github.com/nats-io/stan.go"
	"go_STAN/cmd/configs"
	"go_STAN/internal/testing"
	"log"
)

type StanUsers struct {
	stanPublisher  *stan.Conn
	stanSubscriber *stan.Conn
	subscriber     *stan.Subscription
}

func Run(stanUsers *StanUsers) {
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

	stanUsers.stanPublisher = Connect(&clusterID, &clientID.PublisherID, &natsURL)
	if stanUsers.stanPublisher == nil {
		return
	}

	stanUsers.stanSubscriber = Connect(&clusterID, &clientID.SubscriberID, &natsURL)
	if stanUsers.stanSubscriber == nil {
		return
	}
	stanUsers.subscriber = Subscribe(stanUsers.stanSubscriber, &clientID.SubscriberID, &subject)

	order1, order2 := testing.GetTestOrders()
	Publish(stanUsers.stanPublisher, &clientID.PublisherID, &subject, &order1)
	Publish(stanUsers.stanPublisher, &clientID.PublisherID, &subject, &order2)
}

func Finish(stanUsers *StanUsers) {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Fatalln(err)
		return
	}

	Unsubscribe(stanUsers.subscriber, &configData.STAN.ClientID.SubscriberID)
	Disconnect(stanUsers.stanSubscriber, &configData.STAN.ClientID.SubscriberID)
	Disconnect(stanUsers.stanPublisher, &configData.STAN.ClientID.PublisherID)

	log.Println("STAN: finished!")
}

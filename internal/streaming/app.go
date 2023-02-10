package streaming

import (
	"go_STAN/cmd/configs"
	"log"
)

func Run() {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		clusterID = configData.ClusterID
		clientID  = configData.ClientID
		natsURL   = configData.NatsURL
		subject   = configData.Subject
	)

	stanPublisher := Connect(clusterID, clientID.PublisherID, natsURL)
	if stanPublisher == nil {
		return
	}
	//defer Disconnect(stanPublisher)

	stanSubscriber := Connect(clusterID, clientID.SubscriberID, natsURL)
	if stanSubscriber == nil {
		return
	}
	//defer Disconnect(stanSubscriber)

	Publish(stanPublisher, subject, `{"test": "bruh"}`)
	Publish(stanPublisher, subject, `{"test": "ez25"}`)

	//subscriber := Subscribe(stanSubscriber, subject)

}

func Finish() {
	//Unsubscribe(subscriber)
}

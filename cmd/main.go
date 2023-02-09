package main

import (
	"go_STAN/cmd/configs"
	"go_STAN/internal/streaming"
)

func main() {
	clusterID, clientID, natsURL := configs.GetConfigData()

	stanPublisher := streaming.Connect(clusterID, clientID.PublisherID, natsURL)
	defer streaming.Disconnect(stanPublisher)

	stanSubscriber := streaming.Connect(clusterID, clientID.SubscriberID, natsURL)
	defer streaming.Disconnect(stanSubscriber)

	streaming.Publish(stanPublisher, "orders", `{"test": "bruh"}`)
	streaming.Publish(stanPublisher, "orders", `{"test": "ez25"}`)

	subscriber := streaming.Subscribe(stanSubscriber, "orders")
	// WaitGroup
	streaming.Unsubscribe(subscriber)
}

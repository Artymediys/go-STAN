package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Connection struct {
	ClusterID string `json:"cluster_id"`
	ClientID  Client `json:"client_id"`
	NatsURL   string `json:"nats_url"`
	Subject   string `json:"subject"`
}

type Client struct {
	PublisherID  string `json:"publisher_id"`
	SubscriberID string `json:"subscriber_id"`
}

func GetConfigData() (Connection, error) {
	jsonData, err := os.ReadFile("cmd/configs/connection_config.json")
	if err != nil {
		log.Fatal(err)
		return Connection{}, err
	}

	var configData Connection
	err = json.Unmarshal(jsonData, &configData)

	return configData, err
}

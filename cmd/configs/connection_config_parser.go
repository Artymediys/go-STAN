package configs

import (
	"encoding/json"
	"go_STAN/pkg"
	"os"
)

type Connection struct {
	ClusterID string `json:"cluster_id"`
	ClientID  Client `json:"client_id"`
	NatsURL   string `json:"nats_url"`
}

type Client struct {
	PublisherID  string `json:"publisher_id"`
	SubscriberID string `json:"subscriber_id"`
}

func GetConfigData() (string, Client, string) {
	jsonData, err := os.ReadFile("cmd/configs/connection_config.json")
	pkg.PrintIfError(err)

	var configData Connection
	err = json.Unmarshal(jsonData, &configData)

	return configData.ClusterID, configData.ClientID, configData.NatsURL
}

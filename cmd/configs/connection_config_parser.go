package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Connection struct {
	STAN Stan     `json:"stan"`
	DB   DataBase `json:"db"`
}

type Stan struct {
	ClusterID string `json:"cluster_id"`
	ClientID  Client `json:"client_id"`
	NatsURL   string `json:"nats_url"`
	Subject   string `json:"subject"`
}

type Client struct {
	PublisherID  string `json:"publisher_id"`
	SubscriberID string `json:"subscriber_id"`
}

type DataBase struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	DBName     string `json:"db_name"`
	SSLMode    string `json:"ssl_mode"`
	DriverName string `json:"driver_name"`
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

package db

import (
	"database/sql"
	"fmt"
	"go_STAN/cmd/configs"
	"log"
	"sync"
)

type DataBase struct {
	mutex sync.RWMutex
	cache *Cache
	DB    *sql.DB
}

func InitDB() *DataBase {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Fatal(err)
		return &DataBase{}
	}

	connectionData := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		configData.DB.User,
		configData.DB.Password,
		configData.DB.DBName,
		configData.DB.SSLMode,
	)

	db, openErr := sql.Open(configData.DB.DriverName, connectionData)
	if openErr != nil {
		log.Fatal(openErr)
		return &DataBase{}
	}

	cache := InitCache()

	return &DataBase{
		cache: cache,
		DB:    db,
	}
}

func (db *DataBase) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func fillCache(cache *Cache) {

}

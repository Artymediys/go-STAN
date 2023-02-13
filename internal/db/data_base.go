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

func fillCache(db *sql.DB, cache *Cache) {
	rows, err := db.Query(
		"SELECT orders.id," +
			" orders.order_uid," +
			" orders.customer_id," +
			" payments.transaction," +
			" orders.locale" +
			" FROM orders LEFT OUTER JOIN payments" +
			" ON orders.payment_id = payments.id")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Fatal(err)
		}
	}(rows)

	var orders []MainInfo
	for rows.Next() {
		order := MainInfo{}
		scanErr := rows.Scan(&order.ID, &order.OrderUID, &order.CustomerID, &order.Transaction, &order.Locale)
		if scanErr != nil {
			fmt.Println(scanErr)
			continue
		}
		orders = append(orders, order)
	}
	for _, order := range orders {
		cache.Push(order)
	}
}

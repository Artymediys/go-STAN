package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"go_STAN/cmd/configs"
	"log"
	"sync"
)

type DataBase struct {
	mutex sync.RWMutex
	Cache *Cache
	DB    *sql.DB
}

func InitDB() *DataBase {
	configData, err := configs.GetConfigData()
	if err != nil {
		log.Println(err)
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
		log.Println(openErr)
		return &DataBase{}
	}

	cache := InitCache()
	fillCache(db, cache)

	return &DataBase{
		Cache: cache,
		DB:    db,
	}
}

func (db *DataBase) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Println(err)
	}

	log.Println("DataBase: disconnected!")
}

func (db *DataBase) AddOrder(newOrder Order) {
	var lastInsertID int
	var itemsIDs []int

	db.mutex.Lock()

	// Delivery addition
	deliveryErr := db.DB.QueryRow(
		"INSERT INTO deliveries"+
			" (name, phone, zip, city, address, region, email)"+
			" values ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		newOrder.Delivery.Name, newOrder.Delivery.Phone, newOrder.Delivery.Zip, newOrder.Delivery.City,
		newOrder.Delivery.Address, newOrder.Delivery.Region, newOrder.Delivery.Email).Scan(&lastInsertID)
	if deliveryErr != nil {
		log.Printf("DataBase: unable to insert data - \"delivery\": %v\n", deliveryErr)
		return
	}
	deliveryIDfk := lastInsertID

	// Payment addition
	paymentErr := db.DB.QueryRow(
		"INSERT INTO payments"+
			" (transaction, request_id, currency, provider, amount,"+
			" payment_dt, bank, delivery_cost, goods_total, custom_fee)"+
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		newOrder.Payment.Transaction, newOrder.Payment.RequestID, newOrder.Payment.Currency, newOrder.Payment.Provider,
		newOrder.Payment.Amount, newOrder.Payment.PaymentDT, newOrder.Payment.Bank, newOrder.Payment.DeliveryCost,
		newOrder.Payment.GoodsTotal, newOrder.Payment.CustomFee).Scan(&lastInsertID)
	if paymentErr != nil {
		log.Printf("DataBase: unable to insert data - \"payment\": %v\n", paymentErr)
		return
	}
	paymentIDfk := lastInsertID

	// Item addition
	for _, item := range newOrder.Items {
		itemErr := db.DB.QueryRow(
			"INSERT INTO items "+
				" (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)"+
				" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
			item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size,
			item.TotalPrice, item.NmID, item.Brand, item.Status).Scan(&lastInsertID)
		if itemErr != nil {
			log.Printf("DataBase: unable to insert data - \"item\": %v\n", itemErr)
			return
		}
		itemsIDs = append(itemsIDs, lastInsertID)
	}

	// Order addition
	orderErr := db.DB.QueryRow(
		"INSERT INTO orders"+
			" (order_uid, track_number, entry, delivery_id, payment_id, locale,"+
			" internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)"+
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
		newOrder.OrderUID, newOrder.TrackNumber, newOrder.Entry, deliveryIDfk, paymentIDfk, newOrder.Locale,
		newOrder.InternalSignature, newOrder.CustomerID, newOrder.DeliveryService, newOrder.Shardkey,
		newOrder.SmID, newOrder.DateCreated, newOrder.OofShard).Scan(&lastInsertID)
	if orderErr != nil {
		log.Printf("DataBase: unable to insert data - \"order\": %v\n", orderErr)
		return
	}
	orderIDfk := lastInsertID

	// Order-Items addition
	for _, itemID := range itemsIDs {
		_, err := db.DB.Exec(
			"INSERT INTO orders_items (order_id, item_id) values ($1, $2)",
			orderIDfk, itemID)
		if err != nil {
			log.Printf("DataBase: unable to insert data - \"order-items\": %v\n", err)
			return
		}
	}

	db.mutex.Unlock()

	db.Cache.Push(newOrder)
	log.Printf("DataBase: Order successful added to DB and to cache\n")
}

func fillCache(db *sql.DB, cache *Cache) {
	rows, err := db.Query(
		"SELECT orders.order_uid," +
			" orders.customer_id," +
			" payments.transaction," +
			" orders.locale" +
			" FROM orders LEFT OUTER JOIN payments" +
			" ON orders.payment_id = payments.id")
	if err != nil {
		log.Println(err)
		return
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Println(err)
		}
	}(rows)

	for rows.Next() {
		order := MainInfo{}

		scanErr := rows.Scan(&order.OrderUID, &order.CustomerID, &order.Transaction, &order.Locale)
		if scanErr != nil {
			fmt.Println(scanErr)
			continue
		}

		cache.Push(order)
	}
}

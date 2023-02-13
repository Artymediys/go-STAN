package db

import (
	"sync"
)

type Cache struct {
	mutex       sync.RWMutex
	orders      map[int]MainInfo
	TotalOrders int
}

func InitCache() *Cache {
	order := make(map[int]MainInfo)

	return &Cache{
		orders: order,
	}
}

func (cache *Cache) Push(order interface{}) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.TotalOrders++

	switch order.(type) {
	case Order:
		orderMainInfo := parseOrder(cache.TotalOrders, order.(Order))
		cache.orders[cache.TotalOrders] = *orderMainInfo
	case MainInfo:
		cache.orders[cache.TotalOrders] = order.(MainInfo)
	}

}

func (cache *Cache) Get(id int) (MainInfo, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	order, found := cache.orders[id]
	if !found {
		return MainInfo{}, false
	}

	return order, true
}

func parseOrder(id int, order Order) *MainInfo {
	mainInfo := &MainInfo{
		ID:          id,
		OrderUID:    order.OrderUID,
		CustomerID:  order.CustomerID,
		Transaction: order.Payment.Transaction,
		Locale:      order.Locale,
	}

	return mainInfo
}

package db

import (
	"sync"
)

type Cache struct {
	mutex       sync.RWMutex
	orders      map[uint64]Order
	TotalOrders uint64
}

func InitCache() *Cache {
	order := make(map[uint64]Order)

	return &Cache{
		orders: order,
	}
}

func (cache *Cache) Push(order Order) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.TotalOrders++
	cache.orders[cache.TotalOrders] = order
}

func (cache *Cache) Get(id uint64) (Order, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	order, found := cache.orders[id]
	if !found {
		return Order{}, false
	}

	return order, true
}

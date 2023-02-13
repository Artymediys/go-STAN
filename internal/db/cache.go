package db

import (
	"sync"
)

type Cache struct {
	mutex       sync.RWMutex
	orders      map[uint64]MainInfo
	TotalOrders uint64
}

func InitCache() *Cache {
	order := make(map[uint64]MainInfo)

	return &Cache{
		orders: order,
	}
}

func (cache *Cache) Push(order MainInfo) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.TotalOrders++
	cache.orders[cache.TotalOrders] = order
}

func (cache *Cache) Get(id uint64) (MainInfo, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	order, found := cache.orders[id]
	if !found {
		return MainInfo{}, false
	}

	return order, true
}

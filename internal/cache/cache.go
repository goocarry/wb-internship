package cache

import (
	"sync"

	"github.com/goocarry/wb-internship/internal/model"
)

// TODO: use interface

// Cache ...
type Cache struct {
	sync.RWMutex
	orders map[string]model.Order
}

// New ...
func New() *Cache {

	orders := make(map[string]model.Order)

	cache := Cache{
		orders: orders,
	}

	return &cache
}

// Set ...
func (c *Cache) Set(key string, value *model.Order) {

	c.Lock()

	defer c.Unlock()

	c.orders[key] = *value
}

// Get ...
func (c *Cache) Get(key string) (*model.Order, bool) {

	c.RLock()

	defer c.RUnlock()

	order, found := c.orders[key]

	// key not found
	if !found {
		return nil, false
	}

	return &order, true
}

// SetAll ...
func (c *Cache) SetAll(orders map[string]model.Order) {

	c.Lock()

	defer c.Unlock()

	c.orders = orders
}

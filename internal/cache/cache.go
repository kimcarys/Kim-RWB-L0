package cache

import (
	"order-service/internal/models"
	"sync"
)

type Cache struct {
	mu     sync.RWMutex
	orders map[string]*models.Order
}

func NewCache() *Cache {
	return &Cache{
		orders: make(map[string]*models.Order),
	}
}

func (c *Cache) Set(order *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *Cache) Get(orderUID string) (*models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, exists := c.orders[orderUID]
	return order, exists
}

func (c *Cache) GetAll() map[string]*models.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	// Создаем копию для безопасности
	result := make(map[string]*models.Order)
	for k, v := range c.orders {
		result[k] = v
	}
	return result
}

func (c *Cache) RestoreFromDB(db interface{}) error {
	// Заглушка - позже реализуем восстановление из БД
	// Сейчас просто логируем
	return nil
}
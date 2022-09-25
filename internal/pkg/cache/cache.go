package cache

import "sync"

type Cache struct {
	m  map[string]interface{}
	mx sync.RWMutex
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Cache) Set(key string, val any) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = val
	return nil
}

func (c *Cache) Del(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.m, key)
}

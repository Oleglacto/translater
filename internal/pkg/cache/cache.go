package cache

import (
	"sync"

	"github.com/pkg/errors"
)

var NotFound = errors.New("not found")

type Cache struct {
	m  map[string]interface{}
	mx *sync.RWMutex
}

func New() *Cache {
	m := make(map[string]interface{})
	mx := new(sync.RWMutex)
	return &Cache{m: m, mx: mx}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	if !ok {
		return nil, NotFound
	}
	return val, nil
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

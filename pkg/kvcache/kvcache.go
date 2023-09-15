package kvcache

import (
	"fmt"
	"sync"
	"time"
)

type item[V any] struct {
	Object V
	Expire int64
}

type Cache[K comparable, V any] struct {
	Expire time.Duration //default expiration time for cached item
	Data map[K]item[V] //cached items
	OnDelete func(K, V) //function that's called when cached item is deleted automatically

	mu sync.RWMutex //mutex
}

func New[K comparable, V any](ex time.Duration) *Cache[K, V] {
	return &Cache[K, V] {
		Expire: ex,
		Data: make(map[K]item[V]),
		OnDelete: func(K, V){},
	}
}

func (c *Cache[K, V]) Set(key K, val V) {
	now := time.Now().UnixNano()
	newExpire := now + int64(c.Expire)

	c.mu.Lock()

	c.Data[key] = item[V]{
		Object: val,
		Expire: newExpire,
	}

	c.mu.Unlock()
}

func (c *Cache[K, V]) KeyExists(key K) bool {
	c.mu.RLock()

	if _, ok := c.Data[key]; ok {
		c.mu.RUnlock()
		return true
	}

	c.mu.RUnlock()
	return false
}

func (c *Cache[K, V]) OnDeleteFunc(f func(K, V)) {
	c.OnDelete = f
}

func (c *Cache[K, V]) Add(key K, val V) error {
	if c.KeyExists(key) {
		return fmt.Errorf("kvcache: Data already exists with given key %T", key)
	}
	
	c.Set(key, val)

	return nil
}

func (c *Cache[K, V]) Update(key K, val V) error {
	if !c.KeyExists(key) {
		return fmt.Errorf("kvcache: Data doesn't exists with given key %T", key)
	}

	c.Set(key, val)

	return nil
}

func (c *Cache[K, V]) Get(key K) V {
	c.mu.RLock()

	data := c.Data[key].Object

	c.mu.RUnlock()
	return data
}

func (c *Cache[K, V]) Delete(key K) {
	if c.KeyExists(key) {
		delete(c.Data, key)
	}
}

func (c *Cache[K, V]) IsExpired(key K) bool {
	now := time.Now().UnixNano()
	
	c.mu.RLock()
	if val, ok := c.Data[key]; ok {
		c.mu.RUnlock()
		if val.Expire > 0 && now > val.Expire {
			return true
		}else{
			return false
		}
	}

	c.mu.RUnlock()
	return false
}

func (c *Cache[K, V]) DeleteExpired() {
	for k,v := range c.Data {
		if c.IsExpired(k) {
			c.OnDelete(k, v.Object)
			delete(c.Data, k)
		}
	}
}

func (c *Cache[K, V]) Flush() {
	for k,v := range c.Data {
		c.OnDelete(k, v.Object)
		delete(c.Data, k)
	}
}
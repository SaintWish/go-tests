package kvbyte

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	Object []byte
	Expire int64
}

type Cache[K comparable] struct {
	Expire time.Duration //default expiration time for cached item
	Data map[K]item //cached items
	OnDelete func(K, []byte) //function that's called when cached item is deleted automatically

	mu sync.RWMutex //mutex
}

func New[K comparable](ex time.Duration) *Cache[K] {
	return &Cache[K] {
		Expire: ex,
		Data: make(map[K]item),
		OnDelete: func(K, []byte){},
	}
}

func (c *Cache[K]) Set(key K, val []byte) {
	now := time.Now().UnixNano()
	newExpire := now + int64(c.Expire)

	c.mu.Lock()

	c.Data[key] = item{
		Object: val,
		Expire: newExpire,
	}

	c.mu.Unlock()
}

func (c *Cache[K]) KeyExists(key K) bool {
	c.mu.RLock()

	if _, ok := c.Data[key]; ok {
		c.mu.RUnlock()
		return true
	}

	c.mu.RUnlock()
	return false
}

func (c *Cache[K]) OnDeleteFunc(f func(K, []byte)) {
	c.OnDelete = f
}

func (c *Cache[K]) Add(key K, val []byte) error {
	if c.KeyExists(key) {
		return fmt.Errorf("kvcache: Data already exists with given key %T", key)
	}
	
	c.Set(key, val)

	return nil
}

func (c *Cache[K]) Update(key K, val []byte) error {
	if !c.KeyExists(key) {
		return fmt.Errorf("kvcache: Data doesn't exists with given key %T", key)
	}

	c.Set(key, val)

	return nil
}

func (c *Cache[K]) Get(key K) []byte {
	//c.DeleteExpired()

	c.mu.RLock()

	data := c.Data[key].Object

	c.mu.RUnlock()
	return data
}

func (c *Cache[K]) Delete(key K) {
	if c.KeyExists(key) {
		delete(c.Data, key)
	}
}

func (c *Cache[K]) IsExpired(key K) bool {
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

func (c *Cache[K]) DeleteExpired() {
	for k,v := range c.Data {
		if c.IsExpired(k) {
			c.OnDelete(k, v.Object)
			delete(c.Data, k)
		}
	}
}

func (c *Cache[K]) Flush() {
	for k,v := range c.Data {
		c.OnDelete(k, v.Object)
		delete(c.Data, k)
	}
}
package kvsharded

import (
	"time"
	
	"github.com/dolthub/maphash"
)

const defaultShardCount = 32

type Cache[K comparable, V any] struct {
	shards []*shardMap[K, V]
	OnDelete func(K, V) //function that's called when cached item is deleted automatically
}

func New[K comparable, V any](ex time.Duration) *Cache[K, V] {
	cache := Cache[K, V] {}
	cache.shards = make([]*shardMap[K, V], defaultShardCount)
	cache.OnDelete = func(K, V){}

	for i := 0; i < defaultShardCount; i++ {
		cache.shards[i] = newShardMap[K, V](ex)
	}

	return &cache
}

func (c *Cache[K, V]) getShardIndex(key K) uint64 {
	h := maphash.NewHasher[K]()
	sum := h.Hash(key)

	return sum % defaultShardCount
}

func (c *Cache[K, V]) Set(key K, val V) {
	index := c.getShardIndex(key)
	c.shards[index].set(key, val)
}

func (c *Cache[K, V]) Get(key K) V {
	index := c.getShardIndex(key)
	return c.shards[index].get(key)
}

/*
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
*/
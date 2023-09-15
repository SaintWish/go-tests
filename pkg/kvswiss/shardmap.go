package kvswiss

import (
	"time"
	"sync"

	"github.com/dolthub/swiss"
)

type item[V any] struct {
	Object V
	Expire int64
}

//used internally
type shardMap[K comparable, V any] struct {
	Map *swiss.Map[K, item[V]]
	DefaultExpire int64
	sync.RWMutex //mutex
}

func newShardMap[K comparable, V any](defaultExpire time.Duration, size uint32) *shardMap[K, V] {
	return &shardMap[K, V] {
		Map: swiss.NewMap[K, item[V]]( size/defaultShardCount ), // 
		DefaultExpire: int64(defaultExpire),
	}
}

func (m *shardMap[K, V]) set(key K, val V) {
	now := time.Now().UnixNano()
	newExpire := now + m.DefaultExpire
	itm := item[V]{
		Object: val,
		Expire: newExpire,
	}

	m.Lock()

	m.Map.Put(key, itm)

	m.Unlock()
}

func (m *shardMap[K, V]) get(key K) (V) {
	m.RLock()

	val,_ := m.Map.Get(key)

	m.RUnlock()

	return val.Object
}
package kvsharded

import (
	"time"
	"sync"
)

type item[V any] struct {
	Object V
	Expire int64
}

//used internally
type shardMap[K comparable, V any] struct {
	Map map[K]item[V]
	DefaultExpire int64
	sync.RWMutex //mutex
}

func newShardMap[K comparable, V any](defaultExpire time.Duration) *shardMap[K, V] {
	return &shardMap[K, V] {
		Map: make(map[K]item[V]),
		DefaultExpire: int64(defaultExpire),
	}
}

func (m *shardMap[K, V]) set(key K, val V) {
	now := time.Now().UnixNano()
	newExpire := now + m.DefaultExpire

	m.Lock()

	m.Map[key] = item[V]{
		Object: val,
		Expire: newExpire,
	}

	m.Unlock()
}

func (m *shardMap[K, V]) get(key K) (val V) {
	m.RLock()

	val = m.Map[key].Object

	m.RUnlock()

	return val
}
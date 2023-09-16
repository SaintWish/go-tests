package kvswiss

import (
	"time"
	"sync"

	"github.com/saintwish/go-tests/pkg/kvswiss/swiss"
	//"github.com/dolthub/swiss"
)

type item[V any] struct {
	Object V
	Expire int64
}

//used internally
type shard[K comparable, V any] struct {
	Map *swiss.Map[K, item[V]]
	DefaultExpire int64
	sync.RWMutex //mutex
}

func newShard[K comparable, V any](defaultExpire time.Duration, size uint32) *shard[K, V] {
	return &shard[K, V] {
		Map: swiss.NewMap[K, item[V]]( size/defaultShardCount ), // 
		DefaultExpire: int64(defaultExpire),
	}
}

func (m *shard[K, V]) set(key K, val V) {
	now := time.Now().UnixNano()
	newExpire := now + m.DefaultExpire
	itm := item[V]{
		Object: val,
		Expire: newExpire,
	}

	m.Lock()

	m.Map.Set(key, itm)

	m.Unlock()
}

func (m *shard[K, V]) get(key K) (V) {
	m.RLock()

	val := m.Map.Get(key)

	m.RUnlock()

	return val.Object
}
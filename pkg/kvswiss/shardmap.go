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
	Expiration time.Duration
	sync.RWMutex //mutex
}

func newShard[K comparable, V any](ex time.Duration, size uint64, count uint64) *shard[K, V] {
	return &shard[K, V] {
		Map: swiss.NewMap[K, item[V]]( uint32(size/count) ),
		Expiration: ex,
	}
}

func (m *shard[K, V]) getRaw(key K) (item[V]) {
	m.RLock()

	val := m.Map.Get(key)

	m.RUnlock()

	return val
}

func (m *shard[K, V]) getHasRaw(key K) (item[V], bool) {
	m.RLock()

	val, ok := m.Map.GetHas(key)

	m.RUnlock()

	return val, ok
}

func (m *shard[K, V]) hasRaw(key K) bool {
	m.RLock()

	ok := m.Map.Has(key)

	m.RUnlock()

	return ok
}

func (m *shard[K, V]) set(key K, val V) {
	expire := time.Now().Add(m.Expiration)

	itm := item[V]{
		Object: val,
		Expire: expire.UnixNano(),
	}

	m.Lock()

	m.Map.Set(key, itm)

	m.Unlock()
}

func (m *shard[K, V]) get(key K) V {
	now := time.Now().UnixNano()

	if v, ok := m.getHasRaw(key); ok && v.Expire == 0 || now < v.Expire {
		m.renew(key)
		return v.Object
	}

	var res V
	return res
}

func (m *shard[K, V]) getHas(key K) (V, bool) {
	now := time.Now().UnixNano()

	if v, ok := m.getHasRaw(key); ok && v.Expire == 0 || now < v.Expire {
		m.renew(key)
		return v.Object, true
	}

	var res V
	return res, false
}

func (m *shard[K, V]) has(key K) bool {
	v, ok := m.getHasRaw(key)
	now := time.Now().UnixNano()

	if ok && v.Expire == 0 || now < v.Expire {
		return true
	}

	return false
}

func (m *shard[K, V]) delete(key K) bool {
	m.Lock()

	ok := m.Map.Delete(key)

	m.Unlock()

	return ok
}

func (m *shard[K, V]) isExpired(key K) bool {
	now := time.Now().UnixNano()

	m.RLock()

	if v, ok := m.Map.GetHas(key); ok {
		m.RUnlock()
		if v.Expire > 0 && now > v.Expire {
			return true
		}else{
			return false
		}
	}

	m.RUnlock()
	return false
}

func (m *shard[K, V]) evictItem(key K, cb func(K,V)) bool {
	var res bool
	var val V
	now := time.Now().UnixNano()

	m.Lock()

	if v, ok := m.Map.GetHas(key); ok {
		if v.Expire > 0 && now > v.Expire {
			val = v.Object
			res = m.Map.Delete(key)
		}
	}else{
		m.Unlock()
		return false
	}

	m.Unlock()

	cb(key, val)
	return res
}

func (m *shard[K, V]) evictExpired(cb func(K,V)) {
	now := time.Now().UnixNano()

	m.Lock()

	m.Map.Iter(func (key K, val item[V]) (stop bool) {
		if val.Expire > 0 && now > val.Expire {
			cb(key, val.Object)
			m.Map.Delete(key)
		}
		
		if stop {
			m.Unlock()
			return
		}

		return
	})

	m.Unlock()
}

func (m *shard[K, V]) renew(key K) {
	expire := time.Now().Add(m.Expiration)

	m.Lock()
	
	if v, ok := m.Map.GetHas(key); ok {
		v.Expire = expire.UnixNano()
	}

	m.Unlock()
}
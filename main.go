package main

import (
	"sync"
	"strconv"

	"github.com/kelindar/binary"

	"github.com/SaintWish/go-tests/pkg/kvcache"
	"github.com/SaintWish/go-tests/pkg/kvbyte"
	"github.com/SaintWish/go-tests/pkg/kvsharded"
	"github.com/alphadose/haxmap"
	"github.com/lrita/cmap"
	ccmap "github.com/orcaman/concurrent-map/v2"
	"github.com/cornelk/hashmap"
	"github.com/dolthub/swiss"
	"github.com/mhmtszr/concurrent-swiss-map"
	"github.com/puzpuzpuz/xsync/v2"
	
)

const maxEntries = 500

type test struct {}

type blank struct {
	name test
}

func Map_RW() {
	m := map[int]blank{}
	for i := 1; i <= maxEntries; i++ {
		m[i] = blank{test{}}
	}

	for i := 1; i <= maxEntries; i++ {
		_ = m[i]
	}
}

func SyncMap_RW() {
	var m sync.Map
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func KVCache_RW() {
	c := kvcache.New[int, blank](0)
	for i := 1; i <= maxEntries; i++ {
		c.Set(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
	}
}

func KVByte_RW() {
	c := kvbyte.New[int](0)
	for i := 1; i <= maxEntries; i++ {
		enc, _ := binary.Marshal(blank{test{}})
		c.Set(i, enc)
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
	}
}

func KVSharded_RW() {
	c := kvsharded.New[int, blank](0)
	for i := 1; i <= maxEntries; i++ {
		c.Set(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
	}
}

func HexMap_RW() {
	m := haxmap.New[int, blank]()
	for i := 1; i <= maxEntries; i++ {
		m.Set(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func CMap_RW() {
	var m cmap.Cmap
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func CCMap_RW() {
	m := ccmap.New[blank]()
	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Set(stri, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Get(stri)
	}
}

func HashMap_RW() {
	m := hashmap.New[int, blank]()
	for i := 1; i <= maxEntries; i++ {
		m.Set(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func Swiss_RW() {
	m := swiss.NewMap[int, blank](100)
	for i := 1; i <= maxEntries; i++ {
		m.Put(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func CSSwiss_RW() {
	m := csmap.Create[int, blank](
		csmap.WithSize[int, blank](100),
	)
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func Xsync_RW() {
	m := xsync.NewMap()
	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Store(stri, blank{test{}})
	}

	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Load(stri)
	}
}
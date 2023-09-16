package main

import (
	"sync"
	"strconv"

	"github.com/saintwish/go-tests/pkg/kvcache"
	"github.com/saintwish/go-tests/pkg/kvsharded"
	"github.com/saintwish/go-tests/pkg/kvswiss"
	"github.com/alphadose/haxmap"
	"github.com/lrita/cmap"
	ccmap "github.com/orcaman/concurrent-map/v2"
	"github.com/cornelk/hashmap"
	"github.com/dolthub/swiss"
	"github.com/mhmtszr/concurrent-swiss-map"
	"github.com/puzpuzpuz/xsync/v2"
	
)

const maxEntries = 500

type blank2 struct {}

type blank struct {
	test int
	test2 blank2
}

func SyncMap_RW() {
	var m sync.Map
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func KVCache_RW() {
	c := kvcache.New[int, blank](0)
	for i := 1; i <= maxEntries; i++ {
		c.Set(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
	}
}

func KVSharded_RW() {
	c := kvsharded.New[int, blank](0)
	for i := 1; i <= maxEntries; i++ {
		c.Set(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
		//fmt.Println(res)
	}
}

func KVSwiss_RW() {
	c := kvswiss.New[int, blank](0, 1000)
	for i := 1; i <= maxEntries; i++ {
		c.Set(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		c.Get(i)
		//fmt.Println(res)
	}
}

func HexMap_RW() {
	m := haxmap.New[int, blank]()
	for i := 1; i <= maxEntries; i++ {
		m.Set(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func CMap_RW() {
	var m cmap.Cmap
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func CCMap_RW() {
	m := ccmap.New[blank]()
	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Set(stri, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Get(stri)
	}
}

func HashMap_RW() {
	m := hashmap.New[int, blank]()
	for i := 1; i <= maxEntries; i++ {
		m.Set(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func Swiss_RW() {
	m := swiss.NewMap[int, blank](1000)
	for i := 1; i <= maxEntries; i++ {
		m.Put(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Get(i)
	}
}

func CSSwiss_RW() {
	m := csmap.Create[int, blank](
		csmap.WithShardCount[int, blank](16),
		csmap.WithSize[int, blank](1000),
	)
	for i := 1; i <= maxEntries; i++ {
		m.Store(i, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		m.Load(i)
	}
}

func Xsync_RW() {
	m := xsync.NewMap()
	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Store(stri, blank{test: i, test2: blank2{}})
	}

	for i := 1; i <= maxEntries; i++ {
		stri := strconv.Itoa(i)
		m.Load(stri)
	}
}
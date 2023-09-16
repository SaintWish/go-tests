package main

import (
    "testing"

	"sync"

	"github.com/saintwish/go-tests/pkg/kvcache"
	"github.com/saintwish/go-tests/pkg/kvsharded"
	"github.com/saintwish/go-tests/pkg/kvswiss"
	"github.com/alphadose/haxmap"
	"github.com/cornelk/hashmap"
	"github.com/mhmtszr/concurrent-swiss-map"
)

func Benchmark_SyncMap(b *testing.B) {
	var m sync.Map
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Store(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Load(e)
		}
	}
}

func Benchmark_SyncMap_Parallel(b *testing.B) {
	var m sync.Map
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Store(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Load(e)
			}
		}
	})
}

func Benchmark_KVCache(b *testing.B) {
	m := kvcache.New[int, blank](0)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Set(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Get(e)
		}
	}
}

func Benchmark_KVCache_Parallel(b *testing.B) {
	m := kvcache.New[int, blank](0)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Set(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Get(e)
			}
		}
	})
}

func Benchmark_KVSharded(b *testing.B) {
	m := kvsharded.New[int, blank](0)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Set(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Get(e)
		}
	}
}

func Benchmark_KVSharded_Parallel(b *testing.B) {
	m := kvsharded.New[int, blank](0)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Set(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Get(e)
			}
		}
	})
}

func Benchmark_KVSwiss_SZ1024(b *testing.B) {
	m := kvswiss.New[int, blank](0, 1024)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Set(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Get(e)
		}
	}
}

func Benchmark_KVSwiss_Parallel(b *testing.B) {
	m := kvswiss.New[int, blank](0, 1024)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Set(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Get(e)
			}
		}
	})
}

func Benchmark_HexMap(b *testing.B) {
	m := haxmap.New[int, blank]()
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Set(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Get(e)
		}
	}
}

func Benchmark_HexMap_Parallel(b *testing.B) {
	m := haxmap.New[int, blank]()
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Set(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Get(e)
			}
		}
	})
}

func Benchmark_HashMap(b *testing.B) {
	m := hashmap.New[int, blank]()
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Set(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Get(e)
		}
	}
}

func Benchmark_HashMap_Parallel(b *testing.B) {
	m := hashmap.New[int, blank]()
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Set(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Get(e)
			}
		}
	})
}

func Benchmark_CSSwissMap_SZ1024(b *testing.B) {
	m := csmap.Create[int, blank](
		csmap.WithShardCount[int, blank](32),
		csmap.WithSize[int, blank](1024),
	)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := 1; e <= maxEntries; e++ {
			m.Store(e, s)
		}

		for e := 1; e <= maxEntries; e++ {
			m.Load(e)
		}
	}
}

func Benchmark_CSSwiss_Parallel(b *testing.B) {
	m := csmap.Create[int, blank](
		csmap.WithShardCount[int, blank](32),
		csmap.WithSize[int, blank](1024),
	)
	s := blank{test: 1337, test2: blank2{}}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for e := 1; e <= maxEntries; e++ {
				m.Store(e, s)
			}
	
			for e := 1; e <= maxEntries; e++ {
				m.Load(e)
			}
		}
	})
}
package main

import (
    "testing"
    "runtime"
)

func Benchmark_Map(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		Map_RW()
	}
}

func Benchmark_SyncMap(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		SyncMap_RW()
	}
}

func Benchmark_KVCache(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		KVCache_RW()
	}
}

func Benchmark_KVByte_Serialize(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		KVByte_RW()
	}
}

func Benchmark_KVSharded(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		KVSharded_RW()
	}
}

func Benchmark_KVSwiss_SZ1000(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		KVSwiss_RW()
	}
}

func Benchmark_HexMap(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		HexMap_RW()
	}
}

func Benchmark_CMap(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		CMap_RW()
	}
}

func Benchmark_CCMap_StringKeys(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		CCMap_RW()
	}
}

func Benchmark_HashMap(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		HashMap_RW()
	}
}

func Benchmark_SwissMap_SZ1000(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		Swiss_RW()
	}
}

func Benchmark_CSSwissMap_SZ1000(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		CSSwiss_RW()
	}
}

func Benchmark_XSyncMap_StringKeys(b *testing.B) {
	runtime.GC()

	for i := 0; i < b.N; i++ {
		Xsync_RW()
	}
}
# Non-Concurrent RW Map Tests
```
goos: windows
goarch: amd64
pkg: github.com/SaintWish/go-tests
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
Benchmark_Map-8                            21991             54837 ns/op           60523 B/op         35 allocs/op
Benchmark_SyncMap-8                         6012            197297 ns/op           90282 B/op       2265 allocs/op
Benchmark_KVCache-8                        15302             79344 ns/op           82091 B/op         41 allocs/op
Benchmark_KVByte_Serialize-8                6014            196354 ns/op          163105 B/op       1516 allocs/op
Benchmark_KVSharded-8                      10000            104723 ns/op           71087 B/op        139 allocs/op
Benchmark_KVSwiss_SZ1000-8                 14301             84385 ns/op           46512 B/op         67 allocs/op
Benchmark_HexMap-8                         10000            111364 ns/op           48816 B/op       1019 allocs/op
Benchmark_CMap-8                            8602            132983 ns/op           79627 B/op        838 allocs/op
Benchmark_CCMap_StringKeys-8                9396            123813 ns/op           71488 B/op        968 allocs/op
Benchmark_HashMap-8                         3877            296276 ns/op           48994 B/op       1027 allocs/op
Benchmark_SwissMap_SZ1000-8                50005             23930 ns/op           29904 B/op          3 allocs/op
Benchmark_CSSwissMap_SZ1000-8              17533             68715 ns/op           36048 B/op         67 allocs/op
Benchmark_XSyncMap_StringKeys-8             7077            175017 ns/op           64824 B/op       2395 allocs/op
```
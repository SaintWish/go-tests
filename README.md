# RW Map Tests with and without parallel
``go test -bench . -benchmem``

```
goos: windows
goarch: amd64
pkg: github.com/SaintWish/go-tests
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
Benchmark_SyncMap-8                        15607             74897 ns/op           17964 B/op       1245 allocs/op
Benchmark_SyncMap_Parallel-8               49034             22839 ns/op           17962 B/op       1245 allocs/op
Benchmark_KVCache-8                        34791             33096 ns/op               2 B/op          0 allocs/op
Benchmark_KVCache_Parallel-8                4298            295938 ns/op              19 B/op          0 allocs/op
Benchmark_KVSharded-8                      16153             74316 ns/op             163 B/op          0 allocs/op
Benchmark_KVSharded_Parallel-8             24807             47527 ns/op             105 B/op          0 allocs/op
Benchmark_KVSwiss_SZ1024-8                 26041             46283 ns/op               0 B/op          0 allocs/op
Benchmark_KVSwiss_Parallel-8               16114             72674 ns/op               0 B/op          0 allocs/op
Benchmark_HexMap-8                         33799             35780 ns/op            8001 B/op        500 allocs/op
Benchmark_HexMap_Parallel-8                95028             11006 ns/op            8000 B/op        500 allocs/op
Benchmark_CMap-8                           15470             78254 ns/op            9964 B/op        745 allocs/op
Benchmark_CCMap_StringKeys-8               15732             76208 ns/op            2570 B/op        802 allocs/op
Benchmark_HashMap-8                        35866             33625 ns/op            8001 B/op        500 allocs/op
Benchmark_CSSwissMap_SZ1024-8              22468             52250 ns/op               0 B/op          0 allocs/op
Benchmark_CSSwiss_Parallel-8               22322             52459 ns/op               0 B/op          0 allocs/op
```
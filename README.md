# RW Map Tests with and without parallel
``go test -bench . -benchmem``

```
goos: windows
goarch: amd64
pkg: github.com/SaintWish/go-tests
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
Benchmark_KV1_SZ2048-8                     14672             81210 ns/op               0 B/op          0 allocs/op
Benchmark_KV1_Parallel-8                    9222            126015 ns/op               0 B/op          0 allocs/op
Benchmark_KV2_SZ2048-8                     27018             44900 ns/op               0 B/op          0 allocs/op
Benchmark_KV2_Parallel-8                   13981             84153 ns/op               0 B/op          0 allocs/op
Benchmark_KVCache-8                        34081             35274 ns/op               2 B/op          0 allocs/op
Benchmark_KVCache_Parallel-8                3760            289953 ns/op              22 B/op          0 allocs/op
Benchmark_KVSharded-8                      15558             75821 ns/op             168 B/op          0 allocs/op
Benchmark_KVSharded_Parallel-8             24962             47350 ns/op             105 B/op          0 allocs/op
Benchmark_HexMap-8                         33608             36057 ns/op            8001 B/op        500 allocs/op
Benchmark_HexMap_Parallel-8                96646             11013 ns/op            8000 B/op        500 allocs/op
Benchmark_HashMap-8                        35296             34464 ns/op            8001 B/op        500 allocs/op
Benchmark_HashMap_Parallel-8              110504             10276 ns/op            8001 B/op        500 allocs/op
Benchmark_CSSwissMap_SZ2048-8              26170             46506 ns/op               0 B/op          0 allocs/op
Benchmark_CSSwiss_Parallel-8                9256            109072 ns/op               0 B/op          0 allocs/op
Benchmark_CMap-8                           13874             80982 ns/op           11920 B/op        990 allocs/op
Benchmark_CMap_Parallel-8                  42355             28221 ns/op           11920 B/op        990 allocs/op
Benchmark_SyncMap-8                        14930             78821 ns/op           17964 B/op       1245 allocs/op
Benchmark_SyncMap_Parallel-8               47764             23013 ns/op           17962 B/op       1245 allocs/op
```
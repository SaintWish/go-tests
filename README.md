# Non-Concurrent Map Tests
```
goos: windows
goarch: amd64
pkg: github.com/SaintWish/go-tests
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
Benchmark_Map-8                            26902             44570 ns/op           23438 B/op         31 allocs/op
Benchmark_SyncMap-8                         6336            181773 ns/op           82285 B/op       1765 allocs/op
Benchmark_KVCache-8                        17674             67893 ns/op           42395 B/op         37 allocs/op
Benchmark_KVByte_Serialize-8                6464            170248 ns/op          155114 B/op       1016 allocs/op
Benchmark_KVSharded-8                      13428             89317 ns/op           32958 B/op        177 allocs/op
Benchmark_HexMap-8                         13161             91308 ns/op           40800 B/op        518 allocs/op
Benchmark_CMap-8                            8875            119440 ns/op           76217 B/op        340 allocs/op
Benchmark_CCMap_StringKeys-8                9499            110456 ns/op           38437 B/op        968 allocs/op
Benchmark_HashMap-8                         4010            290086 ns/op           40881 B/op        521 allocs/op
Benchmark_SwissMap-8                       28484             40583 ns/op           19792 B/op          9 allocs/op
Benchmark_CSSwissMap-8                     14764             81271 ns/op           16080 B/op        175 allocs/op
Benchmark_XSyncMap_StringKeys-8             7066            161106 ns/op           56819 B/op       1895 allocs/op
```
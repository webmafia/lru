# LRU cache
A fast, generic and iterable LRU cache written in Go. Comes with a thread-safe variant.

## Documentation
See: [https://pkg.go.dev/github.com/webmafia/lru](https://pkg.go.dev/github.com/webmafia/lru)

## Benchmark
```
goos: darwin
goarch: arm64
pkg: github.com/webmafia/lru
cpu: Apple M1 Pro
Benchmark/cap_008-10      	95708524	        12.43 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_016-10      	74377090	        16.64 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_032-10      	45249507	        25.84 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_064-10      	26544806	        46.38 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_128-10      	13576695	        88.19 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_256-10      	 7849098	       152.8 ns/op	       0 B/op	       0 allocs/op
Benchmark/cap_512-10      	 4349156	       274.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_008-10         	53933190	        21.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_016-10         	54433185	        22.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_032-10         	33993650	        33.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_064-10         	22557008	        52.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_128-10         	12623346	        95.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_256-10         	 7475310	       160.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkThreadsafe/cap_512-10         	 4271628	       281.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/webmafia/lru	19.095s
```
Benchmarking map allocations with values and pointers:

```
goos: linux
goarch: amd64
pkg: github.com/schmichael/play/mapalloc
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkMapAllocValue-8     	117695410	        10.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapAllocPointer-8   	256055787	         4.917 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/schmichael/play/mapalloc	3.928s
```

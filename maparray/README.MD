## Benchmark Run Result:

```
$ go test ./maparray/ -bench=. -run=^a
goos: linux
goarch: amd64
pkg: github.com/nurali/go-play/maparray
BenchmarkMapAdd-4         	35346738	        29.5 ns/op
BenchmarkArrAdd-4         	11327191	       101 ns/op
BenchmarkMapGet-4         	86002251	        12.5 ns/op
BenchmarkArrGet-4         	 3483820	       310 ns/op
BenchmarkMapRemove-4      	100000000	        10.1 ns/op
BenchmarkArrRemove-4      	 3895038	       391 ns/op
BenchmarkMapNextSlot1-4   	   79940	     15870 ns/op
BenchmarkArrNextSlot1-4   	 2835604	       454 ns/op
BenchmarkMapNextSlot2-4   	   54199	     25595 ns/op
BenchmarkArrNextSlot2-4   	385482253	         2.99 ns/op
PASS
ok  	github.com/nurali/go-play/maparray	13.947s
```

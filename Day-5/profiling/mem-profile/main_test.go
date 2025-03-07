package main

import "testing"

func BenchmarkAnalyzeText(b *testing.B) {
	fileName := "moby.txt"
	for i := 0; i < b.N; i++ {
		_, err := AnalyzeText(fileName)
		if err != nil {
			b.Fatalf("Error analyzing text: %v", err)
		}
	}
}

// the below command is to run benchmark for a specific number of time
//go test -run none -bench . -benchtime 3s -benchmem -v
//BenchmarkAnalyzeText-8: This is the name of the benchmark test. The -8 part means that the benchmark was run on 8 threads to parallelize and speed up the execution.
//235: This is the number of iterations that were performed while benchmarking the function. The benchmark function is executed b.N times, where b.N increases in each iteration until the benchmark function lasts long enough to be measured accurately.
//15517327 ns/op: This means that each operation in the benchmark took an average of approximately 15.517327 milliseconds to execute. This is a measurement of time consumed per operation (ns/op stands for "nanoseconds per operation").
//27750673 B/op: This means that each operation in the benchmark caused about 27,750,673 bytes (or approximately 27MB) of allocations in memory. (B/op stands for "bytes per operation"). This can help in identifying memory-intensive operations.
//710 allocs/op: This means that every execution of the benchmarked function resulted in 710 allocations from the heap (allocs/op stands for "heap allocations per operation"). This gives an insight into how much work the garbage collector will need to do as a result of running this function.

func BenchmarkOptimizedAnalyzeText(b *testing.B) {
	fileName := "moby.txt"
	for i := 0; i < b.N; i++ {
		_, err := OptimizedAnalyzeText(fileName)
		if err != nil {
			b.Fatalf("Error analyzing text: %v", err)
		}
	}
}

//below command generates memory profile
//go test -run none -bench . -benchtime 3s -benchmem -v -memprofile p.out
// go tool pprof p.out
// list AnalyzeText
// top 5
// weblist AnalyzeText // to see ui version of the report

//go tool pprof -http=:8080 p.out

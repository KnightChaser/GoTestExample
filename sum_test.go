package calc

import "testing"

// Functions to be benchmarked should be starting with "Benchmark"
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(100, 200)
	}
}

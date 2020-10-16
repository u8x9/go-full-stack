package fib

import "testing"

func benchmark(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func Benchmark1(b *testing.B) {
	benchmark(b, 1)
}
func Benchmark2(b *testing.B) {
	benchmark(b, 2)
}
func Benchmark3(b *testing.B) {
	benchmark(b, 3)
}
func Benchmark5(b *testing.B) {
	benchmark(b, 5)
}

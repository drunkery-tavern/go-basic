package tests

import "testing"

/*func benchmarkFibonacci(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

func BenchmarkFibonacci1(b *testing.B) {
	benchmarkFibonacci(b,1)
}
func BenchmarkFibonacci2(b *testing.B) {
	benchmarkFibonacci(b,2)
}
func BenchmarkFibonacci3(b *testing.B) {
	benchmarkFibonacci(b,3)
}
func BenchmarkFibonacci10(b *testing.B) {
	benchmarkFibonacci(b,10)
}
func BenchmarkFibonacci20(b *testing.B) {
	benchmarkFibonacci(b,20)
}
func BenchmarkFibonacci40(b *testing.B) {
	benchmarkFibonacci(b,40)
}*/

func BenchmarkFibonacci(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(10)
		}
	})
}

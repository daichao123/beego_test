package components

import "testing"

//单元测试
func TestFibonacciNumbersRecursive(t *testing.T) {
	FibonacciNumbersRecursive(10)
}

func TestGetSumRecursive(t *testing.T) {
	GetSumRecursive(13)
}

func TestGetSum(t *testing.T) {
	GetSum(15)
}

//压力测试
func BenchmarkFibonacciNumbersRecursive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciNumbersRecursive(12)
	}
}

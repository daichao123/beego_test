package components

// GetSum 连续自然数的和
func GetSum(n int) (sum int) {
	if n < 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

//GetSumRecursive 递归 连续自然数的和
func GetSumRecursive(n int) (sum int) {
	if n <= 1 {
		return n
	}
	return n + GetSumRecursive(n-1)
}

//FibonacciNumbersRecursive 斐波那契数列 递归
func FibonacciNumbersRecursive(n int) (sum int) {
	if n == 1 || n == 0 {
		return 1
	}
	return FibonacciNumbersRecursive(n-1) + FibonacciNumbersRecursive(n-2)
}

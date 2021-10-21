package tests

// Fibonacci 计算第n个斐波那契数的函数
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	//递归
	return Fibonacci(n-1) + Fibonacci(n-2)
}
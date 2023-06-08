package math

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func Fibonacci(length int) []int {
	var out []int

	for i := 1; i <= length; i++ {
		out = append(out, fib(i))
	}

	return out
}

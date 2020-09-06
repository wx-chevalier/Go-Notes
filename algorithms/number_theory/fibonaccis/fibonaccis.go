package fibonaccis

// Fibonacci returns fibonacci number
func Fibonacci(number int) int {
	if number <= 0 {
		return 0
	}

	n2 := 0
	n1 := 0
	current := 1

	for i := 1; i < number; i++ {
		n2 = n1
		n1 = current
		current = n2 + n1
	}
	return current
}

// Fibonacci returns fibonacci number
func FibonacciRecursively(number int) int {
	if number <= 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	if number == 2 {
		return 1
	}

	return FibonacciRecursively(number-1) + FibonacciRecursively(number-2)
}

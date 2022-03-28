package main

func fibonacci(n int) int {
	a := 1
	b := 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

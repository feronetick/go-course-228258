package main

import "fmt"

func work(x int) int { // объявлено в тесте
	return x
}

func main() {
	cache := make(map[int]int)

	for i := 0; i < 10; i++ {
		var n int

		if _, err := fmt.Scan(&n); err != nil {
			panic(fmt.Sprintf("invalid input: %s", err))
		}

		if _, ok := cache[n]; !ok {
			cache[n] = work(n)
		}

		fmt.Print(cache[n], " ")
	}
}

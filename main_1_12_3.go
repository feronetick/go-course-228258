package main

import "fmt"

func main() {
	var a, b int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)

		if i == 0 || a > b {
			b = a
		}
	}

	fmt.Println(b)
}

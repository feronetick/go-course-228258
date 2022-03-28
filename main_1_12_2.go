package main

import (
	"fmt"
	"os"
)

func scanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return i
}

func main() {
	size := scanInt()
	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		numbers[i] = scanInt()
	}

	for i := 0; i < size; i++ {
		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(numbers[i])
	}
}

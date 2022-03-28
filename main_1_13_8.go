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
	var smallest, count int
	size := scanInt()

	for i := 0; i < size; i++ {
		current := scanInt()

		if i == 0 || current < smallest {
			smallest = current
			count = 1
		} else if current == smallest {
			count++
		}
	}

	fmt.Println(count)
}

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
	var current, biggest, count int

	for current = scanInt(); current != 0; current = scanInt() {
		if current > biggest {
			biggest = current
			count = 1
		} else if current == biggest {
			count++
		}
	}

	fmt.Println(count)
}

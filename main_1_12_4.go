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
	count := 0

	for i := 0; i < size; i++ {
		if scanInt() > 0 {
			count++
		}
	}

	fmt.Println(count)
}

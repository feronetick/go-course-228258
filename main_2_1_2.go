package main

import (
	"fmt"
)

func scanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Printf("Invalid input: %s", err)
		return 0
	}

	return i
}

func minimumFromFour() (result int) {
	for i := 0; i < 4; i++ {
		n := scanInt()

		if i == 0 || n < result {
			result = n
		}
	}

	return result
}

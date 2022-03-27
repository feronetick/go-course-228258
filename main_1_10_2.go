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
	a := scanInt()
	b := scanInt()

	if a >= b {
		fmt.Println("A must be greater than B")
		os.Exit(1)
	}

	s := 0

	for i := a; i <= b; i++ {
		s += i
	}

	fmt.Println(s)
}

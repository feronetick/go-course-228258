package main

import (
	"fmt"
	"os"
)

func scanFloat() (f float64) {
	if _, err := fmt.Scan(&f); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return f
}

func main() {
	a := scanFloat()
	b := scanFloat()

	fmt.Println((a + b) / 2)
}

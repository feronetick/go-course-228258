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
	val := scanFloat()

	if val <= 0 {
		fmt.Printf("число %2.2f не подходит", val)
		return
	}

	if val > 10000 {
		fmt.Printf("%e", val)
		return
	}

	fmt.Printf("%.4f", val*val)
}

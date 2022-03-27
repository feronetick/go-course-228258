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
	var i int

	for i = scanInt(); i <= 100; i = scanInt() {
		if i >= 10 {
			fmt.Println(i)
		}
	}
}

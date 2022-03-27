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
	degree := scanInt()

	if degree > 360 || degree < 0 {
		fmt.Println("Input must be between 0 and 360. Got ", degree, "instead")
		os.Exit(1)
	}

	minutes := degree * 2

	fmt.Println("It is", minutes/60, "hours", minutes%60, "minutes.")
}

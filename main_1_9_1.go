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
	i := scanInt()

	switch {
	case i > 0:
		fmt.Println("Число положительное")
	case i < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}

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
	c := scanInt()

	if a+b < c || b+c < a || a+c < b {
		fmt.Println("Не существует")
	} else {
		fmt.Println("Существует")
	}
}

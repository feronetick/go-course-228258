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
	k := scanInt()

	h := k / 60 / 60
	m := k / 60 % 60

	fmt.Println("It is", h, "hours", m, "minutes.")
}

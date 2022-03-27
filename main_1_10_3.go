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
	c := scanInt()
	r := 0

	for i := c; i > 0; i-- {
		v := scanInt()

		if v > 9 && v < 100 && v%8 == 0 {
			r += v
		}
	}

	fmt.Println(r)
}

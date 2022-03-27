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
	x := scanInt()
	p := float64(scanInt()) / 100
	y := scanInt()

	r := 0

	for x < y {
		x += int(float64(x) * p)

		r++
	}

	fmt.Println(r)
}

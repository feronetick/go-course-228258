package main

import (
	"fmt"
	"math"
)

func ScanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		panic(fmt.Sprintf("invalid input: %s", err))
	}

	return i
}

func main() {
	a := ScanInt()
	b := ScanInt()

	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}

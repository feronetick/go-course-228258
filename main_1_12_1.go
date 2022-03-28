package main

import (
	"fmt"
)

func main() {
	workArray := [10]uint8{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var k1, k2 uint8

	for i := 0; i < 3; i++ {
		fmt.Scan(&k1)
		fmt.Scan(&k2)

		workArray[k1], workArray[k2] = workArray[k2], workArray[k1]
	}

	for _, v := range workArray {
		fmt.Print(v, " ")
	}
}

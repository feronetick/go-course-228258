package main

import (
	"fmt"
	"os"
	"strconv"
)

func scanString() (s string) {
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return s
}

func main() {
	source := scanString()
	sum := 0

	for _, v := range source {
		if i, err := strconv.Atoi(string(v)); err == nil {
			sum += i
		} else {
			fmt.Printf("Invalid input: %s", err)
			return
		}
	}

	fmt.Println(sum)
}

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

	for len(source) > 1 {
		sum := 0

		for _, r := range source {
			if i, err := strconv.Atoi(string(r)); err == nil {
				sum += i
			} else {
				fmt.Printf("Invalid input: %s", err)
				os.Exit(1)
			}
		}

		source = fmt.Sprint(sum)
	}

	fmt.Println(source)
}
